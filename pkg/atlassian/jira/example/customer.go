package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func Customer() {
	j := common.Jira()
	fmt.Printf("Info: %+v\n", j.CustomerInfo())

	if false {
		for _, i := range j.CustomerIssues(true) {
			fmt.Println(i.Format(option.Color))
		}
	}

	if false {
		desks := j.Desks()

		for _, e := range desks.Values {
			fmt.Printf("Desk: %+v\n", e)
			types := j.RequestTypes(strings.ToIntegerStrict(e.ID), 0)

			for _, t := range types.Values {
				fmt.Printf("  Type: %+v\n", t)
			}
		}
	}

	if false {
		i := j.CreateCustomerIssue(
			1,
			4,
			"Test software request",
			"Requesting the software",
		)
		fmt.Printf("Issue: %+v\n", i)
	}
}
