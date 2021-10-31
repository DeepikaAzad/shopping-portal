package utils

import (
	"encoding/base64"
	"strconv"
)

func DecodeBase64(in string) string {
	out := make([]byte, base64.StdEncoding.DecodedLen(len(in)))
	n, err := base64.StdEncoding.Decode(out, []byte(in))
	if err != nil {
		return ""
	}
	return string(out[0:n])
}

func ParseUint(str string) uint {
	intval, _ := strconv.ParseUint(str, 10, 64)
	return uint(intval)
}
