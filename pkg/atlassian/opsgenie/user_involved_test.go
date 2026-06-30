package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/override"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/rotation"
	"github.com/funtimecoding/go-library/pkg/strings/upper"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"testing"
)

func TestUserInvolved(t *testing.T) {
	assert.False(
		t,
		UserInvolved(
			upper.Alfa,
			[]*override.Override{{User: upper.Bravo}},
			[]*rotation.Rotation{},
		),
	)
	assert.True(
		t,
		UserInvolved(
			upper.Alfa,
			[]*override.Override{},
			[]*rotation.Rotation{
				{
					Participants: []og.Participant{{Username: upper.Alfa}},
				},
			},
		),
	)
}
