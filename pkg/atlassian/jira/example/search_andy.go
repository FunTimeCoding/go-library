package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func searchAndy(
	j *jira.Client,
	p string,
	f *option.Format,
) {
	if true {
		fmt.Println("Search")
		issues := j.Search(
			"project = %s AND status != %s",
			p,
			constant.Closed,
		)
		fmt.Printf("  Count: %d\n", len(issues))

		for _, i := range issues {
			fmt.Printf("  Issue: %s\n", i.Format(f))
		}
	}

	if true {
		fmt.Println("SearchLimit")
		issues := j.SearchLimit(
			5,
			"project = %s AND status != %s",
			p,
			constant.Closed,
		)
		fmt.Printf("  Count: %d\n", len(issues))

		for _, i := range issues {
			fmt.Printf("  Issue: %s\n", i.Format(f))
		}
	}
}
