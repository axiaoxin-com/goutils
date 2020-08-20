// 创建 redis 连接对象的函数封装

package goutils

import (
	"github.com/go-redis/redis"
)

// NewRedisClient 返回 redis 的客户端连接对象
func NewRedisClient(addr string, password string, dbindex int) (*redis.Client, error) {
	r := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       dbindex,
	})
	_, err := r.Ping().Result()
	return r, err
}

// NewRedisSentinel 返回 redis sentinel 的连接对象
func NewRedisSentinel(master string, addrs []string, password string, dbindex int) (*redis.Client, error) {
	r := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    master,
		SentinelAddrs: addrs,
		Password:      password,
		DB:            dbindex,
	})
	_, err := r.Ping().Result()
	return r, err
}

// NewRedisCluster 返回 redis cluster 的连接对象
func NewRedisCluster(addrs []string, password string) (*redis.ClusterClient, error) {
	c := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:    addrs,
		Password: password,
	})
	_, err := c.Ping().Result()
	return c, err
}
