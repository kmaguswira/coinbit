package usecases

import (
	"github.com/kmaguswira/coinbit/domain"
)

type ProcessAboveThresholdInput struct {
	WalletID        string
	Wallets         interface{}
	AboveThresholds interface{}
	Amount          interface{}
}

type IProcessAboveThreshold interface {
	Execute(input ProcessAboveThresholdInput) *[]domain.AboveThreshold
}

type processAboveThresholdUsecase struct{}

func NewProcessAboveThresholdUsecase() IProcessAboveThreshold {
	return &processAboveThresholdUsecase{}
}

func (t *processAboveThresholdUsecase) Execute(input ProcessAboveThresholdInput) *[]domain.AboveThreshold {
	// var aboveThresholds []domain.AboveThreshold
	// if input.AboveThresholds != nil {
	// 	aboveThresholds = input.AboveThresholds.([]domain.AboveThreshold)
	// }

	// event := input.Amount.(domain.DepositAmount)

	// foundWallet := false
	// for i := 0; i < len(wallets); i++ {
	// 	if wallets[i].WalletID == input.WalletID {
	// 		foundWallet = true
	// 		wallets[i].Deposit(domain.DepositAmount{
	// 			Amount:    event.Amount,
	// 			Timestamp: event.Timestamp,
	// 		})
	// 	}
	// }

	// if !foundWallet {
	// 	wallets = append(wallets, domain.Wallet{
	// 		WalletID: input.WalletID,
	// 		DepositAmount: []domain.DepositAmount{
	// 			{
	// 				Amount:    event.Amount,
	// 				Timestamp: event.Timestamp,
	// 			},
	// 		},
	// 	})
	// }
	return &[]domain.AboveThreshold{}
}
