package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gabrielluizsf/openai-go/pkg/openai"
	"github.com/gabrielluizsf/openai-go/pkg/openai/chat"
)

func main() {
	openaiAPIKey := os.Getenv("OPENAI_KEY")

	openaiClient := openai.WithContext(context.Background(), openaiAPIKey)
	model := "gpt-3.5-turbo"
	role := chat.Role()
	chatCompletion, err := openaiClient.ChatGPT(&openai.ChatCompletionRequestParams{
		Model: model,
		Messages: []chat.Message{
			{Role: role.System().String(), Content: "Please be as sarcastic as possible."},
			{Role: role.User().String(), Content: "Sure, everything is always perfect."},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID: %s\n", chatCompletion.ID)
	fmt.Printf("Object: %s\n", chatCompletion.Object)
	fmt.Printf("Created: %d\n", chatCompletion.Created)
	fmt.Printf("Model: %s\n", chatCompletion.Model)
	fmt.Printf("System Fingerprint: %s\n", chatCompletion.SystemFingerprint)
	for _, choice := range chatCompletion.Choices {
		fmt.Printf("Choice Index: %d\n", choice.Index)
		fmt.Printf("Choice Finish Reason: %s\n", choice.FinishReason)
		fmt.Printf("Choice Message Role: %s\n", choice.Message.Role)
		fmt.Printf("Choice Message Content: %s\n", choice.Message.Content)
		if choice.Logprobs != nil {
			fmt.Printf("Choice Logprobs: %s\n", *choice.Logprobs)
		}
	}
	fmt.Printf("Prompt Tokens: %d\n", chatCompletion.Usage.PromptTokens)
	fmt.Printf("Completion Tokens: %d\n", chatCompletion.Usage.CompletionTokens)
	fmt.Printf("Total Tokens: %d\n", chatCompletion.Usage.TotalTokens)
}
