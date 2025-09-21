package page

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"testing"
)

func TestNames(t *testing.T) {
	assert.Strings(
		t,
		[]string{"Alfa", "Bravo"},
		Names([]*Page{{Name: strings.Alfa}, {Name: strings.Bravo}}),
	)
}
