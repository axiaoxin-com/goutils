package goutils

import (
	"testing"
)

func TestNewRedisClient(t *testing.T) {
	r, err := NewRedisClient("localhost:6379", "", 0)
	if err != nil {
		t.Error("new redis client return error:", err)
	}
	if r == nil {
		t.Error("new a nil redis client")
	}
}
