package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func CustomValue() {
	i := environment.Exit(constant.TestIssueEnvironment)
	f := environment.Exit(constant.TestFieldEnvironment)
	fmt.Printf(
		"Field value: %s\n",
		internal.Jira().SetVerbose(true).Issue(i).CustomValue(f),
	)
}
