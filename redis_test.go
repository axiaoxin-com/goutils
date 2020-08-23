package goutils

import (
	"testing"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

func TestNewRedisClient(t *testing.T) {
	rdb, err := NewRedisClient(&redis.Options{})
	if err != nil {
		t.Fatal(err)
	}
	if rdb == nil {
		t.Fatal("new a nil redis client")
	}
	rdb.Close()
}

func TestRedisClient(t *testing.T) {
	viper.SetDefault("redis.localhost.addr", "127.0.0.1:6379")
	viper.SetDefault("redis.localhost.password", "")
	viper.SetDefault("redis.localhost.db", 0)
	viper.SetDefault("redis.localhost.dial_timeout", 5)
	viper.SetDefault("redis.localhost.read_timeout", 3)
	viper.SetDefault("redis.localhost.write_timeout", 3)
	viper.SetDefault("redis.localhost.pool_size", 0)
	rdb, err := RedisClient("localhost")
	if err != nil {
		t.Error(err)
	}
	if rdb == nil {
		t.Fatal("get a nil redis client")
	}
	defer CloseRedisInstances()
	if _, err := rdb.Ping().Result(); err != nil {
		t.Error(err)
	}
	if _, err := RedisClient("localhost"); err != nil {
		t.Error(err)
	}
	viper.SetDefault("redis.unittest.addr", "127.0.0.1:6379")
	viper.SetDefault("redis.unittest.password", "")
	viper.SetDefault("redis.unittest.db", 0)
	viper.SetDefault("redis.unittest.dial_timeout", 5)
	viper.SetDefault("redis.unittest.read_timeout", 3)
	viper.SetDefault("redis.unittest.write_timeout", 3)
	viper.SetDefault("redis.unittest.pool_size", 0)
	rdb, err = RedisClient("unittest")
	if err != nil {
		t.Error(err)
	}
	if rdb == nil {
		t.Fatal("get a nil redis client")
	}
	instanceCount := 0
	RedisInstances.Range(func(k, v interface{}) bool {
		instanceCount++
		return true
	})
	if instanceCount != 2 {
		t.Error("instanceCount != 2, ", instanceCount)
	}
}
