package usecases

import (
	"fmt"

	iface "github.com/kmaguswira/coinbit/application/external_services"
	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/kmaguswira/coinbit/utils/error_mapping"
	"github.com/kmaguswira/coinbit/utils/logger"
)

type GetBalanceInput struct {
	WalletID string
}

type GetBalanceOutput struct {
	WalletID       string  `json:"wallet_id"`
	Balance        float64 `json:"balance"`
	AboveThreshold bool    `json:"above_threshold"`
}

type IGetBalance interface {
	Execute(input GetBalanceInput) (*GetBalanceOutput, error)
}

type getBalanceUsecase struct {
	kafka iface.IKafka
}

func NewGetBalanceUsecase() IGetBalance {
	return &getBalanceUsecase{
		kafka: kafka.KafkaClient,
	}
}

func (t *getBalanceUsecase) Execute(input GetBalanceInput) (*GetBalanceOutput, error) {
	logger.Log().Info(fmt.Sprintf("Get balance %s", input.WalletID))

	walletInterface, err := t.kafka.GetWalletByID(input.WalletID)
	if err != nil {
		return nil, err
	}

	if walletInterface == nil {
		return nil, error_mapping.NotFoundError
	}

	aboveThresholdInterface, err := t.kafka.GetIsAboveThresholdByWalletID(input.WalletID)
	if err != nil {
		return nil, err
	}

	wallet := walletInterface.(domain.Wallet)
	aboveThreshold := aboveThresholdInterface.(domain.AboveThreshold)

	output := GetBalanceOutput{
		WalletID:       input.WalletID,
		Balance:        wallet.GetTotal(),
		AboveThreshold: aboveThreshold.AboveThreshold,
	}

	return &output, nil
}
