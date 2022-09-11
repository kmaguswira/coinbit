package main

import (
	"github.com/kmaguswira/coinbit/application/config"
	"github.com/kmaguswira/coinbit/infrastructure/api/server"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/kafka"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/redis"
	"github.com/kmaguswira/coinbit/utils/logger"
)

func main() {
	// initialize
	config.Init()
	logger.Init()
	redis.InitRedis()
	kafka.InitKafka()

	kafkaClient := kafka.KafkaClient
	kafkaClient.InitBalanceView()
	kafkaClient.InitAboveThresholdView()

	server.Init()

	defer kafkaClient.CleanUp()
}
