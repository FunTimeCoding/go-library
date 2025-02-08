package main

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/example_list"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/example_table"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/go-library/pkg/bubbletea/style"
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/example_country"
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/item"
)

func main() {
	if true {
		t := item.New()
		style.Table(t)
		bubbletea.RunAlternative(monitor.New(t))
	}

	if false {
		bubbletea.Run(example_list.New())
	}

	if false {
		t := example_country.New()
		style.Table(t)
		bubbletea.Run(example_table.New(t))
	}
}
