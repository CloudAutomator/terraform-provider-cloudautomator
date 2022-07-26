package client

import (
	"net/http"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
}

func clientDoHandler(t *testing.T, statusCode int, json []byte, requestCount *int) func(w http.ResponseWriter, r *http.Request) {
	t.Helper()

	return func(w http.ResponseWriter, r *http.Request) {
		*requestCount++

		t.Helper()
		w.WriteHeader(statusCode)
		w.Write(json)
	}
}

func TestClient_RequestWithRetry(t *testing.T) {
	tests := []struct {
		name               string
		requestUrl         string
		requestMethod      string
		statusCode         int
		token              string
		json               []byte
		expectRequestCount int
		expectErrorMessage string
	}{
		{
			name:               "invalid url",
			requestUrl:         "http://abc{",
			expectRequestCount: 0,
			expectErrorMessage: "parse \"http://abc{\": invalid character \"{\" in host name",
		},
		{
			name:               "invalid method",
			requestMethod:      "{}",
			expectRequestCount: 0,
			expectErrorMessage: "net/http: invalid method \"{}\"",
		},
		{
			name:               "StatusBadRequest",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusBadRequest,
			expectRequestCount: 1,
			expectErrorMessage: "request failed. StatusCode=400 Reason=",
		},
		{
			name:               "Unauthorized",
			token:              "invalid token",
			requestMethod:      "GET",
			statusCode:         http.StatusUnauthorized,
			expectRequestCount: 1,
			expectErrorMessage: "request failed. StatusCode=401 Reason=",
		},
		{
			name:               "StatusForbidden",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusForbidden,
			expectRequestCount: 1,
			expectErrorMessage: "request failed. StatusCode=403 Reason=",
		},
		{
			name:               "StatusNotFound",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusNotFound,
			expectRequestCount: 1,
			expectErrorMessage: "request failed. StatusCode=404 Reason=",
		},
		{
			name:               "StatusMethodNotAllowed",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusMethodNotAllowed,
			expectRequestCount: 1,
			expectErrorMessage: "request failed. StatusCode=405 Reason=",
		},
		{
			name:               "InternalServerError",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusInternalServerError,
			expectRequestCount: 6,
			expectErrorMessage: "request failed. StatusCode=500 Reason=",
		},
		{
			name:               "StatusBadGateway",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusBadGateway,
			expectRequestCount: 6,
			expectErrorMessage: "request failed. StatusCode=502 Reason=",
		},
		{
			name:               "StatusServiceUnavailable",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusServiceUnavailable,
			expectRequestCount: 6,
			expectErrorMessage: "request failed. StatusCode=503 Reason=",
		},
		{
			name:               "StatusGatewayTimeout",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusGatewayTimeout,
			expectRequestCount: 6,
			expectErrorMessage: "request failed. StatusCode=504 Reason=",
		},
		{
			name:               "unmarshal failed",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusOK,
			json:               []byte(`{"Name": 1}`),
			expectRequestCount: 1,
			expectErrorMessage: "request failed. StatusCode=200 Reason=unmarshal failed",
		},
		{
			name:               "success",
			token:              "valid token",
			requestMethod:      "GET",
			statusCode:         http.StatusOK,
			expectRequestCount: 1,
			json:               []byte(`{"Name": "ok"}`),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup()
			defer server.Close()

			c, _ := New(tt.token, WithDelayBaseSecond(0))
			var requestUrl string
			requestCount := 0

			if tt.requestUrl != "" {
				requestUrl = tt.requestUrl
			} else {
				requestUrl = server.URL + "/test"
			}

			ts := new(TestStruct)
			mux.HandleFunc("/test", clientDoHandler(t, tt.statusCode, tt.json, &requestCount))
			_, err := c.requestWithRetry(tt.requestMethod, requestUrl, nil, ts, defaultRetryCount)

			if tt.expectRequestCount != requestCount {
				t.Errorf("error")
			}

			testErrCheck(t, "c.newRequest()", tt.expectErrorMessage, err)
		})
	}
}
