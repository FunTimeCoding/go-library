package main

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/monitor/item"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/strings"
)

func main() {
	var items []*item.Item
	items = append(items, item.New(strings.Alfa))
	items = append(items, item.New(strings.Bravo))
	items = append(items, item.New(strings.Charlie))
	fmt.Println(notation.Encode(items, true))
}
