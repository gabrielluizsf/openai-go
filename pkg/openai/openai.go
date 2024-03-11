package openai

const BASE_URL = "https://api.openai.com/v1"

type Client struct {
	APIKey string
}

func New(apiKey string) *Client {
	return &Client{
		APIKey: apiKey,
	}
}
