package goutils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMapIntIntSortedKeys(t *testing.T) {
	m := map[int]int{
		1: 1, 3: 3, 6: 6, 4: 4, 2: 2,
	}
	keys := MapIntIntSortedKeys(m, false)
	require.Equal(t, []int{1, 2, 3, 4, 6}, keys)
	keys = MapIntIntSortedKeys(m, true)
	require.Equal(t, []int{6, 4, 3, 2, 1}, keys)
}
