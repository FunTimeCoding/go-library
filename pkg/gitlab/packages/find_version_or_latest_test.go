package packages

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	gitlab "gitlab.com/gitlab-org/api/client-go"
	"testing"
)

func TestFindVersionOrLatest(t *testing.T) {
	assert.Any(
		t,
		&gitlab.Package{Name: "Alfa"},
		FindVersionOrLatest(
			[]*gitlab.Package{{Name: strings.Alfa}},
			"0.0.1",
		),
	)
}
