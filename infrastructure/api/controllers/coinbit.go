package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/coinbit/application/usecases"
)

type coinbitController struct {
	depositUsecase    usecases.IDeposit
	getBalanceUsecase usecases.IGetBalance
}

func NewCoinbitController() *coinbitController {
	return &coinbitController{
		depositUsecase:    usecases.NewDepositUsecase(),
		getBalanceUsecase: usecases.NewGetBalanceUsecase(),
	}
}

func (t coinbitController) Deposit(c *gin.Context) {
	depositInput := usecases.DepositInput{}

	if err := c.BindJSON(&depositInput); err != nil {
		c.Error(err)
		return
	}

	if err := t.depositUsecase.Execute(depositInput); err != nil {
		c.Error(err)
		return
	}

	c.Status(http.StatusOK)
}

func (t coinbitController) GetBalance(c *gin.Context) {
	getBalanceInput := usecases.GetBalanceInput{}

	if err := c.BindJSON(&getBalanceInput); err != nil {
		c.Error(err)
		return
	}

	balance, err := t.getBalanceUsecase.Execute(getBalanceInput)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, balance)
}
