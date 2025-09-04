package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func Search() {
	p := environment.Get(constant.ProjectEnvironment)
	j := internal.Jira()
	issues := j.SearchV3(
		"project = %s AND status != %s",
		p,
		constant.Closed,
	)
	fmt.Printf("Count: %d\n", len(issues))

	for _, i := range issues {
		fmt.Printf("Issue: %s\n", i.Key)
	}

	limitedIssues := j.SearchLimitV3(
		5,
		"project = %s AND status != %s",
		p,
		constant.Closed,
	)
	fmt.Printf("Limited Count: %d\n", len(limitedIssues))

	for _, i := range limitedIssues {
		fmt.Printf("Issue: %s\n", i.Key)
	}
}
