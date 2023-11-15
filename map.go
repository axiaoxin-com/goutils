package goutils

import (
	"sort"
)

// MapIntIntSortedKeys 返回map[int]int{}按value排序的keys
func MapIntIntSortedKeys(m map[int]int, desc bool) []int {
	keys := make([]int, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if m[keys[i]] == m[keys[j]] { // 如果值相同，就按照key的大小排序
			return keys[i] < keys[j]
		}
		if desc {
			return m[keys[i]] > m[keys[j]]
		}
		return m[keys[i]] < m[keys[j]]
	})

	return keys
}

// MapStrIntSortedKeys 返回map[string]int{}按value排序的keys
func MapStrIntSortedKeys(m map[string]int, desc bool) []string {
	keys := make([]string, 0, len(m))

	for key := range m {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		if m[keys[i]] == m[keys[j]] { // 如果值相同，就按照key的字典序排序
			return keys[i] < keys[j]
		}
		if desc {
			return m[keys[i]] >= m[keys[j]]
		}
		return m[keys[i]] < m[keys[j]]
	})

	return keys
}
