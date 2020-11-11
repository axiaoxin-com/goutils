// 创建 redis 连接对象的函数封装

package goutils

import (
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// NewRedisClient 返回 redis 的客户端连接对象
func NewRedisClient(opt *redis.Options) (*redis.Client, error) {
	r := redis.NewClient(opt)
	_, err := r.Ping(context.Background()).Result()
	return r, err
}

// NewRedisFailoverClient 返回带 sentinel 的 redis 连接对象
func NewRedisFailoverClient(opt *redis.FailoverOptions) (*redis.Client, error) {
	r := redis.NewFailoverClient(opt)
	_, err := r.Ping(context.Background()).Result()
	return r, err
}

// NewRedisClusterClient 返回 redis cluster 的连接对象
func NewRedisClusterClient(opt *redis.ClusterOptions) (*redis.ClusterClient, error) {
	c := redis.NewClusterClient(opt)
	_, err := c.Ping(context.Background()).Result()
	return c, err
}

// RedisInstances 按 which key 保存 redis 客户端实例
var RedisInstances sync.Map

// RedisSentinelInstances 按 which key 保存 redis sentinel 客户端实例
var RedisSentinelInstances sync.Map

// RedisClusterInstances 按 which key 保存 redis cluster 客户端实例
var RedisClusterInstances sync.Map

// RedisClient 根据 viper 配置返回 redis 客户端
func RedisClient(which string) (*redis.Client, error) {

	client, loaded := RedisInstances.Load(which)
	if loaded {
		return client.(*redis.Client), nil
	}
	// client 不存在则新建实例存放到 map 中
	// 注意：这里依赖 viper ，必须在外部先对 viper 配置进行加载
	prefix := "redis." + which
	newClient, err := NewRedisClient(&redis.Options{
		Addr:         viper.GetString(prefix + ".addr"),
		Password:     viper.GetString(prefix + ".password"),
		DB:           viper.GetInt(prefix + ".db"),
		DialTimeout:  viper.GetDuration(prefix+".dial_timeout") * time.Second,
		ReadTimeout:  viper.GetDuration(prefix+".read_timeout") * time.Second,
		WriteTimeout: viper.GetDuration(prefix+".write_timeout") * time.Second,
		PoolSize:     viper.GetInt(prefix + ".pool_size"),
	})
	if err != nil {
		return nil, err
	}
	RedisInstances.Store(which, newClient)
	return newClient, nil
}

// RedisSentinelClient 根据 viper 配置返回 redis 客户端
func RedisSentinelClient(which string) (*redis.Client, error) {
	client, loaded := RedisSentinelInstances.Load(which)
	if loaded {
		return client.(*redis.Client), nil
	}
	// client 不存在则新建实例存放到 map 中
	// 注意：这里依赖 viper ，必须在外部先对 viper 配置进行加载
	prefix := "redis.sentinel." + which
	newClient, err := NewRedisFailoverClient(&redis.FailoverOptions{
		MasterName:    viper.GetString(prefix + ".master_name"),
		SentinelAddrs: viper.GetStringSlice(prefix + ".sentinel_addrs"),
		Password:      viper.GetString(prefix + ".password"),
		DB:            viper.GetInt(prefix + ".db"),
		DialTimeout:   viper.GetDuration(prefix + ".dial_timeout"),
		ReadTimeout:   viper.GetDuration(prefix + ".read_timeout"),
		WriteTimeout:  viper.GetDuration(prefix + ".write_timeout"),
		PoolSize:      viper.GetInt(prefix + ".pool_size"),
	})
	if err != nil {
		return nil, err
	}
	RedisInstances.Store(which, newClient)
	return newClient, nil
}

// RedisClusterClient 根据 viper 配置返回 redis cluster 实例
func RedisClusterClient(which string) (*redis.ClusterClient, error) {
	client, loaded := RedisClusterInstances.Load(which)
	if loaded {
		return client.(*redis.ClusterClient), nil
	}
	// client 不存在则新建实例存放到 map 中
	// 注意：这里依赖 viper ，必须在外部先对 viper 配置进行加载
	prefix := "redis.cluster." + which
	newClient, err := NewRedisClusterClient(&redis.ClusterOptions{
		Addrs:        viper.GetStringSlice(prefix + ".addrs"),
		Password:     viper.GetString(prefix + ".password"),
		ReadTimeout:  viper.GetDuration(prefix + ".read_timeout"),
		WriteTimeout: viper.GetDuration(prefix + ".write_timeout"),
	})
	if err != nil {
		return nil, err
	}
	RedisClusterInstances.Store(which, newClient)
	return newClient, nil
}

// CloseRedisInstances 关闭全部的 redis 连接并重置 RedisInstances
func CloseRedisInstances() {
	RedisInstances.Range(func(k, v interface{}) bool {
		if rdb, ok := v.(*redis.Client); ok {
			rdb.Close()
		}
		return true
	})
	RedisInstances = sync.Map{}
}

// SetRedisInstances 设置 redis 对象到 RedisInstances 中
func SetRedisInstances(which string, rdb *redis.Client) error {
	RedisInstances.Store(which, rdb)
	return nil
}
