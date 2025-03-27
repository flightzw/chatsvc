package openai

import (
	"context"

	"github.com/pkg/errors"
)

const (
	_pathChatCompletion string = "/chat/completions"
)

type ChatOptions struct {
	UserID int32
}

type AIChatClient interface {
	ChatCompletion(ctx context.Context, req *ChatCompletionRequest, opts *ChatOptions) (*ChatCompletionReply, error)
}

func NewAIChatClient(modelType ModelType, apiKey string, debug bool) (AIChatClient, error) {
	switch modelType {
	case ModelTypeXFYunLite, ModelTypeXFYunUltraV4:
		return NewXFYunClient(apiKey, debug), nil
	case ModelTypeDeepSeekChat:
		return NewDeepSeekClient(apiKey, debug), nil
	}
	return nil, errors.Errorf("ModelType is invalid: %v", modelType)
}
