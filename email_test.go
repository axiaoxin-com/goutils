package goutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsEmailValid(t *testing.T) {
	b := IsEmailValid("Hello@qq.com")
	require.True(t, b)
}
