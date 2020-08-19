// 字符串相关方法封装

package goutils

import "strings"

// RemoveAllWhitespace 删除字符串中所有的空白符
func RemoveAllWhitespace(s string) string {
	s = strings.Replace(s, "\t", "", -1)
	s = strings.Replace(s, "\n", "", -1)
	s = strings.Replace(s, " ", "", -1)
	return s
}
