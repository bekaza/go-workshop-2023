package request

type GenerateQrRequest struct {
	PromptPayID string  `json:"promptPayId"`
	Amount      float64 `json:"amount"`
}
