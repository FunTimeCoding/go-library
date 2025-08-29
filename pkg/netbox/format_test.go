package netbox

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormat(t *testing.T) {
	fixture := []string{"a", "b"}
	assert.Equal(t, "[a b]", fmt.Sprintf("%v", fixture))
	assert.Equal(t, "[a b]", fmt.Sprintf("%+v", fixture))
	assert.Equal(
		t,
		`[]string{"a", "b"}`,
		fmt.Sprintf("%#v", fixture),
	)
}
