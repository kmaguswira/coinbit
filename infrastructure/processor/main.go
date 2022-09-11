package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/kmaguswira/coinbit/application/config"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/kmaguswira/coinbit/infrastructure/processor/handlers"
	"github.com/kmaguswira/coinbit/utils/logger"
	"golang.org/x/sync/errgroup"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	grp, ctx := errgroup.WithContext(ctx)

	// initialize
	config.Init()
	logger.Init()
	kafka.InitKafka()

	// register handlers
	logger.Log().Info("Starting collectors")
	balanceHandler := handlers.NewBalanceHandler()
	aboveThresholdHandler := handlers.NewAboveThresholdHandler()

	grp.Go(aboveThresholdHandler.Run(ctx))
	grp.Go(balanceHandler.Run(ctx))

	// wait until receive signal
	waiter := make(chan os.Signal, 1)
	signal.Notify(waiter, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-waiter:
	case <-ctx.Done():
	}

	// clean up connection
	kafka.KafkaClient.CleanUp()
	cancel()

	if err := grp.Wait(); err != nil {
		logger.Log().Error(err)
	}

	logger.Log().Info("Shutdown Gracefully")
}
