package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func main() {
	var items []*item.Item
	items = append(
		items,
		item.New(
			strings.Alfa,
			constant.ErrorType,
			"Example detail 1",
			"https://example.com/1",
		),
	)
	items = append(
		items,
		item.New(
			strings.Bravo,
			constant.ErrorType,
			"Example detail 2",
			"https://example.com/2",
		),
	)
	items = append(
		items,
		item.New(
			strings.Charlie,
			constant.ErrorType,
			"Example detail 3",
			"https://example.com/3",
		),
	)
	fmt.Println(notation.Encode(items, true))
}
