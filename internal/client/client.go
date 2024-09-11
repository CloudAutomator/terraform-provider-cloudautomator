package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

const (
	apiEndpoint       = "https://manager.cloudautomator.com/api/v1/"
	defaultRetryCount = 5
	delayBaseSecond   = 1
	retryLimit        = 5
	timeoutSeconds    = 20 * time.Second
)

var (
	userAgentHeader = fmt.Sprintf(
		"go-http/v%s (%s/%s; +go-cloudautomator-client)",
		version,
		runtime.GOOS,
		runtime.GOARCH,
	)

	errBadRequest          = errors.New("bad request")
	errUnauthorized        = errors.New("unauthorized")
	errForbidden           = errors.New("forbidden")
	errNotFound            = errors.New("not found")
	errMethodNotAllowed    = errors.New("method not allowed")
	errInternalServerError = errors.New("internal server error")
	errBadGateway          = errors.New("bad gateway")
	errServiceUnavailable  = errors.New("service unavailable")
	errGatewayTimeout      = errors.New("gateway timeout")
)

type Client struct {
	httpClient      *http.Client
	apiEndpoint     *url.URL
	delayBaseSecond int
	token           string
}

func New(authToken string, options ...ClientOptions) (*Client, error) {
	parsedApiEndpoint, _ := url.Parse(apiEndpoint)

	c := &Client{
		httpClient:      &http.Client{Timeout: timeoutSeconds},
		apiEndpoint:     parsedApiEndpoint,
		delayBaseSecond: delayBaseSecond,
		token:           authToken,
	}

	for _, opt := range options {
		opt(c)
	}

	return c, nil
}

type ClientOptions func(*Client)

func WithAPIEndpoint(endpoint string) ClientOptions {
	apiEndpoint, _ := url.Parse(endpoint)

	return func(c *Client) {
		c.apiEndpoint = apiEndpoint
	}
}

func WithDelayBaseSecond(seconds int) ClientOptions {
	return func(c *Client) {
		c.delayBaseSecond = seconds
	}
}

func (c *Client) requestWithRetry(method, urlStr string, requestBody, v interface{}, retry int) (*http.Response, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	rel.Path = strings.TrimLeft(rel.Path, "/")
	u := c.apiEndpoint.ResolveReference(rel)
	buf := new(bytes.Buffer)
	if requestBody != nil {
		if err := json.NewEncoder(buf).Encode(requestBody); err != nil {
			return nil, fmt.Errorf("failed to encode request body: %w", err)
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgentHeader)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if v == nil {
		return resp, nil
	}

	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if c.shouldRetry(c.parseHTTPError(resp), retry) {
		time.Sleep(c.delayTime(retry))
		return c.requestWithRetry(method, urlStr, requestBody, v, retry-1)
	}

	if c := resp.StatusCode; 200 > c || 299 < c {
		return resp, fmt.Errorf(
			"request failed. StatusCode=%d Reason=%s",
			resp.StatusCode,
			string(responseBody),
		)
	}

	if err := json.Unmarshal(responseBody, v); err != nil {
		return resp, fmt.Errorf(
			"failed to unmarshal response. StatusCode=%d Reason=unmarshal failed",
			resp.StatusCode,
		)
	}
	time.Sleep(time.Second * 1)

	return resp, err
}

func (c *Client) shouldRetry(err error, retry int) bool {
	if retry <= 0 {
		return false
	}

	switch err {
	case errInternalServerError,
		errBadGateway,
		errServiceUnavailable,
		errGatewayTimeout:
		return true
	default:
		return false
	}
}

func (c *Client) parseHTTPError(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusBadRequest: // 400
		return errBadRequest
	case http.StatusUnauthorized: // 401
		return errUnauthorized
	case http.StatusForbidden: // 403
		return errForbidden
	case http.StatusNotFound: // 404
		return errNotFound
	case http.StatusMethodNotAllowed: // 405
		return errMethodNotAllowed
	case http.StatusInternalServerError: // 500
		return errInternalServerError
	case http.StatusBadGateway: // 502
		return errBadGateway
	case http.StatusServiceUnavailable: // 503
		return errServiceUnavailable
	case http.StatusGatewayTimeout: // 504
		return errGatewayTimeout
	default:
		return nil
	}
}

func (c *Client) delayTime(retry int) time.Duration {
	return time.Duration(c.delayBaseSecond*(retryLimit-retry+1)) * time.Second
}
