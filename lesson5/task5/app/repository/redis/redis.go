package redis

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisDB interface {
	Get(h string, target *interface{}) error
	Set(h string, target interface{}) error
}

type Redis struct {
	ctx   context.Context
	redis *redis.Client
}

func NewRedis(r *redis.Client) *Redis {
	return &Redis{
		ctx:   context.Background(),
		redis: r,
	}
}

func (rd *Redis) Get(h string, target *interface{}) error {

	value, err := rd.redis.Get(rd.ctx, h).Result()
	if err != nil {
		return err
	}

	if err := json.Unmarshal([]byte(value), &target); err != nil {
		return err
	}

	return nil
}

func (rd *Redis) Set(h string, target interface{}) error {

	value, err := json.Marshal(target)
	if err != nil {
		return err
	}

	if err := rd.redis.SetNX(rd.ctx, h, value, 10*time.Minute).Err(); err != nil {
		return err
	}

	return nil
}
