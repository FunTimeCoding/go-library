package project

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestProject(t *testing.T) {
	assert.NotNil(t, New("", "", ""))
}
