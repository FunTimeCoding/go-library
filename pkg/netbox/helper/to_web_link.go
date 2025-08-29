package helper

import "strings"

func ToWebLink(v string) string {
	return strings.Replace(v, "/api/", "/", 1)
}
