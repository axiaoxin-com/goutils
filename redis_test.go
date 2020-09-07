package goutils

import (
	"context"
	"testing"

	"github.com/go-redis/redis/v8"
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
	defer viper.Reset()
	viper.Set("redis.localhost.addr", "127.0.0.1:6379")
	viper.Set("redis.localhost.password", "")
	viper.Set("redis.localhost.db", 0)
	viper.Set("redis.localhost.dial_timeout", 5)
	viper.Set("redis.localhost.read_timeout", 3)
	viper.Set("redis.localhost.write_timeout", 3)
	viper.Set("redis.localhost.pool_size", 0)
	rdb, err := RedisClient("localhost")
	if err != nil {
		t.Error(err)
	}
	if rdb == nil {
		t.Fatal("get a nil redis client")
	}
	defer CloseRedisInstances()
	if _, err := rdb.Ping(context.TODO()).Result(); err != nil {
		t.Error(err)
	}
	if _, err := RedisClient("localhost"); err != nil {
		t.Error(err)
	}
	viper.Set("redis.unittest.addr", "127.0.0.1:6379")
	viper.Set("redis.unittest.password", "")
	viper.Set("redis.unittest.db", 0)
	viper.Set("redis.unittest.dial_timeout", 5)
	viper.Set("redis.unittest.read_timeout", 3)
	viper.Set("redis.unittest.write_timeout", 3)
	viper.Set("redis.unittest.pool_size", 0)
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
