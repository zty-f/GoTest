package http_utils

import (
	"net/http"
	"strings"
	"testing"
)

func TestSetCookie(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		maxAge   int64
		path     string
		domain   string
		secure   bool
		httpOnly bool
	}{
		{"key1", "value1", 0, "/", "example.com", true, true},
		{"key2", "value2", 10, "/", "", true, false},
		{"key3", "value3", 20, "/test", "example.org", false, true},
		{"key4", "value4", 30, "/", "something.com", false, false},
	}
	for _, test := range tests {
		var response http.ResponseWriter
		SetCookie(response, test.name, test.value, test.maxAge, test.path, test.domain, test.secure, test.httpOnly)
		cookieStr := response.Header().Get("Set-Cookie")
		if strings.Contains(cookieStr, test.name) {
			t.Logf("SetCookie passed, name=%s, value=%s", test.name, test.value)
		} else {
			t.Errorf("SetCookie failed, name=%s, value=%s", test.name, test.value)
		}
	}
}
