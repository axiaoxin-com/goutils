package goutils

import "testing"

func TestURLKey(t *testing.T) {
	k := URLKey("prefix", "http://abc.com/x?a=b&c=d")
	t.Log(k)
}
