package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console"
	"github.com/funtimecoding/go-library/pkg/text"
	"github.com/funtimecoding/go-library/pkg/text/option"
	"github.com/funtimecoding/go-library/pkg/time"
	"slices"
)

func (i *Issue) printComments() {
	if i.Raw.Fields == nil {
		return
	}

	if i.Raw.Fields.Comments == nil {
		return
	}

	for _, element := range i.Raw.Fields.Comments.Comments {
		if slices.Contains(i.CommentNameFilter, element.Author.Name) {
			continue
		}

		fmt.Printf(
			"  Comment: %s | %s | %s\n",
			element.Author.Name,
			CommentTime(element).Format(time.DateMinute),
			console.Magenta(
				"%s",
				text.OptimizeWhitespace(element.Body, option.Compact()),
			),
		)
	}
}
