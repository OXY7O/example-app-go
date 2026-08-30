package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestNewServerUsesSafeDefaults(t *testing.T) {
	server := newServer("")
	if server.Addr != ":8080" {
		t.Fatalf("address = %q", server.Addr)
	}
	if server.Handler == nil {
		t.Fatal("handler must be configured")
	}
	if server.ReadHeaderTimeout != 5*time.Second || server.ReadTimeout != 10*time.Second || server.WriteTimeout != 10*time.Second || server.IdleTimeout != 60*time.Second {
		t.Fatal("server timeouts do not match the controlled baseline")
	}
}

func TestNewServerAcceptsControlledAddress(t *testing.T) {
	server := newServer("127.0.0.1:9090")
	if server.Addr != "127.0.0.1:9090" {
		t.Fatalf("address = %q", server.Addr)
	}
	recorder := httptest.NewRecorder()
	server.Handler.ServeHTTP(recorder, httptest.NewRequest(http.MethodGet, "/health", nil))
	if recorder.Code != http.StatusOK {
		t.Fatalf("health status = %d", recorder.Code)
	}
}
