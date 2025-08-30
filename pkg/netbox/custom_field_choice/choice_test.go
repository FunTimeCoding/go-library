package custom_field_choice

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/netbox-community/go-netbox/v4"
	"testing"
)

func TestChoice(t *testing.T) {
	assert.NotNil(t, New(&netbox.CustomFieldChoiceSet{}))
}
