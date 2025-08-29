package issue

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestKeys(t *testing.T) {
	assert.Strings(
		t,
		[]string{"ABC-123", "XYZ-456"},
		Keys("Message with keys ABC-123 and XYZ-456."),
	)
}
