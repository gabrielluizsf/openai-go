package openai

import (
	"net/http"

	"github.com/gabrielluizsf/openai-go/internal/request"
)

func (oc *ClientWithContext) TextToSpeech(model, input, voice string) (*TTSResult, error) {
	url := BASE_URL + "/audio/speech"
	openaiAPIKey := oc.GetAPIKey()
	if openaiAPIKey == "" {
		return nil, InvalidAPIKey()
	}

	requestBody, err := request.MarshalJSON(request.JSON{
		"model": model,
		"input": input,
		"voice": voice,
	})
	if err != nil {
		return nil, CreateBodyError(err)
	}

	req, err := request.WithContext(oc.Ctx, http.MethodPost, url, requestBody)
	if err != nil {
		return nil, CreateRequestError(err)
	}
	setHeaders(req, oc, "application/json")
	client := request.Client()
	resp, err := request.Response(req, client)
	if err != nil {
		return nil, SendRequestError(err)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, RequestError(resp.StatusCode)
	}
	return &TTSResult{
		Audio: resp.Body,
	}, nil
}
