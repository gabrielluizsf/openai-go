package request

import (
	"encoding/json"
	"io"
)

type JSON map[string]interface{}

func MarshalJSON(data JSON) ([]byte, error) {
	return json.Marshal(data)
}

func DecodeJSON(body io.Reader, pointer any) error {
	return json.NewDecoder(body).Decode(pointer)
}

func UnmarshalJSON(data []byte, pointer any) error {
	return json.Unmarshal(data, pointer)
}
