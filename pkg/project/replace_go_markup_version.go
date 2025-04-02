package project

func ReplaceGoMarkupVersion(
	content string,
	v string,
) string {
	return ReplaceVersionByPrefix(
		content,
		v,
		"          go-version: ",
	)
}
