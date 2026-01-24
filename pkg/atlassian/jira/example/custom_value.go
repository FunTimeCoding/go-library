package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/system/environment"
	"github.com/funtimecoding/go-library/pkg/tool/common"
)

func CustomValue() {
	i := environment.Required(constant.TestIssueEnvironment)
	f := environment.Required(constant.TestFieldEnvironment)
	fmt.Printf(
		"Field value: %s\n",
		common.Jira().SetVerbose(true).Issue(i).CustomValue(f),
	)
}
