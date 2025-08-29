package age_colorer

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/console/age_colorer/age_fixture"
	"testing"
	"time"
)

func TestGreenRed(t *testing.T) {
	assert.NotNil(t, GreenRed(age_fixture.New(0*time.Hour)))
}
