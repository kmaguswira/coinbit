package usecases

import (
	"github.com/kmaguswira/coinbit/domain"
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
	var wallet domain.Wallet
	var amount domain.DepositAmount

	if input.Wallet != nil {
		wallet = input.Wallet.(domain.Wallet)
	}

	amount = input.Amount.(domain.DepositAmount)
	wallet.Deposit(amount)

	return &domain.AboveThreshold{
		WalletID:       input.WalletID,
		AboveThreshold: wallet.IsAboveThreshold(),
	}

}
