package usecases

import (
	"fmt"

	"github.com/kmaguswira/coinbit/utils/logger"
)

type GetBalanceInput struct {
	WalletID int
}

type GetBalanceOutput struct {
	WalletID       int
	Balance        float64
	AboveThreshold bool
}

type IGetBalance interface {
	Execute(input GetBalanceInput) (*GetBalanceOutput, error)
}

type getBalanceUsecase struct{}

func NewGetBalanceUsecase() IGetBalance {
	return &getBalanceUsecase{}
}

func (t *getBalanceUsecase) Execute(input GetBalanceInput) (*GetBalanceOutput, error) {
	logger.Log().Info(fmt.Sprintf("Get balance %d", input.WalletID))

	return nil, nil
}
