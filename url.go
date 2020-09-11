package goutils

import (
	"bytes"
	"crypto/sha1"
	"io"
	"net/url"
	"reflect"
	"strconv"
)

// URLKey 根据 url 生成 key ，默认使用 url escape ，长度超过 200 则使用 sha1 结果
func URLKey(prefix, u string) string {
	key := url.QueryEscape(u)
	if len(key) > 200 {
		h := sha1.New()
		io.WriteString(h, u)
		key = string(h.Sum(nil))
	}
	var buffer bytes.Buffer
	buffer.WriteString(prefix)
	buffer.WriteString(":")
	buffer.WriteString(key)
	return buffer.String()
}

// StructToURLValues 将结构体指针对象转换为 url.Values ， key 为 json tag ， value 为结构体字段值，没有 json tag 则使用字段名称
func StructToURLValues(i interface{}) (values url.Values) {
	values = url.Values{}

	iv := reflect.ValueOf(i).Elem() // Elem() 则 i 必须传指针，不使用 Elem() 则不传递指针
	it := iv.Type()
	for i := 0; i < iv.NumField(); i++ {
		vf := iv.Field(i)
		tf := it.Field(i)

		k := tf.Tag.Get("json")
		if k == "" {
			k = tf.Name
		}

		v := ""
		switch vf.Interface().(type) {
		case int, int8, int16, int32, int64:
			v = strconv.FormatInt(vf.Int(), 10)
		case uint, uint8, uint16, uint32, uint64:
			v = strconv.FormatUint(vf.Uint(), 10)
		case float32:
			v = strconv.FormatFloat(vf.Float(), 'f', 4, 32)
		case float64:
			v = strconv.FormatFloat(vf.Float(), 'f', 4, 64)
		case []byte:
			v = string(vf.Bytes())
		case string:
			v = vf.String()
		}

		values.Set(k, v)
	}
	return
}
