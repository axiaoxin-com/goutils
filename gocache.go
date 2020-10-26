// 内存缓存相关

package goutils

import (
	"errors"
	"reflect"

	"github.com/patrickmn/go-cache"
)

var (
	// ErrGoCacheKeyNotFound go-cache get key 不存在
	ErrGoCacheKeyNotFound = errors.New("cache key not found")
	// ErrGoCacheSetResultFailed go-cache get 转换为结果对象失败
	ErrGoCacheSetResultFailed = errors.New("cache set result failed")
)

// GetGoCache 封装go-cache 的 Get 方法支持直接获取具体类型
func GetGoCache(c *cache.Cache, key string, resultPointer interface{}) error {
	val, found := c.Get(key)
	if !found {
		return ErrGoCacheKeyNotFound
	}

	v := reflect.ValueOf(resultPointer)
	if v.Type().Kind() == reflect.Ptr && v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))
		return nil
	}
	return ErrGoCacheSetResultFailed
}
