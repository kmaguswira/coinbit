package usecases

import (
	"fmt"
	"time"

	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	pb "github.com/kmaguswira/coinbit/proto"
	"github.com/kmaguswira/coinbit/utils/logger"
	"google.golang.org/protobuf/proto"
)

type DepositInput struct {
	WalletID string  `json:"walletId"`
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
	event := pb.DepositAmount{
		Amount:    input.Amount,
		Timestamp: time.Now().UnixMilli(),
	}

	eventProto, err := proto.Marshal(&event)
	if err != nil {
		logger.Log().Error(err)
	}

	logger.Log().Info(string(eventProto))

	t.kafka.DepositEmmiter(t.kafka.DepositsTopic, input.WalletID, eventProto)

	logger.Log().Info(fmt.Sprintf("Deposit to %s with amount %f", input.WalletID, input.Amount))
	return nil
}
