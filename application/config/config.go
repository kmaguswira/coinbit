package config

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppName                  string
	Env                      string
	Port                     string
	RedisHost                string
	RedisPassword            string
	CorsAllowOrigin          []string
	CorsMaxAge               string
	CorsAllowMethods         string
	CorsAllowHeaders         string
	CorsExposeHeaders        string
	CorsAllowCredentials     string
	KafkaBrokers             []string
	KafkaDepositsTopic       string
	KafkaBalanceGroup        string
	KafkaAboveThresholdGroup string
}

var config Config

func Init() {
	var err error

	viper.AutomaticEnv()

	viper.SetConfigName("env")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./")

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("env.yaml not found, using os env!")
	}

	config = Config{
		AppName:                  viper.GetString("APP_NAME"),
		Env:                      viper.GetString("ENVIRONMENT"),
		Port:                     viper.GetString("PORT"),
		RedisHost:                viper.GetString("REDIS_HOST"),
		RedisPassword:            viper.GetString("REDIS_PASSWORD"),
		CorsAllowOrigin:          strings.Split(viper.GetString("CORS_ALLOW_ORIGIN"), ","),
		CorsMaxAge:               viper.GetString("CORS_MAX_AGE"),
		CorsAllowMethods:         viper.GetString("CORS_ALLOWED_METHODS"),
		CorsAllowHeaders:         viper.GetString("CORS_ALLOWED_HEADERS"),
		CorsExposeHeaders:        viper.GetString("CORS_EXPOSE_HEADERS"),
		CorsAllowCredentials:     viper.GetString("CORS_ALLOW_CREDENTIALS"),
		KafkaBrokers:             strings.Split(viper.GetString("KAFKA_BROKERS"), ","),
		KafkaDepositsTopic:       viper.GetString("KAFKA_DEPOSITS_TOPIC"),
		KafkaBalanceGroup:        viper.GetString("KAFKA_BALANCE_GROUP"),
		KafkaAboveThresholdGroup: viper.GetString("KAFKA_ABOVE_THRESHOLD_GROUP"),
	}
}

func GetConfig() Config {
	return config
}
