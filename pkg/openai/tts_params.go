package openai

import (
	"bytes"
	"net/http"

	"github.com/Simplou/goxios"
)

type TextToSpeechParams struct {
	Model, Input, Voice string
}

func (params *TextToSpeechParams) Response(oc OpenAIClient) (*TTSResult, error) {
	url := BASE_URL + "/audio/speech"
	openaiAPIKey := oc.getAPIKey()
	if openaiAPIKey == "" {
		return nil, InvalidAPIKey()
	}
	json := goxios.JSON{
		"model": params.Model,
		"input": params.Input,
		"voice": params.Voice,
	}
	requestBody, err := json.Marshal()
	if err != nil {
		return nil, CreateBodyError(err)
	}

	client := goxios.New(oc.Context())
	headers := openaiRequestHeaders(oc, "application/json")
	requestOptions := goxios.RequestOpts{
		Headers: headers,
		Body:    bytes.NewBuffer(requestBody),
	}
	res, err := client.Post(url, &requestOptions)
	if err != nil {
		return nil, SendRequestError(err)
	}

	if res.StatusCode != http.StatusOK {
		return nil, RequestError(res.StatusCode)
	}
	return &TTSResult{
		Audio: res.Body,
	}, nil
}
