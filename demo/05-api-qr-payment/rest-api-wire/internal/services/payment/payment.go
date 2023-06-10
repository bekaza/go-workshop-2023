package payment

import (
	"context"
	"example/apiwire/internal/config"
	userRepo "example/apiwire/internal/repository/user"

	"github.com/google/wire"
)

var PaymentServiceSet = wire.NewSet(ProvidePaymentService)

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
	return "", nil
}
