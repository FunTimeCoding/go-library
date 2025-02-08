package main

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/example_list"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/example_table"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/example_country"
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/item"
)

func main() {
	if true {
		bubbletea.RunAlternative(monitor.New(item.New()), true)
	}

	if false {
		bubbletea.Run(example_list.New())
	}

	if false {
		bubbletea.Run(example_table.New(example_country.New()))
	}
}
