package curelyai_go

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

const (
	defaultBaseURL = "http://localhost:8000" // Replace with your actual API base URL
)

// ChatClient represents a client for interacting with the Curely AI chat API.
type ChatClient struct {
	BotKey  string
	BaseURL string
	Client  *http.Client
}

// NewChatClient creates a new ChatClient instance with the specified bot key.
func NewChatClient(botKey string) *ChatClient {
	return &ChatClient{
		BotKey:  botKey,
		BaseURL: defaultBaseURL,
		Client:  &http.Client{Timeout: 10 * time.Second},
	}
}

// Chat sends a message to the chatbot and returns the response.
func (c *ChatClient) Chat(ctx context.Context, message string) (string, error) {
	url := c.BaseURL + "/chat"
	payload, err := json.Marshal(map[string]string{
		"message": message,
	})
	if err != nil {
		return "", err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", url, bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("bot_key", c.BotKey)

	resp, err := c.Client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", errors.New("failed to get a successful response from the server")
	}

	var result struct {
		Message string `json:"message"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Message, nil
}
