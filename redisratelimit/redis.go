package redisratelimit

import (
	"time"

	"github.com/go-redis/redis"
)

// Config configuration of the redis client
type Config struct {
	Address            string
	Password           string
	DB                 int
	DialTimeOutSeconds int
	PoolSize           int
}

// Client is the generic wrapper for a service over redis
type Client struct {
	rdb           *redis.Client
	LimitInterval int
	MaxQuotas     int
}

// NewClient handles the general set up of the client
// it creates a db instance for redis
func NewClient(config Config) *Client {
	client := Client{
		rdb: redis.NewClient(&redis.Options{
			Addr:        config.Address,
			Password:    config.Password,
			DB:          config.DB,
			DialTimeout: time.Duration(config.DialTimeOutSeconds) * time.Second,
			PoolSize:    config.PoolSize,
		}),
	}

	return &client
}

func (c *Client) flushAll() {
	c.rdb.FlushAll()
}
