package compact

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestMail(t *testing.T) {
	assert.String(t, "alfa", Mail("alfa@bravo"))
}
