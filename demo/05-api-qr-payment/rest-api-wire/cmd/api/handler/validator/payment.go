package validator

import (
	"errors"
	"example/apiwire/cmd/api/handler/request"
)

func ValidateQRPayment(req request.GenerateQrRequest) error {
	if req.PromptPayID == "" {
		return errors.New("PromptPayId cannot empty")
	}

	if !(len(req.PromptPayID) == 10 || len(req.PromptPayID) == 13) {
		return errors.New("PromptPayId is invalid")
	}

	if req.Amount < 0 {
		return errors.New("amount less than zero")
	}

	return nil
}
