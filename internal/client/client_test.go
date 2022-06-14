package client

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type TestStruct struct {
	Name string `json:"name"`
}

func testErrCheck(t *testing.T, name string, errorMessage string, err error) bool {
	t.Helper()

	if len(errorMessage) > 0 {
		if err == nil {
			t.Fatalf("%s error = <nil>, should contain %q", name, errorMessage)
			return false
		}

		if errStr := err.Error(); !(errStr == errorMessage) {
			t.Fatalf("%s error = %q, should contain %q", name, errStr, errorMessage)
			return false
		}

		return false
	}

	if err != nil && len(errorMessage) == 0 {
		t.Fatalf("%s unexpected error: %v", name, err)
		return false
	}

	return true
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

			c, _ := New(&token)
			_, err := c.NewRequest(tt.method, tt.requestUrl, "")

			testErrCheck(t, "c.NewRequest()", tt.expectErrorMessage, err)
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
			mux := http.NewServeMux()
			server := httptest.NewServer(mux)

			defer server.Close()

			ts := new(TestStruct)
			mux.HandleFunc("/test", clientDoHandler(t, tt.json))

			c, _ := New(&tt.token, WithAPIEndpoint(server.URL))
			req, _ := c.NewRequest(http.MethodPost, server.URL+"/test", "")
			_, err := c.Do(req, ts)

			testErrCheck(t, "c.Do()", tt.expectErrorMessage, err)

			if len(ts.Name) > 0 && ts.Name != "ok" {
				t.Fatalf("Name = %s, want ok", ts.Name)
			}
		})
	}
}
