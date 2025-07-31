package issue

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/check/issue/option"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/monitor"
	item "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
)

func Check(o *option.Issue) {
	elements := collect()

	if o.Notation {
		printNotation(elements, o)

		return
	}

	f := constant.Format

	for _, e := range elements {
		fmt.Println(e.Format(f))
	}

	if len(elements) == 0 {
		monitor.NoRelevant(item.GoJira.Plural)
	}
}
