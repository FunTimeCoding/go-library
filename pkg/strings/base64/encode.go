package base64

import "encoding/base64"

func Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}
