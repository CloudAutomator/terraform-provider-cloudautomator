package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

const (
	ApiEndpoint       = "https://manager.cloudautomator.com/api/v1/"
	defaultRetryCount = 5
	delayBaseSecond   = 1
	retryLimit        = 5
	timeoutSeconds    = 20 * time.Second
)

var (
	userAgentHeader = fmt.Sprintf(
		"go-http/v%s (%s/%s; +go-cloudautomator-client)",
		Version,
		runtime.GOOS,
		runtime.GOARCH,
	)

	BadRequest          = errors.New("Bad Request")
	Unauthorized        = errors.New("Unauthorized")
	Forbidden           = errors.New("Forbidden")
	NotFound            = errors.New("Not Found")
	MethodNotAllowed    = errors.New("Method Not Allowed")
	InternalServerError = errors.New("Internal Server Error")
	BadGateway          = errors.New("Bad Gateway")
	ServiceUnavailable  = errors.New("Service Unavailable")
	GatewayTimeout      = errors.New("Gateway Timeout")
)

type Client struct {
	httpClient      *http.Client
	apiEndpoint     *url.URL
	delayBaseSecond int
	token           string
}

func New(authToken string, options ...ClientOptions) (*Client, error) {
	parsedApiEndpoint, _ := url.Parse(ApiEndpoint)

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
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", userAgentHeader)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if v == nil {
		return resp, nil
	}

	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, errors.New("read data failed")
	}

	if c.shouldRetry(c.getError(resp), retry) {
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
			"request failed. StatusCode=%d Reason=%s",
			resp.StatusCode,
			"unmarshal failed",
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
	case InternalServerError,
		BadGateway,
		ServiceUnavailable,
		GatewayTimeout:
		return true
	default:
		return false
	}
}

func (c *Client) getError(res *http.Response) error {
	switch res.StatusCode {
	case http.StatusBadRequest: // 400
		return BadRequest
	case http.StatusUnauthorized: // 401
		return Unauthorized
	case http.StatusForbidden: // 403
		return Forbidden
	case http.StatusNotFound: // 404
		return NotFound
	case http.StatusMethodNotAllowed: // 405
		return MethodNotAllowed
	case http.StatusInternalServerError: // 500
		return InternalServerError
	case http.StatusBadGateway: // 502
		return BadGateway
	case http.StatusServiceUnavailable: // 503
		return ServiceUnavailable
	case http.StatusGatewayTimeout: // 504
		return GatewayTimeout
	default:
		return nil
	}
}

func (c *Client) delayTime(retry int) time.Duration {
	return time.Duration(c.delayBaseSecond*(retryLimit-retry+1)) * time.Second
}
