package httpapi_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/OXY7O/example-app-go/internal/httpapi"
)

func request(t *testing.T, handler http.Handler, method, path string) *httptest.ResponseRecorder {
	t.Helper()
	recorder := httptest.NewRecorder()
	handler.ServeHTTP(recorder, httptest.NewRequest(method, path, nil))
	return recorder
}

func TestHealth(t *testing.T) {
	response := request(t, httpapi.NewHandler(), http.MethodGet, "/health")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", response.Code, http.StatusOK)
	}
	if response.Header().Get("Content-Type") != "application/json" {
		t.Fatalf("content type = %q", response.Header().Get("Content-Type"))
	}
	var body map[string]string
	if err := json.Unmarshal(response.Body.Bytes(), &body); err != nil {
		t.Fatal(err)
	}
	if body["status"] != "ok" {
		t.Fatalf("body = %#v", body)
	}
}

func TestExampleByIDIsDeterministic(t *testing.T) {
	response := request(t, httpapi.NewHandler(), http.MethodGet, "/examples/42")
	if response.Code != http.StatusOK {
		t.Fatalf("status = %d", response.Code)
	}
	if got, want := response.Body.String(), "{\"id\":42,\"name\":\"example-42\"}\n"; got != want {
		t.Fatalf("body = %q, want %q", got, want)
	}
}

func TestInvalidID(t *testing.T) {
	for _, path := range []string{"/examples/not-a-number", "/examples/0", "/examples/1/extra"} {
		response := request(t, httpapi.NewHandler(), http.MethodGet, path)
		if response.Code != http.StatusBadRequest {
			t.Fatalf("path %q status = %d, want %d", path, response.Code, http.StatusBadRequest)
		}
	}
}

func TestMethodRejected(t *testing.T) {
	for _, path := range []string{"/health", "/examples/42"} {
		response := request(t, httpapi.NewHandler(), http.MethodPost, path)
		if response.Code != http.StatusMethodNotAllowed {
			t.Fatalf("path %q status = %d, want %d", path, response.Code, http.StatusMethodNotAllowed)
		}
	}
}
