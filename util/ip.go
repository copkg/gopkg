package util

import (
	"net"
	"net/http"
	"strings"
)

// 获取客户端 IP 地址
func GetClientIP(r *http.Request) string {
	// 从 X-Forwarded-For 头部中提取 IP
	xff := r.Header.Get("X-Forwarded-For")
	if xff != "" {
		// X-Forwarded-For 可能包含多个 IP，以逗号分隔，取第一个
		ips := strings.Split(xff, ",")
		if len(ips[0]) > 0 {
			return strings.TrimSpace(ips[0])
		}
	}

	// 从 X-Real-Ip 头部中提取 IP
	xrip := r.Header.Get("X-Real-IP")
	if xrip != "" {
		return xrip
	}

	// 从 RemoteAddr 中提取 IP
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil {
		return ip
	}

	// 如果以上都不成功，返回 empty
	return ""
}
