package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/query"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func main() {
	c := jira.NewEnvironment()
	f := option.ExtendedColor

	for _, i := range query.Open(c) {
		fmt.Println(i.Format(f))
	}
}
