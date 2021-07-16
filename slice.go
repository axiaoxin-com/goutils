// slice 相关通用方法

package goutils

import "sort"

// RemoveStringSliceItemByIndex 根据下标删除 string slice 中的元素
func RemoveStringSliceItemByIndex(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

// IsEqualStringSlice 判断两个 string slice 是否相同
func IsEqualStringSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, c := range s1 {
		if c != s2[i] {
			return false
		}
	}
	return true
}

// ChunkFloat64Slice float64 slice 按指定大小进行切块
func ChunkFloat64Slice(data []float64, chunkSize int) [][]float64 {
	result := [][]float64{}
	dataLen := len(data)
	for i := 0; i < dataLen; i += chunkSize {
		end := i + chunkSize
		if end > dataLen {
			end = dataLen
		}
		result = append(result, data[i:end])
	}
	return result
}

// IsStrInSlice 判断字符串是否在给定的字符串列表中
func IsStrInSlice(i string, s []string) bool {
	if !sort.StringsAreSorted(s) {
		sort.Strings(s)
	}
	index := sort.SearchStrings(s, i)
	return (index < len(s) && s[index] == i)
}

// IsIntInSlice 判断字符串是否在给定的字符串列表中
func IsIntInSlice(i int, s []int) bool {
	if !sort.IntsAreSorted(s) {
		sort.Ints(s)
	}
	index := sort.SearchInts(s, i)
	return (index < len(s) && s[index] == i)
}

// ReverseFloat64Slice 直接反转 float64 列表，无返回值
func ReverseFloat64Slice(numbers []float64) {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
}

// ReversedFloat64Slice 反转 float64 列表，有返回值
func ReversedFloat64Slice(numbers []float64) []float64 {
	for i, j := 0, len(numbers)-1; i < j; i, j = i+1, j-1 {
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}
