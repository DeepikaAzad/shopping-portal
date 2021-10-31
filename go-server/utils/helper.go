package utils

import "encoding/base64"

func DecodeBase64(in string) string {
	out := make([]byte, base64.StdEncoding.DecodedLen(len(in)))
	n, err := base64.StdEncoding.Decode(out, []byte(in))
	if err != nil {
		return ""
	}
	return string(out[0:n])
}
