package main

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/example_list"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/example_table"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/go-library/pkg/bubbletea/style"
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/example_countries"
)

func main() {
	if false {
		bubbletea.Run(example_list.New())
	}

	if false {
		bubbletea.Run(monitor.New())
	}

	if true {
		t := example_countries.New()
		style.Table(t)
		bubbletea.Run(example_table.New(t))
	}
}
