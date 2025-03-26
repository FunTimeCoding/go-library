package console

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/funtimecoding/go-library/pkg/face"
)

func NewColor(hex string) face.SprintFunction {
	return func(
		format string,
		a ...any,
	) string {
		return lipgloss.NewStyle().Foreground(lipgloss.Color(hex)).Render(
			fmt.Sprintf(
				format,
				a...,
			),
		)
	}
}
