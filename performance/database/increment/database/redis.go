package database

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

type Redis struct {
	ctx    context.Context
	client *redis.Client
}

func NewRedis(ctx context.Context) *Redis {
	opts := redis.Options{Addr: "localhost:6379", DB: 0}
	rdb := redis.NewClient(&opts)
	redis := new(Redis)
	redis.ctx = ctx
	redis.client = rdb
	return redis
}

func (r *Redis) Increment(value uint16) (uint16, error) {
	result, ok := r.client.IncrBy(r.ctx, "product_1", int64(value)).Result()
	if ok != nil {
		return 0, errors.New("can't increment")
	}
	return uint16(result), nil
}

func (r *Redis) Set(value uint16) (uint16, error) {
	if _, ok := r.client.Set(r.ctx, "product_1", int64(value), 0).Result(); ok != nil {
		return 0, errors.New("can't set")
	}
	return value, nil
}
