package console

import "github.com/muesli/termenv"

func Link(
	link string,
	text string,
	osc8 bool,
) string {
	// https://iterm2.com/feature-reporting/Hyperlinks_in_Terminal_Emulators.html
	// https://iterm2.com/3.2/documentation-escape-codes.html
	// https://en.wikipedia.org/wiki/ANSI_escape_code#OSC

	if osc8 {
		return termenv.Hyperlink(link, text)
	}

	return link
}
