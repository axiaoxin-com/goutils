// 时间相关方法

package goutils

// UnixTimestampTrim 将 Unix 时间戳（秒）以指定秒数进行规整
func UnixTimestampTrim(ts int64, sec int64) int64 {
	return ts - ts%sec
}
