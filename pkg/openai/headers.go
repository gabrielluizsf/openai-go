package openai

import "github.com/Simplou/goxios"

func openaiRequestHeaders(oc OpenAIClient, contentType string) []goxios.Header {
	return []goxios.Header{
		{
			Key:   "Content-Type",
			Value: contentType,
		},
		{
			Key:   "Authorization",
			Value: "Bearer " + oc.getAPIKey(),
		},
	}
}
