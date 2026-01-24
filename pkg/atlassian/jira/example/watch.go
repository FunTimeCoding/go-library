package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/console/status/option"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func Watch() {
	for _, i := range common.Jira().WatchedIssues() {
		fmt.Println(i.Format(option.Color))
	}
}
