package score_colorer

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/console/score_colorer/score_fixture"
	"testing"
)

func TestGreenRed(t *testing.T) {
	assert.True(t, GreenRed(score_fixture.New(1)) != nil)
}
