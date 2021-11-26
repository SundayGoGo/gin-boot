package utils

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"
)

func GetMD5Encode(data string) string {
	h := md5.New()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

// str 2 int64
func Str2Int64(str string) int64 {
	uid, _ := strconv.ParseInt(str, 10, 64)
	return uid
}
