package handlers

import (
	"context"

	"github.com/kmaguswira/coinbit/application/usecases"
	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/lovoo/goka"
)

type aboveThresholdHandler struct {
	kafkaClient                  *kafka.KafkaClient
	processAboveThresholdUsecase usecases.IProcessAboveThreshold
}

func NewAboveThresholdHandler() aboveThresholdHandler {
	return aboveThresholdHandler{
		kafkaClient:                  kafka.GetClient(),
		processAboveThresholdUsecase: usecases.NewProcessAboveThresholdUsecase(),
	}
}

func (t *aboveThresholdHandler) Run(ctx context.Context) func() error {
	return func() error {
		aboveThresholdGroup := goka.DefineGroup(t.kafkaClient.AboveThresholdGroup,
			goka.Input(t.kafkaClient.DepositsTopic, new(domain.DepositAmount), t.collect),
			goka.Lookup(t.kafkaClient.BalanceTable, new(domain.Wallet)),
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
	aboveThreshold := t.processAboveThresholdUsecase.Execute(usecases.ProcessAboveThresholdInput{
		WalletID: ctx.Key(),
		Wallet:   ctx.Lookup(t.kafkaClient.BalanceTable, ctx.Key()),
		Amount:   amount,
	})

	ctx.SetValue(*aboveThreshold)
}
