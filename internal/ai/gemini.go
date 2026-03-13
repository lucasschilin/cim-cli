package ai

import (
	"context"

	"google.golang.org/genai"
)

type Gemini struct {
	client *genai.Client
	model  string
}

func NewGemini(ctx context.Context, apiKey string, model string) (*Gemini, error) {

	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey: apiKey,
	})

	if err != nil {
		return nil, err
	}

	return &Gemini{
		client: client,
		model:  model,
	}, nil
}

type geminiRequest struct {
	Contents []geminiContent `json:"contents"`
}

type geminiContent struct {
	Parts []geminiPart `json:"parts"`
}

type geminiPart struct {
	Text string `json:"text"`
}

type geminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func (g *Gemini) ImproveCommitMessage(
	ctx context.Context,
	prompt string,
) (string, error) {

	resp, err := g.client.Models.GenerateContent(
		ctx,
		g.model,
		genai.Text(prompt),
		nil,
	)

	if err != nil {
		return "", err
	}

	return resp.Text(), nil
}
