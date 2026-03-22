package ai

import (
	"context"
	"fmt"

	"github.com/lucasschilin/cim-cli/internal/config"
)

func NewProvider(ctx context.Context, cfg *config.Config) (Provider, error) {

	switch cfg.Provider {

	case "gemini":
		return NewGemini(ctx, cfg.Gemini.APIKey, cfg.Model)

	case "openai":
		return NewOpenAI(cfg.Openai.APIKey, cfg.Model), nil

	case "ollama":
		return NewOllama(
			cfg.Ollama.BaseURL,
			cfg.Model,
			*cfg.ImprovementRequestTimeout,
		), nil

	default:
		return nil, fmt.Errorf("unknown AI provider: %s", cfg.Provider)

	}

}
