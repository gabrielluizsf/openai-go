package openai

import (
	"fmt"
	"net/http"
)

type OpenAIError struct {
	message string
	statusCode  int
}

func (oerr *OpenAIError) Error() string {
	return oerr.message
}

func (oerr *OpenAIError) StatusCode() int{
	return oerr.statusCode
}

func CreateBodyError(err error) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Erro ao criar o corpo da solicitação: %v", err),
		statusCode: http.StatusBadRequest,
	}
}

func CreateRequestError(err error) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Erro ao criar a solicitação HTTP: %v", err),
		statusCode: http.StatusInternalServerError,
	}
}

func SendRequestError(err error) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Erro ao enviar a solicitação HTTP: %v", err),
		statusCode: http.StatusServiceUnavailable,
	}
}

func RequestError(statusCode int) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Request Error\nStatus Code: %d", statusCode),
		statusCode: statusCode,
	}
}

func DecodeJSONError(err error) *OpenAIError {
	return &OpenAIError{
		message: fmt.Sprintf("Erro ao decodificar a resposta: %v", err),
		statusCode: http.StatusInternalServerError,
	}
}

func InvalidAPIKey() *OpenAIError {
	return &OpenAIError{
		message: "Invalid API KEY",
		statusCode: http.StatusUnauthorized,
	}
}
