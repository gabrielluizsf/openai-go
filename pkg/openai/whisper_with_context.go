package openai

// AudioTranscription performs audio transcription using the specified template.
//
// Example of use:
//  client := openai.WithContext(context.Background(),os.Getenv("OPENAI_KEY"))
//  model := "model_name"
//  filename := "audio.mp3"
//  audioFilePath := "./path/to/audio.mp3"
//
//	resp, err := client.AudioTranscription(&openai.WhisperParams{
//				Model:"whisper-1",
//				Filename: filename,
//				AudioFilePath: filePath,
//	})
//
//  if err != nil {
//    log.Fatalf("Error transcribing audio: %v", err)
//  }
//
//  fmt.Printf("Audio transcription: %s\n", resp.Text)
//
// WhisperParams:
//  - Model: Name of the model to be used for audio transcription. See documentation for supported models: https://platform.openai.com/docs/models/model-endpoint-compatibility
//  - Filename: Name of the audio file.
//  - AudioFilePath: Path of the audio file on the local file system.
//
// Returns:
//  A pointer to an AudioTranscriptionResponse containing the audio transcription and possible errors found.
//
// Possible errors:
//  - If an error occurs when opening the audio file.
//  - If an error occurs while creating the multipart form.
//  - If an error occurs while copying the file contents to the request body.
//  - If an error occurs when creating the HTTP request.
//  - If an error occurs when making the HTTP request to the OpenAI service.
//  - If an error occurs while decoding the JSON response from the OpenAI service.
//  - If the HTTP response status code is not 200 (OK).
func (oc *ClientWithContext) AudioTranscription(requestParams *WhisperParams) (*AudioTranscriptionResponse, error) {
	return requestParams.Response(oc)
}
