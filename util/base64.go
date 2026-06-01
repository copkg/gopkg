package util

import (
	"regexp"
	"strings"
)

// 安全的base64转标准base64
func UrlSafeBase64ToStd(urlSafeString string) string {
	reg := regexp.MustCompile(`[^A-Za-z0-9+/=]`)
	cleaned := reg.ReplaceAllString(urlSafeString, "")
	//// 2. 将URL安全格式转换为标准格式
	cleaned = strings.ReplaceAll(urlSafeString, "-", "+")
	cleaned = strings.ReplaceAll(cleaned, "_", "/")
	pad := len(cleaned) % 4
	if pad == 0 {
		cleaned += strings.Repeat("=", 4-pad)
	}
	cleaned = strings.Replace(cleaned, " ", "", -1)
	return cleaned
}
