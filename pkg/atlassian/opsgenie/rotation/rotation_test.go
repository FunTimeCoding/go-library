package rotation

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"testing"
)

func TestRotation(t *testing.T) {
	assert.True(t, New(&schedule.Rotation{}) != nil)
}
