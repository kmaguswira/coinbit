package handlers

import (
	"context"

	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	pb "github.com/kmaguswira/coinbit/proto"
	"github.com/lovoo/goka"
)

type balanceHandler struct {
	kafkaClient *kafka.KafkaClient
}

func NewBalanceHandler() balanceHandler {
	return balanceHandler{
		kafkaClient: kafka.GetClient(),
	}
}

func (t *balanceHandler) Run(ctx context.Context) func() error {
	return func() error {
		balanceGroup := goka.DefineGroup(t.kafkaClient.BalanceGroup,
			goka.Input(t.kafkaClient.DepositsTopic, new(domain.DepositAmount), t.collect),
			goka.Persist(new(domain.Wallet)),
		)
		processor, err := goka.NewProcessor(t.kafkaClient.Brokers, balanceGroup)
		if err != nil {
			return err
		}
		return processor.Run(ctx)
	}
}

func (t *balanceHandler) collect(ctx goka.Context, msg interface{}) {
	var wallets []pb.Wallet
	if v := ctx.Value(); v != nil {
		wallets = v.([]pb.Wallet)
	}

	event := msg.(*pb.DepositAmount)

	foundWallet := false
	for i := 0; i < len(wallets); i++ {
		if wallets[i].WalletID == ctx.Key() {
			foundWallet = true
			wallets[i].DepositAmount = append(wallets[i].DepositAmount, &pb.DepositAmount{
				Amount:    event.Amount,
				Timestamp: event.Timestamp,
			})
		}
	}

	if !foundWallet {
		wallets = append(wallets, pb.Wallet{
			WalletID: ctx.Key(),
			DepositAmount: []*pb.DepositAmount{
				{
					Amount:    event.Amount,
					Timestamp: event.Timestamp,
				},
			},
			AboveThreshold: false,
		})
	}

	ctx.SetValue(wallets)
}
