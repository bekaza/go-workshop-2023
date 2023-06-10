package payment_test

import (
	"context"
	"example/apiwire/internal/config"
	mock_user_repo "example/apiwire/internal/repository/user/mock_user"
	"example/apiwire/internal/services/payment"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type testPaymentServiceSuite struct {
	suite.Suite
	mockUserRepo *mock_user_repo.MockIUserRepository
	underTest    payment.PaymentService
}

func (s *testPaymentServiceSuite) SetupSuite() {
	ctrl := gomock.NewController(s.T())
	s.mockUserRepo = mock_user_repo.NewMockIUserRepository(ctrl)
	appConfig := config.AppConfig{
		QrEnv: config.QrConfig{},
	}
	s.underTest = payment.ProvidePaymentService(appConfig, s.mockUserRepo)
}

func TestPaymentServiceSuite(t *testing.T) {
	suite.Run(t, &testPaymentServiceSuite{})
}

func (s *testPaymentServiceSuite) TestPaymentService_GenerateQr_PhoneNumber() {
	var (
		ctx = context.TODO()
	)

	promptPayID := "0882314328"

	qrStr, err := s.underTest.GenerateQr(ctx, promptPayID)

	s.Require().NoError(err)
	s.Equal("00020101021229370016A000000677010111011300668823143285802TH530376463048007", qrStr)
}

func (s *testPaymentServiceSuite) TestPaymentService_GenerateQr_IDCardNumber() {
	var (
		ctx = context.TODO()
	)

	promptPayID := "1100000000000"

	qrStr, err := s.underTest.GenerateQr(ctx, promptPayID)

	s.Require().NoError(err)
	s.Equal("00020101021229370016A0000006770101110211000000000005802TH530376463043A8B", qrStr)
}
