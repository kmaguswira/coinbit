package usecases

import (
	"fmt"

	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/utils/logger"
)

type ProcessAboveThresholdInput struct {
	WalletID string
	Wallet   interface{}
	Amount   interface{}
}

type IProcessAboveThreshold interface {
	Execute(input ProcessAboveThresholdInput) *domain.AboveThreshold
}

type processAboveThresholdUsecase struct{}

func NewProcessAboveThresholdUsecase() IProcessAboveThreshold {
	return &processAboveThresholdUsecase{}
}

func (t *processAboveThresholdUsecase) Execute(input ProcessAboveThresholdInput) *domain.AboveThreshold {
	logger.Log().Info(fmt.Sprintf("AboveThresholdProcessor::receiving event %v with walletID %s", input.Amount, input.WalletID))
	var wallet domain.Wallet
	var amount domain.DepositAmount

	if input.Wallet != nil {
		wallet = input.Wallet.(domain.Wallet)
	}

	amount = input.Amount.(domain.DepositAmount)
	wallet.Deposit(amount)

	logger.Log().Info(fmt.Sprintf("AboveThresholdProcessor::event %v with walletID %s processed", input.Amount, input.WalletID))
	return &domain.AboveThreshold{
		WalletID:       input.WalletID,
		AboveThreshold: wallet.IsAboveThreshold(),
	}

}
