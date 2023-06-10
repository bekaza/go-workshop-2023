package handler_test

import (
	"context"
	"errors"
	"net/http"
	"testing"

	"example/apiwire/cmd/api/handler"
	"example/apiwire/cmd/api/handler/request"
	mock_payment_service "example/apiwire/internal/services/payment/mock_payment"
	"example/apiwire/pkg/testutils"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type testPaymentHandlerSuite struct {
	suite.Suite
	mockPaymentService *mock_payment_service.MockPaymentService
	underTest          handler.PaymentHandler
}

func (h *testPaymentHandlerSuite) SetupSuite() {
	ctrl := gomock.NewController(h.T())
	h.mockPaymentService = mock_payment_service.NewMockPaymentService(ctrl)
	h.underTest = handler.ProvidePaymentHandler(h.mockPaymentService)
}

func TestPaymentHandlerSuite(t *testing.T) {
	suite.Run(t, &testPaymentHandlerSuite{})
}

func (h *testPaymentHandlerSuite) TestPaymentHandler_BindJson() {
	res, gCtx := testutils.NewWithRequestContext(http.MethodPost, "/qr", testutils.JSON("///"))

	h.underTest.GenerateQr(gCtx)

	h.Equal(http.StatusBadRequest, res.Code)
	var response map[string]interface{}
	testutils.DecodeJSON(res.Body, &response)
	h.Equal("invalid request", response["message"])
}

func (h *testPaymentHandlerSuite) TestPaymentHandler_ValidateInvalid() {
	req := request.GenerateQrRequest{
		PromptPayID: "",
		Amount:      -10,
	}
	res, gCtx := testutils.NewWithRequestContext(http.MethodPost, "/qr", testutils.JSON(req))

	h.underTest.GenerateQr(gCtx)

	h.Equal(http.StatusBadRequest, res.Code)
	var response map[string]interface{}
	testutils.DecodeJSON(res.Body, &response)
	h.Equal("invalid request", response["message"])
}

func (h *testPaymentHandlerSuite) TestPaymentHandler_QrServiceInternalError() {
	var (
		ctx = context.TODO()
	)
	req := request.GenerateQrRequest{
		PromptPayID: "0888888888",
		Amount:      10,
	}
	res, gCtx := testutils.NewWithRequestContext(http.MethodPost, "/qr", testutils.JSON(req))
	h.mockPaymentService.EXPECT().GenerateQr(ctx, req.PromptPayID).Return("", errors.New("internal error"))

	h.underTest.GenerateQr(gCtx)

	h.Equal(http.StatusInternalServerError, res.Code)
	var response map[string]interface{}
	testutils.DecodeJSON(res.Body, &response)
	h.Equal("internal error", response["message"])
}
