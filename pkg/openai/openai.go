package openai

import (
	"context"

	"github.com/gabrielluizsf/openai-go/pkg/openai/chat"
)

const BASE_URL = "https://api.openai.com/v1"

type OpenAIClient interface {
	GetAPIKey() string
	ChatGPT(string, []chat.Message, ...int) (*chat.ChatCompletion, *OpenAIError)
	AudioTranscription(string, string, string) (*AudioTranscriptionResponse, error)
	TextToSpeech(string, string, string) (*TTSResult, error)
}

type Client struct {
	APIKey string
}

func (oc *Client) GetAPIKey() string {
	return oc.APIKey
}

func New(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}

type ClientWithContext struct {
	APIKey string
	Ctx    context.Context
}

func (oc *ClientWithContext) GetAPIKey() string {
	return oc.APIKey
}

func WithContext(ctx context.Context, apiKey string) *ClientWithContext {
	return &ClientWithContext{
		APIKey: apiKey,
		Ctx:    ctx,
	}
}
