package page

import "regexp"

var (
	richTextMacroPattern = regexp.MustCompile(
		`<ac:structured-macro[^>]*ac:name="([^"]+)"[^>]*><ac:rich-text-body>([\s\S]*?)</ac:rich-text-body></ac:structured-macro>`,
	)
	plainTextMacroPattern = regexp.MustCompile(
		`<ac:structured-macro[^>]*ac:name="([^"]+)"[^>]*><ac:plain-text-body><!\[CDATA\[(.*?)]]></ac:plain-text-body></ac:structured-macro>`,
	)
	macroCommentPattern = regexp.MustCompile(
		`<!-- ac:(\w+) -->\n([\s\S]*?)<!-- /ac:\w+ -->`,
	)
)
