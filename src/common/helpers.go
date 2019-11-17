package common

import (
	"net/http"
	"strconv"
)

func GetInt64(r *http.Request, key string) int64 {
	val, _ := strconv.Atoi(r.FormValue(key))
	return int64(val)
}
