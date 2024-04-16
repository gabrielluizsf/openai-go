package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gabrielluizsf/openai-go/pkg/openai"
)

func SaveTextToSpeech(client openai.OpenAIClient, model, input, voice, filePath string) error {
	ttsResult, err := client.TextToSpeech(&openai.TextToSpeechParams{
		Model: model,
		Input: input,
		Voice: voice,
	})
	if err != nil {
		return err
	}
	defer ttsResult.Audio.Close()

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, ttsResult.Audio)
	if err != nil {
		return err
	}

	fmt.Printf("Áudio salvo com sucesso em: %s\n", filePath)
	return nil
}

func main() {
	openaiClient := openai.WithContext(context.Background(),os.Getenv("OPENAI_KEY"))
	message := "oi"
	n, err := rand.Read([]byte(message))
	if err != nil {
		log.Fatal(err)
	}
	filename := fmt.Sprintf("%d%s.mp3", n, message)
	filePath := "./temp/" + filename
	if err := SaveTextToSpeech(openaiClient, "tts-1", message, "onyx", filePath); err != nil {
		log.Fatal(err)
	}
	transcription, err := openaiClient.AudioTranscription(&openai.WhisperParams{
		Model:"whisper-1", 
		Filename: filename, 
		AudioFilePath: filePath,
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf(transcription.Text)
}
