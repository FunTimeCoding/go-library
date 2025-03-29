package team

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/team"
	"testing"
)

func TestTeam(t *testing.T) {
	assert.True(t, New(&team.ListedTeams{}) != nil)
}
