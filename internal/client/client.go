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
	ApiEndpoint           = "https://manager.cloudautomator.com/api/v1/"
	contentTypeHeader     = "application/json"
	defaultTimeoutSeconds = 20 * time.Second
)

var (
	userAgentHeader = fmt.Sprintf(
		"go-http/v%s (%s/%s; +go-cloudautomator-client)",
		Version,
		runtime.GOOS,
		runtime.GOARCH,
	)
)

type Client struct {
	httpClient  *http.Client
	apiEndpoint *url.URL
	token       string
}

func New(authToken string, options ...ClientOptions) (*Client, error) {
	parsedApiEndpoint, _ := url.Parse(ApiEndpoint)

	c := &Client{
		httpClient:  &http.Client{Timeout: defaultTimeoutSeconds},
		apiEndpoint: parsedApiEndpoint,
		token:       authToken,
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

func (c *Client) newRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	rel.Path = strings.TrimLeft(rel.Path, "/")
	u := c.apiEndpoint.ResolveReference(rel)
	buf := new(bytes.Buffer)
	if body != nil {
		if err := json.NewEncoder(buf).Encode(body); err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	req.Header.Set("Content-Type", contentTypeHeader)
	req.Header.Set("User-Agent", userAgentHeader)

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if v == nil {
		return resp, nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp, errors.New("read data failed")
	}

	if c := resp.StatusCode; 200 > c || 299 < c {
		return resp, fmt.Errorf(
			"request failed. StatusCode=%d Reason=%s",
			resp.StatusCode,
			string(body),
		)
	}

	if err := json.Unmarshal(body, v); err != nil {
		return resp, fmt.Errorf(
			"request failed. StatusCode=%d Reason=%s",
			resp.StatusCode,
			"unmarshal failed",
		)
	}
	time.Sleep(time.Second * 1)

	return resp, err
}
