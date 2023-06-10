package payment

import (
	"context"
	"example/apiwire/internal/config"
	userRepo "example/apiwire/internal/repository/user"
	"fmt"

	"github.com/google/wire"
	"github.com/sigurn/crc16"
)

const (
	InitPayload         = "000201"
	PointOfMethod       = "010212"
	MerchantAccountCode = "2937"
	AID                 = "0016A000000677010111"
	CountryCode         = "5802TH"
	CurrencyCode        = "5303764"
	CheckSumCode        = "6304"

	ThaiPhoneTypeCode    = "01130066"
	IDCardNumberCode     = "02"
	IDCardNumberTypeCode = "02"
	AmountCode           = "54"

	PhoneLength = 10
)

var PaymentServiceSet = wire.NewSet(ProvidePaymentService)

//go:generate mockgen -source=./payment.go -destination=./mock_payment/mock_payment_service.go -package=mock_payment_service
type PaymentService interface {
	GenerateQr(ctx context.Context, promptPayID string) (string, error)
}

type paymentServiceImpl struct {
	QrConfig config.QrConfig
	UserRepo userRepo.IUserRepository
}

func ProvidePaymentService(appConfig config.AppConfig, userRepo userRepo.IUserRepository) PaymentService {
	return &paymentServiceImpl{
		QrConfig: appConfig.QrEnv,
		UserRepo: userRepo,
	}
}

func (p paymentServiceImpl) GenerateQr(ctx context.Context, promptPayID string) (string, error) {
	resultQR := InitPayload + PointOfMethod + MerchantAccountCode + AID
	if PhoneLength == len(promptPayID) {
		resultQR += ThaiPhoneTypeCode + promptPayID[1:]
	} else {
		resultQR += IDCardNumberCode + promptPayID
	}
	resultQR += CountryCode + CurrencyCode + CheckSumCode
	h := crc16.New(crc16.MakeTable(crc16.CRC16_CCITT_FALSE))
	_, err := h.Write([]byte(resultQR))
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s%X", resultQR, h.Sum16()), nil
}
