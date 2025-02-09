package main

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/example_list"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/example_table"
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/example_country"
)

func main() {
	if false {
		bubbletea.Run(example_list.New())
	}

	if true {
		bubbletea.Run(example_table.New(example_country.New()))
	}
}
