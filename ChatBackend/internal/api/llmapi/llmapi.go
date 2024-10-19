package llmapi

import (
	"bytes"
	"chat"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// JSON запроса к LLM
type Request struct {
	Messages  []chat.Message `json:"messages"`
	MaxTokens uint           `json:"max_tokens"`
}

type Client struct {
	URL string
}

func NewClient(url string) *Client {
	return &Client{URL: url}
}

func (client *Client) Generate(req *Request) (*chat.Message, error) {
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

	var message *chat.Message
	err = json.Unmarshal(jsonResponse, message)
	if err != nil {
		return nil, fmt.Errorf("failed to unserialize response: %w", err)
	}

	return message, nil
}
