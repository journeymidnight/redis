package redis

import (
	"github.com/go-redis/redis/v7"
	"time"
)

const (
	maxRetries         = 10
	dialTimeout        = time.Second
	readTimeout        = time.Second
	writeTimeout       = time.Second
	idleTimeout        = 30 * time.Second
	minIdleConnections = 12
)

func NewRedisClient(addresses []string, password string) redis.UniversalClient {
	options := redis.UniversalOptions{
		MaxRetries:   maxRetries,
		DialTimeout:  dialTimeout,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
		IdleTimeout:  idleTimeout,
		Addrs:        addresses,
		Password:     password,
		MinIdleConns: minIdleConnections,
	}
	return redis.NewUniversalClient(&options)
}
