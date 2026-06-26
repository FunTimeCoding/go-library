package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
)

func SetStatus() {
	if false {
		c := confluence.NewEnvironment()
		result, e := c.DraftOverlay("160497673")

		if e != nil {
			fmt.Printf("error: %v\n", e)

			return
		}

		fmt.Printf(
			"status=%s body=%s\n",
			result.Raw.Status,
			page.ToMarkdown(result.Raw.Body.Storage.Value),
		)
	}
}
