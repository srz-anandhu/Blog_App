package e

import (
	"net/http"
	"strconv"
)

func GetErrorMsg(c int) string {
	msg, ok := ErrTypeMap[c]
	if !ok {
		return http.StatusText(c)
	}
	return msg
}

// GetHttpStatusCode used to get status code from code provided
func GetHttpStatusCode(c int) int {
	str := strconv.Itoa(c)
	code := str[:3]
	r, _ := strconv.Atoi(code)
	if r < 100 || r >= 600 {
		return http.StatusInternalServerError
	}
	return r
}
