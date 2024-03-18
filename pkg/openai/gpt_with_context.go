package openai

import (
	"net/http"

	"github.com/gabrielluizsf/openai-go/internal/request"
	"github.com/gabrielluizsf/openai-go/pkg/openai/chat"
)

// model https://platform.openai.com/docs/models/model-endpoint-compatibility
func (oc *ClientWithContext) ChatGPT(model string, messages []chat.Message, maxTokens ...int) (*chat.ChatCompletion, *OpenAIError) {
	url := BASE_URL + "/chat/completions"
	openaiAPIKey := oc.GetAPIKey()
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

	req, err := request.WithContext(oc.Ctx, http.MethodPost, url, requestBodyJSON)
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
