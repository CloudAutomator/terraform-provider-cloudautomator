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

var (
	apiEndpoint           = "https://manager.cloudautomator.com/api/v1/"
	contentTypeHeader     = "application/json"
	defaultTimeoutSeconds = 20 * time.Second
	userAgentHeader       = fmt.Sprintf(
		"go-http/v%s (%s/%s; +go-cloudautomator-client)",
		Version,
		runtime.GOOS,
		runtime.GOARCH,
	)
)

type Client struct {
	HttpClient  *http.Client
	ApiEndpoint *url.URL
	Token       *string
}

func New(authToken *string) (*Client, error) {
	parsedApiEndpoint, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}

	c := &Client{
		HttpClient:  &http.Client{Timeout: defaultTimeoutSeconds},
		ApiEndpoint: parsedApiEndpoint,
		Token:       authToken,
	}

	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	rel.Path = strings.TrimLeft(rel.Path, "/")
	u := c.ApiEndpoint.ResolveReference(rel)
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
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", *c.Token))
	req.Header.Set("Content-Type", contentTypeHeader)
	req.Header.Set("User-Agent", userAgentHeader)

	return req, nil
}

func (c *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}

	if c := resp.StatusCode; 200 > c || 299 < c {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return resp, errors.New("read data failed")
		}

		return resp, fmt.Errorf(
			"request failed. StatusCode=%d Reason=%s",
			resp.StatusCode,
			string(body),
		)
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	if v == nil {
		return resp, nil
	}

	if err := json.Unmarshal(body, v); err != nil {
		return resp, errors.New("unmarshall failed")
	}
	time.Sleep(time.Second * 1)

	return resp, err
}
