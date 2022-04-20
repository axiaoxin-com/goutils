// IP 相关方法封装

package goutils

import (
	"context"
	"fmt"
	"net"
)

// GetLocalIP 获取当前 IP
func GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String(), nil
}

// AnonymousName 返回游客昵称
func AnonymousName(ctx context.Context, serviceid int, ip, ua string) string {
	nickname := "游客"
	h, err := NewHashids(fmt.Sprint(ip, ua), 6, "")
	if err != nil {
		return nickname
	}
	hid, err := h.Encode(int64(serviceid))
	if err != nil {
		return nickname
	}
	nickname = "游客" + hid
	return nickname
}
