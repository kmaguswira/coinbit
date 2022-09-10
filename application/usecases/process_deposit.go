package usecases

import (
	"github.com/kmaguswira/coinbit/domain"
)

type ProcessDepositInput struct {
	WalletID string
	Wallet   interface{}
	Amount   interface{}
}

type IProcessDeposit interface {
	Execute(input ProcessDepositInput) *domain.Wallet
}

type processDepositUsecase struct{}

func NewProcessDepositUsecase() IProcessDeposit {
	return &processDepositUsecase{}
}

func (t *processDepositUsecase) Execute(input ProcessDepositInput) *domain.Wallet {
	event := input.Amount.(domain.DepositAmount)

	var wallet domain.Wallet
	if input.Wallet != nil {
		wallet = input.Wallet.(domain.Wallet)
		wallet.DepositAmount = append(wallet.DepositAmount, event)
	} else {
		wallet = domain.Wallet{
			WalletID: input.WalletID,
			DepositAmount: []domain.DepositAmount{
				{
					Amount:    event.Amount,
					Timestamp: event.Timestamp,
				},
			},
		}
	}

	return &wallet
}
