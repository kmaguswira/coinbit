package usecases

import (
	"fmt"
	"time"

	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/kmaguswira/coinbit/utils/logger"
)

type EmitDepositInput struct {
	WalletID string  `json:"walletId"`
	Amount   float64 `json:"amount"`
}

type IEmitDeposit interface {
	Execute(input EmitDepositInput) error
}

type emitDepositUsecase struct {
	kafka *kafka.KafkaClient
}

func NewEmitDepositUsecase() IEmitDeposit {
	return &emitDepositUsecase{
		kafka: kafka.GetClient(),
	}
}

func (t *emitDepositUsecase) Execute(input EmitDepositInput) error {
	event := domain.DepositAmount{
		Amount:    input.Amount,
		Timestamp: time.Now(),
	}

	eventProto, err := event.Encode(event)
	if err != nil {
		logger.Log().Error(err)
	}

	t.kafka.DepositEmmiter(t.kafka.DepositsTopic, input.WalletID, string(eventProto))

	logger.Log().Info(fmt.Sprintf("Deposit to %s with amount %f", input.WalletID, input.Amount))
	return nil
}
