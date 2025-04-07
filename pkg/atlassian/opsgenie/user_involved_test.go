package opsgenie

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/override"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/rotation"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/opsgenie/opsgenie-go-sdk-v2/og"
	"testing"
)

func TestUserInvolved(t *testing.T) {
	assert.False(
		t,
		UserInvolved(
			strings.Alfa,
			[]*override.Override{{User: strings.Bravo}},
			[]*rotation.Rotation{},
		),
	)
	assert.True(
		t,
		UserInvolved(
			strings.Alfa,
			[]*override.Override{},
			[]*rotation.Rotation{
				{
					Participants: []og.Participant{{Username: strings.Alfa}},
				},
			},
		),
	)
}
