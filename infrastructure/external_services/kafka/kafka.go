package kafka

import (
	"context"
	"fmt"

	"github.com/kmaguswira/coinbit/application/config"
	iface "github.com/kmaguswira/coinbit/application/external_services"
	"github.com/kmaguswira/coinbit/domain"
	"github.com/kmaguswira/coinbit/utils/error_mapping"
	"github.com/kmaguswira/coinbit/utils/logger"
	"github.com/lovoo/goka"
)

type kafkaClient struct {
	brokers             []string
	depositsTopic       goka.Stream
	balanceGroup        goka.Group
	aboveThresholdGroup goka.Group
	balanceTable        goka.Table
	aboveThresholdTable goka.Table
	depositEmitter      *goka.Emitter
	balanceView         *goka.View
	aboveThresholdView  *goka.View
}

var KafkaClient iface.IKafka

func InitKafka() {
	depositEmitter, err := goka.NewEmitter(config.GetConfig().KafkaBrokers, goka.Stream(config.GetConfig().KafkaDepositsTopic), new(domain.DepositAmount))
	if err != nil {
		logger.Log().Error(err)
	}

	KafkaClient = &kafkaClient{
		brokers:             config.GetConfig().KafkaBrokers,
		depositsTopic:       goka.Stream(config.GetConfig().KafkaDepositsTopic),
		balanceGroup:        goka.Group(config.GetConfig().KafkaBalanceGroup),
		aboveThresholdGroup: goka.Group(config.GetConfig().KafkaAboveThresholdGroup),
		balanceTable:        goka.GroupTable(goka.Group(config.GetConfig().KafkaBalanceGroup)),
		aboveThresholdTable: goka.GroupTable(goka.Group(config.GetConfig().KafkaAboveThresholdGroup)),
		depositEmitter:      depositEmitter,
	}

	// create topic if not exists
	KafkaClient.EnsureStreamExists(config.GetConfig().KafkaDepositsTopic)
	KafkaClient.EnsureTableExists(fmt.Sprintf("%s-table", config.GetConfig().KafkaBalanceGroup))
	KafkaClient.EnsureTableExists(fmt.Sprintf("%s-table", config.GetConfig().KafkaAboveThresholdGroup))
}

func (t *kafkaClient) GetBrokers() []string {
	return t.brokers
}

func (t *kafkaClient) GetDepositTopic() goka.Stream {
	return t.depositsTopic
}

func (t *kafkaClient) GetBalanceGroup() goka.Group {
	return t.balanceGroup
}

func (t *kafkaClient) GetAboveThresholdGroup() goka.Group {
	return t.aboveThresholdGroup
}

func (t *kafkaClient) GetBalanceTable() goka.Table {
	return t.balanceTable

}
func (t *kafkaClient) GetAboveThresholdTable() goka.Table {
	return t.aboveThresholdTable
}

func (t *kafkaClient) CleanUp() {
	t.depositEmitter.Finish()
}

func (t *kafkaClient) EnsureStreamExists(topic string) {
	tm := t.CreateTopicManager()
	defer tm.Close()
	if err := tm.EnsureStreamExists(topic, 8); err != nil {
		logger.Log().Error(err)
	}
}

func (t *kafkaClient) EnsureTableExists(topic string) {
	tm := t.CreateTopicManager()
	defer tm.Close()
	if err := tm.EnsureTableExists(topic, 8); err != nil {
		logger.Log().Error(err)
	}
}

func (t *kafkaClient) CreateTopicManager() goka.TopicManager {
	tmc := goka.NewTopicManagerConfig()
	tm, err := goka.NewTopicManager(t.brokers, goka.DefaultConfig(), tmc)
	if err != nil {
		logger.Log().Error(err)
	}
	return tm
}

func (t *kafkaClient) DepositEmmiter(key string, value interface{}) {
	if err := t.depositEmitter.EmitSync(key, value); err != nil {
		logger.Log().Error(err)
	}
}

func (t *kafkaClient) InitBalanceView() {
	view, err := goka.NewView(t.brokers, t.balanceTable, new(domain.Wallet))
	if err != nil {
		panic(err)
	}
	t.balanceView = view

	go t.balanceView.Run(context.Background())
}

func (t *kafkaClient) InitAboveThresholdView() {
	view, err := goka.NewView(t.brokers, t.aboveThresholdTable, new(domain.AboveThreshold))
	if err != nil {
		panic(err)
	}
	t.aboveThresholdView = view

	go t.aboveThresholdView.Run(context.Background())
}

func (t *kafkaClient) GetWalletByID(id string) (interface{}, error) {
	if t.balanceView == nil {
		err := fmt.Errorf("balance view not initialize")
		logger.Log().Error(err)
		return nil, err
	}

	val, err := t.balanceView.Get(id)

	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	return val, nil
}

func (t *kafkaClient) GetIsAboveThresholdByWalletID(walletID string) (interface{}, error) {
	if t.aboveThresholdView == nil {
		logger.Log().Error(fmt.Errorf("aboveThreshold view not initialize"))
		return nil, error_mapping.InternalServerError
	}

	val, err := t.aboveThresholdView.Get(walletID)

	if err != nil {
		logger.Log().Error(err)
		return nil, err
	}

	return val, nil
}
