package usecases

import (
	"fmt"
	"time"

	iface "github.com/kmaguswira/coinbit/application/external_services"
	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/kmaguswira/coinbit/utils/logger"
)

type EmitDepositInput struct {
	WalletID string  `json:"wallet_id"`
	Amount   float64 `json:"amount"`
}

type IEmitDeposit interface {
	Execute(input EmitDepositInput) error
}

type emitDepositUsecase struct {
	kafka iface.IKafka
}

func NewEmitDepositUsecase() IEmitDeposit {
	return &emitDepositUsecase{
		kafka: kafka.KafkaClient,
	}
}

func (t *emitDepositUsecase) Execute(input EmitDepositInput) error {
	event := domain.DepositAmount{
		Amount:    input.Amount,
		Timestamp: time.Now(),
	}

	t.kafka.DepositEmmiter(input.WalletID, event)

	logger.Log().Info(fmt.Sprintf("Deposit to %s with amount %f", input.WalletID, input.Amount))
	return nil
}
