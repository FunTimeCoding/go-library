package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/internal"
	"github.com/funtimecoding/go-library/pkg/system/environment"
)

func CustomValue() {
	issue := environment.Get("JIRA_TEST_ISSUE", 1)
	field := environment.Get("JIRA_TEST_FIELD", 1)
	fmt.Printf(
		"Field value: %s\n",
		internal.Jira().SetVerbose(true).Issue(issue).CustomValue(field),
	)
}
