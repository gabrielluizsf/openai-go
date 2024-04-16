package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"os"

	"github.com/gabrielluizsf/goxios"
)

type WhisperParams struct {
	Model, Filename, AudioFilePath string
}

func (params *WhisperParams) Response(oc OpenAIClient) (*AudioTranscriptionResponse, error) {
	file, err := os.Open(params.AudioFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", params.Filename)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	writer.WriteField("model", params.Model)

	writer.Close()

	url := BASE_URL + "/audio/transcriptions"
	client := goxios.New(oc.Context())
	headers := openaiRequestHeaders(oc, writer.FormDataContentType())
	requestOptions := goxios.RequestOpts{
		Headers: headers,
		Body:    body,
	}
	res, err := client.Post(url, &requestOptions)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var result AudioTranscriptionResponse

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		return nil, err
	}
	if res.StatusCode != 200 {
		return nil, RequestError(res.StatusCode)
	}
	return &result, nil

}
