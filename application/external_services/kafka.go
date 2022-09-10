package external_services

import "github.com/lovoo/goka"

type IKafka interface {
	CleanUp()
	InitBalanceView()
	GetBrokers() []string
	InitAboveThresholdView()
	GetBalanceGroup() goka.Group
	GetBalanceTable() goka.Table
	GetDepositTopic() goka.Stream
	EnsureTableExists(topic string)
	EnsureStreamExists(topic string)
	GetAboveThresholdTable() goka.Table
	GetAboveThresholdGroup() goka.Group
	CreateTopicManager() goka.TopicManager
	GetWalletByID(id string) (interface{}, error)
	DepositEmmiter(key string, value interface{})
	GetIsAboveThresholdByWalletID(walletID string) (interface{}, error)
}
