package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func Watch() {
	for _, i := range internal.Jira().WatchedIssues() {
		fmt.Println(i.Format(option.Color))
	}
}
