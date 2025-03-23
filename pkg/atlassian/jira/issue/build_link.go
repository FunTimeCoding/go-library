package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue/option"
)

func buildLink(
	o *option.Issue,
	key string,
) string {
	return fmt.Sprintf("%s/browse/%s", o.Locator, key)
}
