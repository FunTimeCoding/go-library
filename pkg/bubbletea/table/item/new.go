package item

import (
	"github.com/charmbracelet/bubbles/table"
	"github.com/funtimecoding/go-library/pkg/bubbletea/style"
)

func New(user bool) *table.Model {
	columns := []table.Column{
		{Title: IdentifierColumn},
		{Title: TypeColumn},
		{Title: DetailColumn},
	}

	if user {
		columns = append(columns, table.Column{Title: UserColumn})
	}

	result := table.New(
		table.WithColumns(columns),
		table.WithFocused(true),
	)
	style.Table(&result)

	return &result
}
