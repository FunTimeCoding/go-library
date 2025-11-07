package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func Customer() {
	j := internal.Jira()

	for _, i := range j.CustomerIssues(true) {
		fmt.Println(i.Format(option.Color))
	}

	desks := j.Desks()

	for _, e := range desks.Values {
		fmt.Printf("Desk: %+v\n", e)
		types := j.RequestTypes(strings.ToIntegerStrict(e.ID), 0)

		for _, t := range types.Values {
			fmt.Printf("  Type: %+v\n", t)
		}
	}

	i := j.CreateCustomerIssue(
		1,
		4,
		"Test software request",
		"Requesting the software",
	)
	fmt.Printf("Issue: %+v\n", i)
}
