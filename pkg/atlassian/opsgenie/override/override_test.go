package override

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
	"testing"
)

func TestOverride(t *testing.T) {
	assert.NotNil(t, New(&schedule.ScheduleOverride{}))
}
