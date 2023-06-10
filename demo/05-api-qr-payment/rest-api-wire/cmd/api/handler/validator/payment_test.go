package validator_test

import (
	"example/apiwire/cmd/api/handler/request"
	"example/apiwire/cmd/api/handler/validator"
	"testing"

	"github.com/stretchr/testify/suite"
)

type testPaymentValidatorSuite struct {
	suite.Suite
}

func (t *testPaymentValidatorSuite) SetupSuite() {

}

func TestPaymentValidatorSuite(t *testing.T) {
	suite.Run(t, &testPaymentValidatorSuite{})
}

func (t *testPaymentValidatorSuite) TestPaymentValidator_PromptPayIDEmpty() {
	var req request.GenerateQrRequest

	err := validator.ValidateQRPayment(req)

	t.Require().Error(err)
	t.Equal("PromptPayId cannot empty", err.Error())
}

func (t *testPaymentValidatorSuite) TestPaymentValidator_PromptPayIDInvalid() {
	t.Run("PromptPayID is 94855 (5 len)", func() {
		req := request.GenerateQrRequest{
			PromptPayID: "94855",
		}

		err := validator.ValidateQRPayment(req)

		t.Require().Error(err)
		t.Equal("PromptPayId is invalid", err.Error())
	})

	t.Run("PromptPayID is 11111111111 (11 len)", func() {
		req := request.GenerateQrRequest{
			PromptPayID: "11111111111",
		}

		err := validator.ValidateQRPayment(req)

		t.Require().Error(err)
		t.Equal("PromptPayId is invalid", err.Error())
	})

	t.Run("PromptPayID is 1111111111111111 (16 len)", func() {
		req := request.GenerateQrRequest{
			PromptPayID: "1111111111111111",
		}

		err := validator.ValidateQRPayment(req)

		t.Require().Error(err)
		t.Equal("PromptPayId is invalid", err.Error())
	})
}

func (t *testPaymentValidatorSuite) TestPaymentValidator_AmountLessThanZero() {
	req := request.GenerateQrRequest{
		PromptPayID: "0888888888",
		Amount:      -10,
	}

	err := validator.ValidateQRPayment(req)

	t.Require().Error(err)
	t.Equal("Amount less than zero", err.Error())
}

func (t *testPaymentValidatorSuite) TestPaymentValidator_Success() {

}
