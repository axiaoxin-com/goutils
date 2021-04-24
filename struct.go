package goutils

import (
	"net/url"
	"reflect"
	"strconv"
)

// StructToURLValues 将结构体指针对象转换为 url.Values ， key 为 json tag ， value 为结构体字段值，没有 json tag 则使用字段名称
func StructToURLValues(item interface{}, tag string) (values url.Values) {
	values = url.Values{}

	iv := reflect.ValueOf(item).Elem() // Elem() 则 i 必须传指针，不使用 Elem() 则不传递指针
	it := iv.Type()
	for i := 0; i < iv.NumField(); i++ {
		vf := iv.Field(i)
		tf := it.Field(i)

		k := tf.Tag.Get(tag)
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

// StructToMap 结构体转 Map
func StructToMap(item interface{}, tag string) map[string]interface{} {

	res := map[string]interface{}{}
	if item == nil {
		return res
	}
	v := reflect.TypeOf(item)
	reflectValue := reflect.ValueOf(item)
	reflectValue = reflect.Indirect(reflectValue)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	for i := 0; i < v.NumField(); i++ {
		tag := v.Field(i).Tag.Get(tag)
		field := reflectValue.Field(i).Interface()
		if tag != "" && tag != "-" {
			if v.Field(i).Type.Kind() == reflect.Struct {
				res[tag] = StructToMap(field, tag)
			} else {
				res[tag] = field
			}
		}
	}
	return res
}

// StructTagList 获取结构体 tag 列表
func StructTagList(item interface{}, tag string) []string {
	tags := []string{}
	s := reflect.TypeOf(item).Elem()
	for i := 0; i < s.NumField(); i++ {
		tag := s.Field(i).Tag.Get(tag)
		if tag != "" && tag != "-" {
			tags = append(tags, tag)
		}
	}
	return tags
}
