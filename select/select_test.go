package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestWebsiteRacer(t *testing.T) {

	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	fastURL := fastServer.URL
	slowURL := slowServer.URL

	got := WebsiteRacer(fastURL, slowURL)
	if got != fastURL {
		t.Fatalf("got %q, want %q", got, fastURL)
	}
}

func makeDelayedServer(sleep time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(sleep)
		w.WriteHeader(http.StatusOK)
	}))
}
