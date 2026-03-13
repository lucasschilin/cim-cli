package ai

import "context"

type Provider interface {
	ImproveCommitMessage(ctx context.Context, prompt string) (string, error)
}
