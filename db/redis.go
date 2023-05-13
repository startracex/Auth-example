package db

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

func Connect() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr: "",
		DB:   0,
	})
	return client
}

type RedisInstance struct {
	*redis.Client
}

func (r *RedisInstance) Connect(
	addr string,
	db int,
) {
	r.Client = redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   db,
	})
}

func (r *RedisInstance) Close() error {
	return r.Client.Close()
}

func (r *RedisInstance) Get(key string) (string, error) {
	return r.Client.Get(context.TODO(), key).Result()
}

func (r *RedisInstance) Set(key string, value any, expiration ...int) error {
	var exp int
	if len(expiration) > 0 {
		exp = expiration[0]
	}
	return r.Client.Set(context.TODO(), key, value, time.Duration(exp)*time.Second).Err()
}

func (r *RedisInstance) Del(keys ...string) {
	r.Client.Del(context.TODO(), keys...)
}
