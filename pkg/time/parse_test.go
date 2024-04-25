package time

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestParse(t *testing.T) {
	result := Parse("2006-01-02", "2019-01-01")
	assert.Integer(t, 2019, result.Year())
	assert.Integer(t, 1, int(result.Month()))
	assert.Integer(t, 1, result.Day())
}
