package mux

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/b-nnett/codex-subscription-router/internal/state"
)

func TestUsageStatusUsesSelectedAccountCredentials(t *testing.T) {
	root := t.TempDir()
	primaryHome := filepath.Join(root, "primary")
	if err := os.MkdirAll(primaryHome, 0o700); err != nil {
		t.Fatal(err)
	}
	writeResetTestAuth(t, primaryHome)
	store, err := state.Open(filepath.Join(root, "mux"), primaryHome)
	if err != nil {
		t.Fatal(err)
	}
	server := httptest.NewServer(http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet {
			t.Fatalf("method = %s, want GET", request.Method)
		}
		if got := request.Header.Get("Authorization"); got != "Bearer secret-token" {
			t.Fatalf("authorization = %q", got)
		}
		if got := request.Header.Get("ChatGPT-Account-ID"); got != "chatgpt-account" {
			t.Fatalf("ChatGPT-Account-ID = %q", got)
		}
		if got := request.Header.Get("OAI-App-Brand"); got != "codex" {
			t.Fatalf("OAI-App-Brand = %q", got)
		}
		_, _ = response.Write([]byte(`{"rate_limit":{"allowed":true}}`))
	}))
	defer server.Close()

	multiplexer := &Multiplexer{
		store:               store,
		profileClient:       server.Client(),
		usageStatusEndpoint: server.URL,
	}
	result, err := multiplexer.UsageStatus(context.Background(), "primary")
	if err != nil {
		t.Fatal(err)
	}
	if string(result) != `{"rate_limit":{"allowed":true}}` {
		t.Fatalf("unexpected response: %s", result)
	}
}
