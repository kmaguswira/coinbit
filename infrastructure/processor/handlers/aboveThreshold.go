package handlers

import (
	"context"
	"fmt"

	"github.com/kmaguswira/coinbit/application/usecases"
	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/lovoo/goka"
)

type aboveThresholdHandler struct {
	kafkaClient           *kafka.KafkaClient
	processDepositUsecase usecases.IProcessDeposit
}

func NewAboveThresholdHandler() aboveThresholdHandler {
	return aboveThresholdHandler{
		kafkaClient:           kafka.GetClient(),
		processDepositUsecase: usecases.NewProcessDepositUsecase(),
	}
}

func (t *aboveThresholdHandler) Run(ctx context.Context) func() error {
	return func() error {
		aboveThresholdGroup := goka.DefineGroup(t.kafkaClient.AboveThresholdGroup,
			goka.Input(t.kafkaClient.DepositsTopic, new(domain.DepositAmount), t.collect),
			goka.Join(t.kafkaClient.BalanceTable, new(domain.Wallet)),
			goka.Persist(new(domain.AboveThreshold)),
		)
		processor, err := goka.NewProcessor(t.kafkaClient.Brokers, aboveThresholdGroup)
		if err != nil {
			return err
		}
		return processor.Run(ctx)
	}
}

func (t *aboveThresholdHandler) collect(ctx goka.Context, amount interface{}) {

	var wallets []domain.Wallet

	if wallet := ctx.Lookup(t.kafkaClient.BalanceTable, ctx.Key()); wallet != nil {
		wallets = wallet.([]domain.Wallet)
	}

	fmt.Println(wallets)
	// wallets := t.processDepositUsecase.Execute(usecases.ProcessDepositInput{
	// 	WalletID: ctx.Key(),
	// 	Wallets:  ctx.Value(),
	// 	Amount:   amount,
	// })

}
