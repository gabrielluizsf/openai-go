package openai

import (
	"net/http"

	"github.com/gabrielluizsf/openai-go/internal/request"
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
//	    {Role: "user", Content: "Hello"},
//	}
//
//	// Send chat request to GPT
//	res, err := client.ChatGPT("gpt-3.5-turbo", messages)
//	if err != nil {
//	    fmt.Println("Error:", err)
//	    return
//	}
//	fmt.Println("Response:", res)
func (oc *Client) ChatGPT(model string, messages []chat.Message, maxTokens ...int) (*chat.ChatCompletion, *OpenAIError) {
	url := BASE_URL + "/chat/completions"
	openaiAPIKey := oc.getAPIKey()
	if openaiAPIKey == "" {
		return nil, InvalidAPIKey()
	}
	requestBody := request.JSON{
		"model": model,
		"messages": func() []request.JSON {
			var requestMessages []request.JSON
			for _, message := range messages {
				requestMessages = append(requestMessages, request.JSON{
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

	requestBodyJSON, err := request.MarshalJSON(requestBody)
	if err != nil {
		return nil, CreateBodyError(err)
	}

	req, err := request.New(url, http.MethodPost, requestBodyJSON)
	if err != nil {
		return nil, CreateRequestError(err)
	}
	setHeaders(req, oc, "application/json")
	client := request.Client()
	resp, err := request.Response(req, client)
	if err != nil {
		return nil, SendRequestError(err)
	}
	defer resp.Body.Close()

	var chatCompletion chat.ChatCompletion
	err = request.DecodeJSON(resp.Body, &chatCompletion)
	if err != nil {
		return nil, DecodeJSONError(err)
	}

	return &chatCompletion, nil
}
