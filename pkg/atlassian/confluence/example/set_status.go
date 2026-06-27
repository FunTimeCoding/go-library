package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page/page_put"
)

func SetStatus() {
	if false {
		c := confluence.NewEnvironment()
		result, e := c.SetPageStatus("6717441", constant.DraftStatus)

		if e != nil {
			fmt.Printf("error: %v\n", e)

			return
		}

		fmt.Printf(
			"status=%s version=%d\n",
			result.Raw.Status,
			result.Raw.Version.Number,
		)
		_ = constant.Page
		_ = page_put.NewWithStatus
	}
}
