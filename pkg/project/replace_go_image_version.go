package project

func ReplaceGoImageVersion(
	content string,
	v string,
) string {
	return ReplaceVersionByPrefix(content, v, `image: golang:`)
}
