package adapter

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type OpenAIClient struct {
	APIKey     string
	APIBase    string
	HTTPClient *http.Client
}

// NewOpenAIClient creates a new OpenAIClient instance
func NewOpenAIClient(apiKey string, apiBases ...string) *OpenAIClient {
	apiBase := "http://localhost:1234/v1"
	if len(apiBases) > 0 {
		apiBase = apiBases[0]
	}
	return &OpenAIClient{
		APIKey:     apiKey,
		APIBase:    apiBase,
		HTTPClient: &http.Client{},
	}
}

// ChatCompletion sends a chat completion request to OpenAI
func (c *OpenAIClient) ChatCompletion(prompt string, model string, maxTokens int) (string, error) {
	url := fmt.Sprintf("%s/chat/completions", c.APIBase)

	payload := map[string]interface{}{
		"model": model,
		"messages": []map[string]string{
			{"role": "system", "content": "You are a helpful assistant."},
			{"role": "user", "content": prompt},
		},
		"max_tokens": maxTokens,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("failed to marshal payload: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+c.APIKey)

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		respBody, _ := io.ReadAll(resp.Body)
		return "", fmt.Errorf("unexpected status code: %d, response: %s", resp.StatusCode, string(respBody))
	}

	var result struct {
		Choices []struct {
			Message struct {
				Content string `json:"content"`
			} `json:"message"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(result.Choices) == 0 {
		return "", errors.New("no choices in response")
	}

	return result.Choices[0].Message.Content, nil
}
