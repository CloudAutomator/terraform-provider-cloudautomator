package client

import (
	"net/http"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
}

func clientDoHandler(t *testing.T, json []byte) func(w http.ResponseWriter, r *http.Request) {
	t.Helper()

	return func(w http.ResponseWriter, r *http.Request) {
		t.Helper()

		auth := r.Header.Get("Authorization")
		if auth != "Bearer validToken" {
			w.WriteHeader(http.StatusUnauthorized)
			_, _ = w.Write([]byte("Unauthorized. Invalid access token."))
			return
		}

		w.Write(json)
	}
}

func TestClient_NewRequest(t *testing.T) {
	tests := []struct {
		name               string
		requestUrl         string
		method             string
		expectErrorMessage string
	}{
		{
			name:       "valid url",
			requestUrl: "http://localhost",
		},
		{
			name:               "invalid url",
			requestUrl:         "http://abc{",
			expectErrorMessage: "parse \"http://abc{\": invalid character \"{\" in host name",
		},
		{
			name:       "valid method",
			method:     "GET",
			requestUrl: "http://localhost",
		},
		{
			name:               "invalid method",
			method:             "{}",
			requestUrl:         "http://localhost",
			expectErrorMessage: "net/http: invalid method \"{}\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := "example"

			c, _ := New(token)
			_, err := c.newRequest(tt.method, tt.requestUrl, "")

			testErrCheck(t, "c.newRequest()", tt.expectErrorMessage, err)
		})
	}
}

func TestClient_Do(t *testing.T) {
	tests := []struct {
		name               string
		token              string
		json               []byte
		expectErrorMessage string
	}{
		{
			name:  "valid token",
			token: "validToken",
			json:  []byte(`{"Name": "ok"}`),
		},
		{
			name:               "invalid token",
			token:              "invalidToken",
			expectErrorMessage: "request failed. StatusCode=401 Reason=Unauthorized. Invalid access token.",
		},
		{
			name:               "unmarshal failed",
			token:              "validToken",
			json:               []byte(`{"Name": 1}`),
			expectErrorMessage: "request failed. StatusCode=200 Reason=unmarshal failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup()
			defer server.Close()

			ts := new(TestStruct)
			mux.HandleFunc("/test", clientDoHandler(t, tt.json))

			c, _ := New(tt.token, WithAPIEndpoint(server.URL))
			req, _ := c.newRequest(http.MethodPost, server.URL+"/test", "")
			_, err := c.do(req, ts)

			testErrCheck(t, "c.do()", tt.expectErrorMessage, err)

			if len(ts.Name) > 0 && ts.Name != "ok" {
				t.Fatalf("Name = %s, want ok", ts.Name)
			}
		})
	}
}
