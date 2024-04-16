package openai

import (
	"context"

	"github.com/gabrielluizsf/openai-go/pkg/openai/chat"
)

const BASE_URL = "https://api.openai.com/v1"

type OpenAIClient interface {
	getAPIKey() string
	ChatGPT(*ChatCompletionRequestParams) (*chat.ChatCompletion, *OpenAIError)
	AudioTranscription(string, string, string) (*AudioTranscriptionResponse, error)
	TextToSpeech(*TextToSpeechParams) (*TTSResult, error)
	Context() context.Context
}

type Client struct {
	apiKey string
}

func (oc *Client) Context() context.Context{
	return context.Background()
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
	Client
	Ctx    context.Context
}

func (oc *ClientWithContext) Context() context.Context{
	return oc.Ctx
}


func WithContext(ctx context.Context, apiKey string) *ClientWithContext {
	return &ClientWithContext{
		Client: Client{
			apiKey: apiKey,
		},
		Ctx:    ctx,
	}
}
