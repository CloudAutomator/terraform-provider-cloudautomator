package client

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	mux    *http.ServeMux
	server *httptest.Server
)

func setup() {
	mux = http.NewServeMux()
	server = httptest.NewServer(mux)
}

func testHttpMethod(t *testing.T, r *http.Request, expect string) {
	t.Helper()

	if got := r.Method; got != expect {
		t.Errorf("Request method: %v, want %v", got, expect)
	}
}

func testEqual(t *testing.T, expect interface{}, got interface{}) {
	t.Helper()

	if diff := cmp.Diff(expect, got); diff != "" {
		t.Errorf("values not equal (-want / +got):\n%s", diff)
	}
}

func testErrCheck(t *testing.T, name string, errorMessage string, err error) bool {
	t.Helper()

	if len(errorMessage) > 0 {
		if err == nil {
			t.Fatalf("%s error = <nil>, should error %q", name, errorMessage)
			return false
		}

		if errStr := err.Error(); !(errStr == errorMessage) {
			t.Fatalf("%s error = %q, should error %q", name, errStr, errorMessage)
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
