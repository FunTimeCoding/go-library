package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func Search() {
	p := environment.Required(constant.DefaultProjectNameEnvironment)
	j := common.Jira()
	f := constant.Format
	searchAndy(j, p, f)
	searchOwn(j, p)
	searchOwnFull(j, p, f)
}

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

func searchOwn(
	j *jira.Client,
	p string,
) {
	if true {
		fmt.Println("SearchV3")
		issues := j.SearchV3(
			"project = %s AND status != %s",
			p,
			constant.Closed,
		)
		fmt.Printf("  Count: %d\n", len(issues))

		for _, i := range issues {
			fmt.Printf("  Issue: %s\n", i.Key)
		}
	}

	if true {
		fmt.Println("SearchLimitV3")
		issues := j.SearchLimitV3(
			5,
			"project = %s AND status != %s",
			p,
			constant.Closed,
		)
		fmt.Printf("  Count: %d\n", len(issues))

		for _, i := range issues {
			fmt.Printf("  Issue: %s\n", i.Key)
		}
	}
}

func searchOwnFull(
	j *jira.Client,
	p string,
	f *option.Format,
) {
	if true {
		fmt.Println("SearchFull")
		issues := j.SearchFull(
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
		fmt.Println("SearchLimitFull")
		issues := j.SearchLimitFull(
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
