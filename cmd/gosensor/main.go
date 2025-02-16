package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/constant"
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func main() {
	var result []*item.Item
	result = append(
		result,
		item.New(
			fmt.Sprintf("%s-%d", constant.ExamplePrefix, 1),
			constant.ErrorType,
			strings.Alfa,
			"https://example.org/1",
		),
	)
	result = append(
		result,
		item.New(
			fmt.Sprintf("%s-%d", constant.ExamplePrefix, 2),
			constant.ErrorType,
			strings.Bravo,
			"https://example.org/2",
		),
	)
	result = append(
		result,
		item.New(
			fmt.Sprintf("%s-%d", constant.ExamplePrefix, 3),
			constant.ErrorType,
			strings.Charlie,
			"https://example.org/3",
		),
	)
	fmt.Println(notation.Encode(result, true))
}
