package project

func ReplaceGoFromVersion(
	content string,
	v string,
) string {
	return ReplaceVersionByPrefix(content, v, `FROM golang:`)
}
