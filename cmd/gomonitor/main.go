package main

import (
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/item"
)

func main() {
	bubbletea.Run(monitor.New(item.New()), true)
}
