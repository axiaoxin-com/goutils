// 时间相关方法

package goutils

import "time"

// UnixTimestampTrim 将 Unix 时间戳（秒）以指定秒数进行规整
func UnixTimestampTrim(ts int64, sec int64) int64 {
	return ts - ts%sec
}

// StrToTime 字符串时间转 time 对象
func StrToTime(layout, value string, loc ...*time.Location) (time.Time, error) {
	location := time.Local
	if len(loc) > 0 {
		location = loc[0]
	}
	return time.ParseInLocation(layout, value, location)
}
