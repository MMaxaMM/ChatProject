package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type MessageRole string

const (
	RoleSystem    = "system"
	RoleUser      = "user"
	RoleAssistent = "assistent"
)

// JSON ответа от LLM
type Message struct {
	Role    MessageRole `json:"role"`
	Content string      `json:"content"`
}

// JSON запроса к LLM
type ChatRequest struct {
	Messages  []Message `json:"messages"`
	MaxTokens uint      `json:"max_tokens"`
}

type Client struct {
	URL       string
	MaxTokens uint
}

func (client *Client) Generate(req *ChatRequest) (*Message, error) {
	jsonRequest, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to serialize request: %w", err)
	}

	resp, err := http.Post(client.URL, "application/json", bytes.NewReader(jsonRequest))
	if err != nil {
		return nil, fmt.Errorf("failed access to LLM server: %w", err)
	}

	jsonResponse, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var message *Message
	err = json.Unmarshal(jsonResponse, message)
	if err != nil {
		return nil, fmt.Errorf("failed to unserialize response: %w", err)
	}

	return message, nil
}

func NewClient(url string, maxTokens uint) *Client {
	return &Client{URL: url, MaxTokens: maxTokens}
}
