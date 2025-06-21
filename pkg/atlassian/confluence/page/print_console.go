package page

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func (p *Page) PrintConsole() {
	fmt.Println(
		console.Cyan(
			"%s",
			strings.PrefixMultiline(bodyToMarkdown(p.Raw.Body), "> "),
		),
	)
}
