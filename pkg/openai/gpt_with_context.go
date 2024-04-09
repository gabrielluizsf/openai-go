package openai

import (
	"bytes"

	"github.com/gabrielluizsf/goxios"
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
//	res, err := client.ChatGPT("gpt-3.5-turbo", messages)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//	fmt.Println("Response:", res)
func (oc *ClientWithContext) ChatGPT(model string, messages []chat.Message, maxTokens ...int) (*chat.ChatCompletion, *OpenAIError) {
	url := BASE_URL + "/chat/completions"
	openaiAPIKey := oc.getAPIKey()
	if openaiAPIKey == "" {
		return nil, InvalidAPIKey()
	}
	requestBody := goxios.JSON{
		"model": model,
		"messages": func() []goxios.JSON {
			var requestMessages []goxios.JSON
			for _, message := range messages {
				requestMessages = append(requestMessages, goxios.JSON{
					"role":    message.Role,
					"content": message.Content,
				})
			}
			return requestMessages
		}(),
	}

	if len(maxTokens) > 0 {
		requestBody["max_tokens"] = func() int {
			sum := 0
			for _, number := range maxTokens {
				sum += number
			}
			return sum
		}()
	}

	requestBodyJSON, err := requestBody.Marshal()
	if err != nil {
		return nil, CreateBodyError(err)
	}

	client := goxios.NewClient(oc.Ctx)
	headers := openaiRequestHeaders(oc, "application/json")
	res, err := client.Post(url, headers, bytes.NewBuffer(requestBodyJSON))
	if err != nil {
		return nil, SendRequestError(err)
	}
	defer res.Body.Close()

	var chatCompletion chat.ChatCompletion
	err = goxios.DecodeJSON(res.Body, &chatCompletion)
	if err != nil {
		return nil, DecodeJSONError(err)
	}

	return &chatCompletion, nil
}
