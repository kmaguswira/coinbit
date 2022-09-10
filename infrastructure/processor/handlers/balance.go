package handlers

import (
	"context"

	iface "github.com/kmaguswira/coinbit/application/external_services"
	"github.com/kmaguswira/coinbit/application/usecases"
	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/lovoo/goka"
)

type balanceHandler struct {
	kafkaClient           iface.IKafka
	processDepositUsecase usecases.IProcessDeposit
}

func NewBalanceHandler() balanceHandler {
	return balanceHandler{
		kafkaClient:           kafka.KafkaClient,
		processDepositUsecase: usecases.NewProcessDepositUsecase(),
	}
}

func (t *balanceHandler) Run(ctx context.Context) func() error {
	return func() error {
		balanceGroup := goka.DefineGroup(t.kafkaClient.GetBalanceGroup(),
			goka.Input(t.kafkaClient.GetDepositTopic(), new(domain.DepositAmount), t.collect),
			goka.Persist(new(domain.Wallet)),
		)
		processor, err := goka.NewProcessor(t.kafkaClient.GetBrokers(), balanceGroup)
		if err != nil {
			return err
		}
		return processor.Run(ctx)
	}
}

func (t *balanceHandler) collect(ctx goka.Context, amount interface{}) {
	wallets := t.processDepositUsecase.Execute(usecases.ProcessDepositInput{
		WalletID: ctx.Key(),
		Wallet:   ctx.Value(),
		Amount:   amount,
	})

	ctx.SetValue(*wallets)
}
