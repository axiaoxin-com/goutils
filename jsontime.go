package goutils

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// JSONTime 用于在 json 中自定义时间格式
// json marshal 一个带有 time.Time 字段类型的结构体时，时间格式固定为 RFC3339 格式
// 将 time.Time 类型替换为 JSONTime 类型，可设置时间格式为 JSONTimeFormat 中定义的格式
type JSONTime struct {
	time.Time
}

// JSONTimeFormat 定义 JSONTime 的时间格式
// 该值可以被外部修改为指定的其他格式
var (
	JSONTimeFormat = "2006-01-02 15:04:05"
)

// NewJSONTime 创建 JSONTime 对象
func NewJSONTime(t time.Time) JSONTime {
	return JSONTime{t}
}

// MarshalJSON 使用 JSONTimeFormat 覆盖 time 包中实现的 json.Marshaler 接口
func (t JSONTime) MarshalJSON() ([]byte, error) {
	formatted := fmt.Sprintf("\"%s\"", t.Format(JSONTimeFormat))
	return []byte(formatted), nil
}

// String 实现 String 方法 用于 print
func (t JSONTime) String() string {
	return t.Time.Format(JSONTimeFormat)
}

// Value 在 gorm 中只重写 MarshalJSON 是不够的，只写这个方法会在写数据库的时候会提示 delete_at 字段不存在，需要加上 database/sql 的 Value 和 Scan 方法
func (t JSONTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	if t.Time.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return t.Time, nil
}

// Scan 在 gorm 中只重写 MarshalJSON 是不够的，只写这个方法会在写数据库的时候会提示 delete_at 字段不存在，需要加上 database/sql 的 Value 和 Scan 方法
func (t *JSONTime) Scan(v interface{}) error {
	value, ok := v.(time.Time)
	if ok {
		*t = JSONTime{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
