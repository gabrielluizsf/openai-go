package openai

import (
	"context"

	"github.com/gabrielluizsf/openai-go/pkg/openai/chat"
)

const BASE_URL = "https://api.openai.com/v1"

type OpenAIClient interface {
	getAPIKey() string
	ChatGPT(string, []chat.Message, ...int) (*chat.ChatCompletion, *OpenAIError)
	AudioTranscription(string, string, string) (*AudioTranscriptionResponse, error)
	TextToSpeech(string, string, string) (*TTSResult, error)
}

type Client struct {
	apiKey string
}

func (oc *Client) getAPIKey() string {
	return oc.apiKey
}
// New creates a new OpenAI client with the provided API key.
//
// Example of use:
//
// 	apiKey := "your-api-key-here"
// 	client := openai.New(apiKey)
// 
func New(apiKey string) *Client {
	return &Client{
		apiKey: apiKey,
	}
}

type ClientWithContext struct {
	apiKey string
	Ctx    context.Context
}

func (oc *ClientWithContext) getAPIKey() string {
	return oc.apiKey
}

func WithContext(ctx context.Context, apiKey string) *ClientWithContext {
	return &ClientWithContext{
		apiKey: apiKey,
		Ctx:    ctx,
	}
}
