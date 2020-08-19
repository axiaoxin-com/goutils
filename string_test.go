package goutils

import "testing"

func TestRemoveAllWhitespace(t *testing.T) {
	rs := RemoveAllWhitespace(" a\tb\n \n\nc d   e ")
	if rs != "abcde" {
		t.Fatal("RemoveAllWhiteSpace error:", rs)
	}
}
