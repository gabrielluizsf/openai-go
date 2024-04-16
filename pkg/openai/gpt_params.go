package openai

import (
	"bytes"

	"github.com/gabrielluizsf/goxios"
	"github.com/gabrielluizsf/openai-go/pkg/openai/chat"
)

type ChatCompletionRequestParams struct {
	Model     string
	Messages  []chat.Message
	MaxTokens []int
}

func (param *ChatCompletionRequestParams) Response(oc OpenAIClient) (*chat.ChatCompletion, *OpenAIError) {
	url := BASE_URL + "/chat/completions"
	openaiAPIKey := oc.getAPIKey()
	if openaiAPIKey == "" {
		return nil, InvalidAPIKey()
	}
	requestBody := goxios.JSON{
		"model": param.Model,
		"messages": func() []goxios.JSON {
			var requestMessages []goxios.JSON
			for _, message := range param.Messages {
				requestMessages = append(requestMessages, goxios.JSON{
					"role":    message.Role,
					"content": message.Content,
				})
			}
			return requestMessages
		}(),
	}

	if len(param.MaxTokens) > 0 {
		requestBody["max_tokens"] = func() int {
			sum := 0
			for _, number := range param.MaxTokens {
				sum += number
			}
			return sum
		}()
	}

	requestBodyJSON, err := requestBody.Marshal()
	if err != nil {
		return nil, CreateBodyError(err)
	}

	client := goxios.New(oc.Context())
	headers := openaiRequestHeaders(oc, "application/json")
	requestOptions := goxios.RequestOpts{
		Headers: headers,
		Body:    bytes.NewBuffer(requestBodyJSON),
	}
	res, err := client.Post(url, &requestOptions)
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
