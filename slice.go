// slice 相关通用方法

package goutils

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
