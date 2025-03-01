package item

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/bubbletea/style"
)

func New() *table.Model {
	result := table.New(
		table.WithColumns(
			[]table.Column{
				{Title: IdentifierColumn},
				{Title: DetailColumn},
				{Title: UserColumn},
			},
		),
		table.WithFocused(true),
	)
	style.Table(&result)

	return &result
}
