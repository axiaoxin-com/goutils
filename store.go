// 参考 https://github.com/gin-contrib/cache/blob/master/persistence/

package goutils

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"reflect"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/patrickmn/go-cache"
)

var (
	// ErrStoreSetVal 命中缓存时，将缓存中的值保存到具体对象时发生错误
	ErrStoreSetVal = errors.New("store: not set value")
	// ErrStoreKeyMiss key 不存在
	ErrStoreKeyMiss = errors.New("store: key not fount")
	// ErrStoreKeyExists key 已存在
	ErrStoreKeyExists = errors.New("store: key exists")
)

// Store is the interface of a store backend
type Store interface {
	// Get retrieves an item from the cache. Returns the item or nil, and a bool indicating
	// whether the key was found.
	Get(ctx context.Context, key string, value interface{}) error

	// Set sets an item to the cache, replacing any existing item.
	Set(ctx context.Context, key string, value interface{}, expire time.Duration) error

	// Add adds an item to the cache only if an item doesn't already exist for the given
	// key, or if the existing item has expired. Returns an error otherwise.
	Add(ctx context.Context, key string, value interface{}, expire time.Duration) error

	// Replace sets a new value for the cache key only if it already exists. Returns an
	// error if it does not.
	Replace(ctx context.Context, key string, data interface{}, expire time.Duration) error

	// Delete removes an item from the cache. Does nothing if the key is not in the cache.
	Delete(ctx context.Context, key string) error

	// Increment increments a real number, and returns error if the value is not real
	Increment(ctx context.Context, key string, data uint64) (uint64, error)

	// Decrement decrements a real number, and returns error if the value is not real
	Decrement(ctx context.Context, key string, data uint64) (uint64, error)

	// Flush deletes all items from the cache.
	Flush(ctx context.Context) error
}

// InmemStore represents the store with internel memory
type InmemStore struct {
	cache.Cache
}

// NewInmemStore new a InmemStore
func NewInmemStore(defaultExpire time.Duration, purgeDuration time.Duration) *InmemStore {
	return &InmemStore{
		*cache.New(defaultExpire, purgeDuration),
	}
}

// Get implement Store interface
func (c *InmemStore) Get(ctx context.Context, key string, ptrValue interface{}) error {
	val, found := c.Cache.Get(key)
	if !found {
		return ErrStoreKeyMiss
	}
	// set found val to value pointer
	v := reflect.ValueOf(ptrValue)
	if v.Type().Kind() == reflect.Ptr && v.Elem().CanSet() {
		v.Elem().Set(reflect.ValueOf(val))
		return nil
	}
	return ErrStoreSetVal
}

// Set implement Store interface
// Add an item to the cache, replacing any existing item. If the duration is 0
// (DefaultExpiration), the cache's default expiration time is used. If it is -1
// (NoExpiration), the item never expires.
func (c *InmemStore) Set(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	c.Cache.Set(key, value, expire)
	return nil
}

// Add implement Store interface
// Add an item to the cache only if an item doesn't already exist for the given
// key, or if the existing item has expired. Returns an error otherwise.
func (c *InmemStore) Add(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	return c.Cache.Add(key, value, expire)
}

// Replace implement Store interface
// Set a new value for the cache key only if it already exists, and the existing
// item hasn't expired. Returns an error otherwise.
func (c *InmemStore) Replace(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	return c.Cache.Replace(key, value, expire)
}

// Delete implement Cache interface
func (c *InmemStore) Delete(ctx context.Context, key string) error {
	c.Cache.Delete(key)
	return nil
}

// Increment implement Store interface
func (c *InmemStore) Increment(ctx context.Context, key string, n int64) error {
	err := c.Cache.Increment(key, n)
	return err
}

// Decrement implement Store interface
func (c *InmemStore) Decrement(ctx context.Context, key string, n int64) error {
	err := c.Cache.Decrement(key, n)
	return err
}

// Flush implement Store interface
func (c *InmemStore) Flush(ctx context.Context) error {
	c.Cache.Flush()
	return nil
}

// -------------------------------------------------------------------------------

// RedisStore represents the store with redis
type RedisStore struct {
	*redis.Client
}

// Encode returns a []byte representing the passed value
func (c *RedisStore) Encode(value interface{}) ([]byte, error) {
	if bytes, ok := value.([]byte); ok {
		return bytes, nil
	}

	switch v := reflect.ValueOf(value); v.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return []byte(strconv.FormatInt(v.Int(), 10)), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return []byte(strconv.FormatUint(v.Uint(), 10)), nil
	}

	// 其他类型使用 gob 保存
	var b bytes.Buffer
	encoder := gob.NewEncoder(&b)
	if err := encoder.Encode(value); err != nil {
		return nil, err
	}
	return b.Bytes(), nil
}

// Decode deserialices the passed []byte into a the passed ptr interface{}
func (c *RedisStore) Decode(byt []byte, ptr interface{}) (err error) {
	if bytes, ok := ptr.(*[]byte); ok {
		*bytes = byt
		return nil
	}

	if v := reflect.ValueOf(ptr); v.Kind() == reflect.Ptr {
		switch p := v.Elem(); p.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			var i int64
			i, err = strconv.ParseInt(string(byt), 10, 64)
			if err != nil {
				return err
			}

			p.SetInt(i)
			return nil

		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i uint64
			i, err = strconv.ParseUint(string(byt), 10, 64)
			if err != nil {
				return err
			}

			p.SetUint(i)
			return nil
		}
	}

	b := bytes.NewBuffer(byt)
	decoder := gob.NewDecoder(b)
	if err = decoder.Decode(ptr); err != nil {
		return err
	}
	return nil
}

// NewRedisStore return a RedisStore
func NewRedisStore(rdb *redis.Client) *RedisStore {
	return &RedisStore{
		rdb,
	}
}

// Get implement Store interface
func (c *RedisStore) Get(ctx context.Context, key string, ptrValue interface{}) error {
	raw, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return ErrStoreKeyMiss
	} else if err != nil {
		return err
	}
	return c.Decode([]byte(raw), ptrValue)
}

// Set implement Store interface
func (c *RedisStore) Set(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	v, err := c.Encode(value)
	if err != nil {
		return err
	}
	return c.Client.Set(ctx, key, string(v), expire).Err()
}

// Add implement Store interface
func (c *RedisStore) Add(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	_, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		v, e := c.Encode(value)
		if e != nil {
			return e
		}
		return c.Client.Set(ctx, key, string(v), expire).Err()
	} else if err == nil {
		return ErrStoreKeyExists
	}
	return err
}

// Replace implement Store interface
func (c *RedisStore) Replace(ctx context.Context, key string, value interface{}, expire time.Duration) error {
	_, err := c.Client.Get(ctx, key).Result()
	if err == redis.Nil {
		return ErrStoreKeyMiss
	} else if err != nil {
		return err
	}
	v, e := c.Encode(value)
	if e != nil {
		return e
	}
	return c.Client.Set(ctx, key, string(v), expire).Err()
}

// Delete implement Store interface
func (c *RedisStore) Delete(ctx context.Context, key string) error {
	return c.Client.Del(ctx, key).Err()
}

// Increment implement Store interface
func (c *RedisStore) Increment(ctx context.Context, key string, n int64) error {
	return c.Client.IncrBy(ctx, key, n).Err()
}

// Decrement implement Store interface
func (c *RedisStore) Decrement(ctx context.Context, key string, n int64) error {
	return c.Client.DecrBy(ctx, key, n).Err()
}

// Flush implement Store interface
func (c *RedisStore) Flush(ctx context.Context) error {
	return c.Client.FlushDB(ctx).Err()
}
