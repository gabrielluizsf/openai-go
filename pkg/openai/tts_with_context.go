package openai

import (
	"bytes"
	"net/http"

	"github.com/gabrielluizsf/goxios"
)

// TextToSpeech converts the given text input into speech using the specified model and voice.
//
// Parameters:
//   - model: A string representing the model to use for text-to-speech conversion.
//   - input: The input text to be converted into speech.
//   - voice: A string representing the voice to be used for speech synthesis.
//
// Returns:
//   - *TTSResult: A pointer to a TTSResult struct containing the synthesized audio as an io.ReadCloser.
//   - error: An error if any occurred during the text-to-speech conversion process.
//
// Example:
//
//	ttsResult, err := openai.New("your-api-key").TextToSpeech("gpt-3.5-tts", "Hello, world!", "onyx")
//	if err != nil {
//	    log.Fatal("Text-to-speech conversion failed:", err)
//	}
//	defer ttsResult.Audio.Close()
//	// Use ttsResult.Audio for further processing, e.g., saving to a file or streaming to a client.
func (oc *ClientWithContext) TextToSpeech(model, input, voice string) (*TTSResult, error) {
	url := BASE_URL + "/audio/speech"
	openaiAPIKey := oc.getAPIKey()
	if openaiAPIKey == "" {
		return nil, InvalidAPIKey()
	}
	json := goxios.JSON{
		"model": model,
		"input": input,
		"voice": voice,
	}
	requestBody, err := json.Marshal()
	if err != nil {
		return nil, CreateBodyError(err)
	}

	client := goxios.NewClient(oc.Ctx)
	headers := openaiRequestHeaders(oc, "application/json")
	res, err := client.Post(url, headers, bytes.NewBuffer(requestBody))
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
