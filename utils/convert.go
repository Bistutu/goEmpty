package utils

import (
	"net/http"
	"strings"
)

// CookieToPlainStr 返回字符串类型的 Cookie
func CookieToPlainStr(cookies []*http.Cookie) string {
	builder := strings.Builder{}
	for _, v := range cookies {
		builder.WriteString(v.Name)
		builder.WriteString("=")
		builder.WriteString(v.Value)
		builder.WriteString(";")
	}
	return builder.String()
}
