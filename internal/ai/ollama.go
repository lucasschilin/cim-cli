package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type Ollama struct {
	baseURL string       // URL
	model   string       // modelo
	client  *http.Client // client HTTP reutilizável
}

func NewOllama(baseURL string, model string, timeout int) *Ollama {
	return &Ollama{
		baseURL: baseURL,
		model:   model,
		client: &http.Client{
			Timeout: time.Duration(timeout) * time.Second,
		},
	}
}

type ollamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type ollamaResponse struct {
	Response string `json:"response"`
	Error    string `json:"error"` // captura erro do Ollama
}

func (o *Ollama) ImproveCommitMessage(
	ctx context.Context,
	prompt string,
) (string, error) {

	reqBody := ollamaRequest{
		Model:  o.model,
		Prompt: prompt,
		Stream: false,
	}

	jsonData, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		o.baseURL+"/api/generate",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := o.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ollama request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	var result ollamaResponse
	_ = json.Unmarshal(body, &result)

	if resp.StatusCode != http.StatusOK {
		if result.Error != "" {
			return "", fmt.Errorf("ollama error: %s", result.Error)
		}
		return "", fmt.Errorf(
			"ollama returned status %d: %s",
			resp.StatusCode,
			string(body),
		)
	}

	if result.Error != "" {
		return "", fmt.Errorf("ollama error: %s", result.Error)
	}

	if strings.TrimSpace(result.Response) == "" {
		return "", fmt.Errorf("ollama returned empty response")
	}

	return result.Response, nil
}
