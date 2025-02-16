package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func main() {
	var result []*item.Item

	for _, i := range sentry.NewEnvironment().AllIssues() {
		result = append(
			result,
			item.New(
				fmt.Sprintf("%s-%s", constant.SentryPrefix, *i.Raw.ID),
				constant.ErrorType,
				i.Title,
				i.Link,
			),
		)
	}

	fmt.Println(notation.Encode(result, true))
}
