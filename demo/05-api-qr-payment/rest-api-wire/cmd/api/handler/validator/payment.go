package validator

import (
	"errors"
	"example/apiwire/cmd/api/handler/request"
)

func ValidateQRPayment(req request.GenerateQrRequest) error {
	if req.PromptPayID == "" {
		return errors.New("PromptPayID")
	}

	return nil
}
