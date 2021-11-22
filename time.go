// 时间相关方法

package goutils

import (
	"fmt"
	"time"
)

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

// GetLatestTradingDay 返回最近一个交易日string类型日期：YYYY-mm-dd
func GetLatestTradingDay() string {
	today := time.Now()
	holidays := ChinaHolidays[fmt.Sprint(today.Year())]
	day := today
	for {
		date := day.Format("2006-01-02")
		if holidays[date] {
			day = day.AddDate(0, 0, -1)
		} else {
			break
		}
	}
	weekday := day.Weekday()
	switch weekday {
	case time.Saturday:
		day = day.AddDate(0, 0, -1)
	case time.Sunday:
		day = day.AddDate(0, 0, -2)
	}
	return day.Format("2006-01-02")
}

// IsTradingDay 返回当期是否为交易日
func IsTradingDay() bool {
	today := time.Now()
	weekday := today.Weekday()
	if weekday == time.Saturday || weekday == time.Sunday {
		return false
	}
	holidays := ChinaHolidays[fmt.Sprint(today.Year())]
	if holidays[today.Format("2006-01-02")] {
		return false
	}
	return true
}
