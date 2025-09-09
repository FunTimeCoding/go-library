package packages

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestLatest(t *testing.T) {
	assert.Any(
		t,
		map[string]*gitlab.Package{
			"Alfa":  {Name: "Alfa", Version: "v1.0.2"},
			"Bravo": {Name: "Bravo", Version: "v1.1.0"},
		},
		Latest(
			[]*gitlab.Package{
				{Name: strings.Alfa, Version: "v1.0.0"},
				{Name: strings.Alfa, Version: "v1.0.2"},
				{Name: strings.Alfa, Version: "v1.0.1"},
				{Name: strings.Bravo, Version: "v1.1.0"},
			},
		),
	)
}
