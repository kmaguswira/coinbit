package usecases

import (
	"fmt"

	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/utils/logger"
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
	logger.Log().Info(fmt.Sprintf("BalanceProcessor::receiving event %v with walletID %s", input.Amount, input.WalletID))
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

	logger.Log().Info(fmt.Sprintf("BalanceProcessor::event %v with walletID %s processed", input.Amount, input.WalletID))
	return &wallet
}
