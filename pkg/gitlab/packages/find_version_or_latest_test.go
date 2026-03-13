package packages

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/strings"
	"gitlab.com/gitlab-org/api/client-go/v2"
	"testing"
)

func TestFindVersionOrLatest(t *testing.T) {
	assert.Any(
		t,
		&gitlab.Package{Name: "Alfa", Version: "v1.0.2"},
		FindVersionOrLatest(
			[]*gitlab.Package{
				{Name: strings.Alfa, Version: "v1.0.0"},
				{Name: strings.Alfa, Version: "v1.0.2"},
				{Name: strings.Alfa, Version: "v1.0.1"},
			},
			constant.LatestVersion,
		),
	)
	assert.Any(
		t,
		&gitlab.Package{Name: "Alfa", Version: "v1.0.1"},
		FindVersionOrLatest(
			[]*gitlab.Package{
				{Name: strings.Alfa, Version: "v1.0.0"},
				{Name: strings.Alfa, Version: "v1.0.2"},
				{Name: strings.Alfa, Version: "v1.0.1"},
			},
			"v1.0.1",
		),
	)
}
