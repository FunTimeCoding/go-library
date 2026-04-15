package page

func ToStorage(markdown string) string {
	if macroCommentPattern.MatchString(markdown) {
		return markersToMacros(markdown)
	}

	return markdownToHTML(markdown)
}
