package packages

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestFindVersion(t *testing.T) {
	assert.Any(
		t,
		&gitlab.Package{Name: "Alfa", Version: "v1.0.1"},
		FindVersion(
			[]*gitlab.Package{
				{Name: strings.Alfa, Version: "v1.0.0"},
				{Name: strings.Alfa, Version: "v1.0.2"},
				{Name: strings.Alfa, Version: "v1.0.1"},
			},
			"v1.0.1",
		),
	)
}
