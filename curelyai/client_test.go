package curelyai

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestChatClient(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message":"Hello, World!"}`))
	}))
	defer server.Close()

	client := NewChatClient("test-key", 5*time.Second)
	client.BaseURL = server.URL

	response, err := client.Chat(context.Background(), "Hello")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if response != "Hello, World!" {
		t.Fatalf("expected 'Hello, World!', got %s", response)
	}
}
