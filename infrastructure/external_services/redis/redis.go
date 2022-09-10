package redis

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/kmaguswira/coinbit/application/config"
	iface "github.com/kmaguswira/coinbit/application/external_services"
	"github.com/kmaguswira/coinbit/utils/logger"
)

var RedisClient iface.IRedis

type redisClient struct {
	context context.Context
	client  *redis.Client
}

func newRedis() iface.IRedis {
	return &redisClient{
		context: context.Background(),
		client: redis.NewClient(&redis.Options{
			Addr:     config.GetConfig().RedisHost,
			Password: config.GetConfig().RedisPassword,
			DB:       0,
		}),
	}
}

func InitRedis() {
	RedisClient = newRedis()
}

func (t *redisClient) GetValue(key string) (string, error) {
	value, err := t.client.Get(t.context, key).Result()

	if errors.Is(err, redis.Nil) {
		return "", err
	}

	if errors.Is(err, redis.TxFailedErr) {
		logger.Log().Error(err)
		return "", err
	}

	return value, nil
}

func (t *redisClient) SetValue(key string, value string, duration int) error {
	err := t.client.Set(t.context, key, value, time.Second*time.Duration(duration)).Err()

	if errors.Is(err, redis.Nil) {
		logger.Log().Error(err)
		return err
	}

	if errors.Is(err, redis.TxFailedErr) {
		logger.Log().Error(err)
		return err
	}

	return nil
}
