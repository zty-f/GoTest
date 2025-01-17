package http_utils

import (
	"net/http"
	"strings"
	"time"
)

// SetCookie
// go 官方的setcookie会默认去掉domain前边的.，可能有一些浏览器的请求无法携带cookie
func SetCookie(response http.ResponseWriter, name string, value string, maxAge int64, path string, domain string, secure bool, httpOnly bool) {
	expires := time.Unix(time.Now().Unix()+maxAge, 0)
	var cookieStr = (&http.Cookie{
		Name:     name,
		Value:    value,
		Expires:  expires,
		Path:     "/",
		Domain:   domain,
		Secure:   secure,
		HttpOnly: httpOnly,
	}).String()
	if len(domain) > 0 && domain[0] == '.' {
		cookieStr = strings.Replace(cookieStr, "; Domain=", "; Domain=.", 1)
	}
	response.Header().Add("Set-Cookie", cookieStr)
}
