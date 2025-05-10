package unit

import (
	"github.com/funtimecoding/go-library/pkg/strings/separator"
	"strings"
)

func FromList(line string) *Unit {
	fields := strings.Fields(line)

	return &Unit{
		Name:        fields[0],
		Load:        fields[1],
		Active:      fields[2],
		Sub:         fields[3],
		Description: strings.Join(fields[4:], separator.Space),
	}
}
