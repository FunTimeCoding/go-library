package container

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/google/go-github/v84/github"
	"testing"
)

func TestContainer(t *testing.T) {
	c := New(
		&github.Package{
			Name:       new(strings.Alfa),
			Repository: &github.Repository{Name: new(strings.Bravo)},
		},
	)
	c.Raw = nil
	assert.Any(
		t,
		&Container{Name: strings.Alfa, Repository: strings.Bravo},
		c,
	)
}
