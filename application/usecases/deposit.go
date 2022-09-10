package usecases

import (
	"fmt"

	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/kmaguswira/coinbit/utils/logger"
)

type DepositInput struct {
	WalletID int     `json:"walletId"`
	Amount   float64 `json:"amount"`
}

type IDeposit interface {
	Execute(input DepositInput) error
}

type depositUsecase struct {
	kafka *kafka.KafkaClient
}

func NewDepositUsecase() IDeposit {
	return &depositUsecase{
		kafka: kafka.GetClient(),
	}
}

func (t *depositUsecase) Execute(input DepositInput) error {
	t.kafka.DepositEmmiter(t.kafka.DepositsTopic, fmt.Sprintf("%d", input.WalletID), fmt.Sprintf("%f", input.Amount))
	logger.Log().Info(fmt.Sprintf("Deposit to %d with amount %f", input.WalletID, input.Amount))
	return nil
}
