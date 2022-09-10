package handlers

import (
	"context"

	iface "github.com/kmaguswira/coinbit/application/external_services"
	"github.com/kmaguswira/coinbit/application/usecases"
	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/lovoo/goka"
)

type aboveThresholdHandler struct {
	kafkaClient                  iface.IKafka
	processAboveThresholdUsecase usecases.IProcessAboveThreshold
}

func NewAboveThresholdHandler() aboveThresholdHandler {
	return aboveThresholdHandler{
		kafkaClient:                  kafka.KafkaClient,
		processAboveThresholdUsecase: usecases.NewProcessAboveThresholdUsecase(),
	}
}

func (t *aboveThresholdHandler) Run(ctx context.Context) func() error {
	return func() error {
		aboveThresholdGroup := goka.DefineGroup(t.kafkaClient.GetAboveThresholdGroup(),
			goka.Input(t.kafkaClient.GetDepositTopic(), new(domain.DepositAmount), t.collect),
			goka.Lookup(t.kafkaClient.GetBalanceTable(), new(domain.Wallet)),
			goka.Persist(new(domain.AboveThreshold)),
		)
		processor, err := goka.NewProcessor(t.kafkaClient.GetBrokers(), aboveThresholdGroup)
		if err != nil {
			return err
		}
		return processor.Run(ctx)
	}
}

func (t *aboveThresholdHandler) collect(ctx goka.Context, amount interface{}) {
	aboveThreshold := t.processAboveThresholdUsecase.Execute(usecases.ProcessAboveThresholdInput{
		WalletID: ctx.Key(),
		Wallet:   ctx.Lookup(t.kafkaClient.GetBalanceTable(), ctx.Key()),
		Amount:   amount,
	})

	ctx.SetValue(*aboveThreshold)
}
