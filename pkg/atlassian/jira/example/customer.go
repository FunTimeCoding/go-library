package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
)

func Customer() {
	for _, i := range internal.Jira().CustomerIssues(false) {
		fmt.Println(i.Format(option.Color))
	}
}
