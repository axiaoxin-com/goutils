package goutils

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type ateststruct struct {
	Name string
}

func (s ateststruct) X() bool {
	return true
}

func TestInmemStore(t *testing.T) {
	store := NewInmemStore(2*time.Second, 1*time.Minute)
	assert.NotNil(t, store)

	value := ""
	err := store.Get(nil, "no-exists", &value)
	assert.NotNil(t, err)

	store.Set(nil, "kkk", "vvv", 0)
	err = store.Get(nil, "kkk", &value)
	assert.Nil(t, err)
	assert.Equal(t, "vvv", value)

	err = store.Add(nil, "k", "v", 0)
	assert.Nil(t, err)
	err = store.Get(nil, "k", &value)
	assert.Nil(t, err)
	assert.Equal(t, "v", value)
	err = store.Add(nil, "k", "v1", 0)
	assert.NotNil(t, err)

	err = store.Replace(nil, "k", "v2", 0)
	assert.Nil(t, err)
	err = store.Get(nil, "k", &value)
	assert.Equal(t, "v2", value)

	v := 0
	store.Set(nil, "i", 1, 0)
	err = store.Get(nil, "i", &v)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	err = store.Increment(nil, "i", 10)
	assert.Nil(t, err)
	err = store.Get(nil, "i", &v)
	assert.Nil(t, err)
	assert.Equal(t, 11, v)

	err = store.Decrement(nil, "i", 5)
	assert.Nil(t, err)
	err = store.Get(nil, "i", &v)
	assert.Nil(t, err)
	assert.Equal(t, 6, v)

	store.Delete(nil, "kkk")
	err = store.Get(nil, "kkk", &value)
	assert.NotNil(t, err)

	err = store.Get(nil, "i", &v)
	assert.Nil(t, err)
	store.Flush(nil)
	err = store.Get(nil, "i", &v)
	assert.NotNil(t, err)
}

func TestResidStore(t *testing.T) {
	ctx := context.Background()
	rdb, err := RedisClient("localhost")
	assert.Nil(t, err)
	store := NewRedisStore(rdb)
	assert.NotNil(t, store)

	value := ""
	err = store.Get(ctx, "no-exists", &value)
	assert.NotNil(t, err)

	err = store.Set(ctx, "kkk", "vvv", 0)
	assert.Nil(t, err)
	err = store.Get(ctx, "kkk", &value)
	assert.Nil(t, err)
	assert.Equal(t, "vvv", value)

	err = store.Add(ctx, "k", "v", 0)
	assert.Nil(t, err)
	err = store.Get(ctx, "k", &value)
	assert.Nil(t, err)
	assert.Equal(t, "v", value)
	err = store.Add(ctx, "k", "v1", 0)
	assert.NotNil(t, err)

	err = store.Replace(ctx, "k", "v2", 0)
	assert.Nil(t, err)
	err = store.Get(ctx, "k", &value)
	assert.Equal(t, "v2", value)

	v := 0
	store.Set(ctx, "i", 1, 0)
	err = store.Get(ctx, "i", &v)
	assert.Nil(t, err)
	assert.Equal(t, 1, v)
	err = store.Increment(ctx, "i", 10)
	assert.Nil(t, err)
	err = store.Get(ctx, "i", &v)
	assert.Nil(t, err)
	assert.Equal(t, 11, v)

	err = store.Decrement(ctx, "i", 5)
	assert.Nil(t, err)
	err = store.Get(ctx, "i", &v)
	assert.Nil(t, err)
	assert.Equal(t, 6, v)

	store.Delete(ctx, "kkk")
	err = store.Get(ctx, "kkk", &value)
	assert.NotNil(t, err)

	err = store.Get(ctx, "i", &v)
	assert.Nil(t, err)
	store.Flush(ctx)
	err = store.Get(ctx, "i", &v)
	assert.NotNil(t, err)

	si := []int{1, 2, 3, 4, 5}
	err = store.Set(ctx, "si", si, 0)
	assert.Nil(t, err)
	nsi := []int{}
	err = store.Get(ctx, "si", &nsi)
	assert.Nil(t, err)
	assert.True(t, reflect.DeepEqual(si, nsi))

	st := ateststruct{
		Name: "xx",
	}
	err = store.Set(ctx, "st", st, 0)
	assert.Nil(t, err)
	nst := ateststruct{}
	err = store.Get(ctx, "st", &nst)
	assert.Nil(t, err)
	assert.True(t, nst.X())
	assert.Equal(t, "xx", nst.Name)
}
