package request

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

type Header struct {
	Key   string
	Value any
}

func WithContext(ctx context.Context, method, url string, requestBody []byte) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(requestBody))
}

func New(url, method string, requestBody []byte) (*http.Request, error) {
	return http.NewRequest(method, url, bytes.NewBuffer(requestBody))
}

func NewWithBuffer(url, method string, requestBody io.Reader) (*http.Request, error) {
	return http.NewRequest(method, url, requestBody)
}

func NewWithContextBuffer(ctx context.Context, url, method string, requestBody io.Reader) (*http.Request, error) {
	return http.NewRequestWithContext(ctx, method, url, requestBody)
}

func SetHeaders(req *http.Request, headers []Header) {
	for _, header := range headers {
		req.Header.Set(header.Key, fmt.Sprintf("%v", header.Value))
	}
}

func ClientWithTimeout(timeout time.Duration) *http.Client {
	return &http.Client{
		Timeout: timeout,
	}
}

func Client() *http.Client {
	return &http.Client{}
}

func Response(req *http.Request, client *http.Client) (*http.Response, error) {
	return client.Do(req)
}
