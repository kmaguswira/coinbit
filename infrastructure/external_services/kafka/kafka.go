package kafka

import (
	"github.com/kmaguswira/coinbit/application/config"
	"github.com/kmaguswira/coinbit/utils/logger"
	"github.com/lovoo/goka"
	"github.com/lovoo/goka/codec"
)

type KafkaClient struct {
	Brokers             []string
	DepositsTopic       goka.Stream
	BalanceGroup        goka.Group
	AboveThresholdGroup goka.Group
	BalanceTable        goka.Table
	AboveThresholdTable goka.Table
}

var client *KafkaClient

func InitKafka() {
	client = &KafkaClient{
		Brokers:             config.GetConfig().KafkaBrokers,
		DepositsTopic:       goka.Stream(config.GetConfig().KafkaDepositsTopic),
		BalanceGroup:        goka.Group(config.GetConfig().KafkaBalanceGroup),
		AboveThresholdGroup: goka.Group(config.GetConfig().KafkaAboveThresholdGroup),
	}

	client.BalanceTable = goka.GroupTable(client.BalanceGroup)
	client.AboveThresholdTable = goka.GroupTable(client.AboveThresholdGroup)

	client.ensureStreamExists(config.GetConfig().KafkaDepositsTopic)
}

func GetClient() *KafkaClient {
	return client
}

func (t *KafkaClient) ensureStreamExists(topic string) {
	tm := t.createTopicManager()
	defer tm.Close()
	if err := tm.EnsureStreamExists(topic, 8); err != nil {
		logger.Log().Error(err)
	}
}

func (t *KafkaClient) ensureTableExists(topic string) {
	tm := t.createTopicManager()
	defer tm.Close()
	if err := tm.EnsureTableExists(topic, 8); err != nil {
		logger.Log().Error(err)
	}
}

func (t *KafkaClient) createTopicManager() goka.TopicManager {
	tmc := goka.NewTopicManagerConfig()
	tm, err := goka.NewTopicManager(t.Brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		logger.Log().Error(err)
	}
	return tm
}

func (t *KafkaClient) DepositEmmiter(topic goka.Stream, key string, value interface{}) {
	emitter, err := goka.NewEmitter(t.Brokers, topic, new(codec.String))
	if err != nil {
		logger.Log().Error(err)
	}

	if err := emitter.EmitSync(key, value); err != nil {
		logger.Log().Error(err)
	}

}