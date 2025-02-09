package main

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/bubbletea"
	"github.com/funtimecoding/go-library/pkg/bubbletea/model/monitor"
	"github.com/funtimecoding/go-library/pkg/bubbletea/table/item"
)

func main() {
	bubbletea.RunAlternative(monitor.New(item.New([]table.Row{})), true)
}
