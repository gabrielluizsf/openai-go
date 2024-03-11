package openai

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/gabrielluizsf/openai-go/internal/request"
)

type AudioTranscriptionResponse struct {
	Text string `json:"text"`
}

// model https://platform.openai.com/docs/models/model-endpoint-compatibility
func (oc *Client) AudioTranscription(model, fileName, audioFilePath string) (*AudioTranscriptionResponse, error) {
	file, err := os.Open(audioFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile("file", fileName)
	if err != nil {
		return nil, err
	}
	_, err = io.Copy(part, file)
	if err != nil {
		return nil, err
	}

	writer.WriteField("model", model)

	writer.Close()

	url := BASE_URL + "/audio/transcriptions"
	req, err := request.NewWithBuffer(url, http.MethodPost, body)
	if err != nil {
		return nil, err
	}
	setHeaders(req, oc, writer.FormDataContentType())

	client := request.Client()
	resp, err := request.Response(req, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result AudioTranscriptionResponse

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, RequestError(resp.StatusCode)
	}
	return &result, nil

}
