package openai

import (
	"io"
)

type TTSResult struct {
	Audio io.ReadCloser
}

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
//	ttsResult, err := openai.New("your-api-key").TextToSpeech("gpt-3.5-tts", "Hello, world!", "nova")
//	if err != nil {
//	    log.Fatal("Text-to-speech conversion failed:", err)
//	}
//	defer ttsResult.Audio.Close()
//	// Use ttsResult.Audio for further processing, e.g., saving to a file or streaming to a client.
func (oc *Client) TextToSpeech(requestParams *TextToSpeechParams) (*TTSResult, error) {
	return requestParams.Response(oc)
}
