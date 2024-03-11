package chat

type ChatCompletion struct {
	ID                string      `json:"id"`
	Object            string      `json:"object"`
	Created           int         `json:"created"`
	Model             string      `json:"model"`
	SystemFingerprint string      `json:"system_fingerprint"`
	Choices           []Choice    `json:"choices"`
	Usage             Usage       `json:"usage"`
}