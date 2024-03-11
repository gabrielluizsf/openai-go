package openai

import (
	"net/http"

	"github.com/gabrielluizsf/openai-go/internal/request"
)

func setHeaders(req *http.Request, oc *Client, contentType string) {
	request.SetHeaders(
		req,
		[]request.Header{
			{
				Key:   "Content-Type",
				Value: contentType,
			},
			{
				Key:   "Authorization",
				Value: "Bearer " + oc.APIKey,
			},
		})
}
