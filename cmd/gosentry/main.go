package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/sentry"
)

func main() {
	var items []*item.Item

	for _, i := range sentry.NewEnvironment().AllIssues() {
		items = append(
			items,
			item.New(
				fmt.Sprintf("%s-%s", constant.SentryPrefix, *i.Raw.ID),
				constant.ErrorType,
				i.Title,
				i.Link,
			),
		)
	}

	fmt.Println(notation.Encode(items, true))
}
