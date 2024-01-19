package project

func ReplaceDelveVersion(
	content string,
	v string,
) string {
	return ReplaceVersionByPrefix(
		content,
		v,
		`github\.com/go-delve/delve/cmd/dlv@v`,
	)
}
