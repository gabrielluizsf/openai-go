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

// AudioTranscription performs audio transcription using the specified template.
//
// Example of use:
// client := openai.New("your-api-key")
// model := "model_name"
// fileName := "audio.mp3"
// audioFilePath := "./path/to/audio.mp3"
//
// resp, err := client.AudioTranscription(model, fileName, audioFilePath)
// if err != nil {
// log.Fatalf("Error transcribing audio: %v", err)
// }
//
// fmt.Printf("Audio transcription: %s\n", resp.Text)
//
// Parameters:
// - model: Name of the model to be used for audio transcription. See documentation for supported models: https://platform.openai.com/docs/models/model-endpoint-compatibility
// - fileName: Name of the audio file.
// - audioFilePath: Path of the audio file on the local file system.
//
// Returns:
// A pointer to an AudioTranscriptionResponse containing the audio transcription and possible errors found.
//
// Possible errors:
// - If an error occurs when opening the audio file.
// - If an error occurs while creating the multipart form.
// - If an error occurs while copying the file contents to the request body.
// - If an error occurs when creating the HTTP request.
// - If an error occurs when making the HTTP request to the OpenAI service.
// - If an error occurs while decoding the JSON response from the OpenAI service.
// - If the HTTP response status code is not 200 (OK).
//
func (oc *ClientWithContext) AudioTranscription(model, fileName, audioFilePath string) (*AudioTranscriptionResponse, error) {
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
	req, err := request.NewWithContextBuffer(oc.Ctx, url, http.MethodPost, body)
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
