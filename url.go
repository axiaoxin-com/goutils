package goutils

import (
	"bytes"
	"crypto/sha1"
	"io"
	"net/url"
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
