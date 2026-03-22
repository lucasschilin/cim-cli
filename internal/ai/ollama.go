package ai

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
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
		return "", err
	}

	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodPost,
		o.baseURL+"/api/generate",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := o.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("ollama request failed: %w", err)
	}
	defer resp.Body.Close()

	var result ollamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	return result.Response, nil
}
