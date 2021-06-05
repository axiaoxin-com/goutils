package goutils

import "testing"

func TestURLKey(t *testing.T) {
	k := URLKey("prefix", "http://abc.com/x?a=b&c=d")
	t.Log(k)
}

func TestStructToURLValues(t *testing.T) {
	v := StructToURLValues(&struct {
		I int `json:"int_i"`
		S string
	}{666, "testing"}, "json")
	if v.Get("int_i") != "666" || v.Get("S") != "testing" {
		t.Fatalf("convert failed: %+v", v)
	}
	if v.Encode() != "S=testing&int_i=666" {
		t.Fatal("encode error:", v.Encode())
	}
}
