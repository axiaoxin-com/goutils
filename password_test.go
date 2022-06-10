package goutils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestHashPasswd(t *testing.T) {
	h, err := HashPassword("123")
	require.Nil(t, err)
	t.Log("hashed 123:", h)
	time.Sleep(5 * time.Second)
	pass := CheckPasswordHash("123", h)
	t.Log("check pass:", pass)
}
