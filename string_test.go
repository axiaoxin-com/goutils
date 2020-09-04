package goutils

import "testing"

func TestRemoveAllWhitespace(t *testing.T) {
	rs := RemoveAllWhitespace(" a\tb\n \n\nc d   e ")
	if rs != "abcde" {
		t.Fatal("RemoveAllWhiteSpace error:", rs)
	}
}

func TestReverseString(t *testing.T) {
	s := ReverseString("12345")
	if s != "54321" {
		t.Error("reverse error:", s)
	}
}
