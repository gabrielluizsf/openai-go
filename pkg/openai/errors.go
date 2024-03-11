package openai

import "fmt"

type OpenAIError struct {
	message string
}

func (oerr *OpenAIError) Error() string {
	return oerr.message
}

func CreateBodyError(err error) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Erro ao criar o corpo da solicitação: %v", err),
	}
}

func CreateRequestError(err error) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Erro ao criar a solicitação HTTP: %v", err),
	}
}

func SendRequestError(err error) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Erro ao enviar a solicitação HTTP: %v", err),
	}
}

func RequestError(statusCode int) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Request Error\nStatus Code: %d", statusCode),
	}
}

func DecodeJSONError(err error) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Erro ao decodificar a resposta: %v", err),
	}
}

func InvalidAPIKey() *OpenAIError {
	return &OpenAIError{
		message: "Invalid API KEY",
	}
}
