package openai

import (
	"github.com/gabrielluizsf/openai-go/pkg/openai/chat"
)

// ChatGPT sends a chat request to GPT.
//
// Example:
//
//	// Create a new OpenAI client
//	client := openai.New("your-api-key")
//
//	// Define messages for the conversation
//	messages := []chat.Message{
//	    {Role: "system", Content: "You are a helpful assistant"},
//	    {Role: "user", Content: "Hi"},
//	}
//
//	// Send chat request to GPT
//	res, err := client.ChatGPT(&openai.ChatCompletionRequestParams{
//		Model: "gpt-3.5-turbo",
//		Messages: messages,
//	})
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//	fmt.Println("Response:", res)
func (oc *ClientWithContext) ChatGPT(requestParams *ChatCompletionRequestParams) (*chat.ChatCompletion, *OpenAIError) {
	return requestParams.Response(oc)
}
