package space

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestNames(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		Names([]*Space{{Name: strings.Alfa}, {Name: strings.Bravo}}),
	)
}
