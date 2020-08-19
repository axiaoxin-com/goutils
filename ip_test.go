package goutils

import "testing"

func TestGetLocalIP(t *testing.T) {
	ip, err := GetLocalIP()
	if err != nil {
		t.Fatal("GetLocalIP err:", err)
	}
	if ip == "" {
		t.Fatal("GetLocalIP ip is empty")
	}
}
