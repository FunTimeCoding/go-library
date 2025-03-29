package team_map

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"
	"testing"
)

func TestMap(t *testing.T) {
	assert.True(t, New([]*team.Team{}) != nil)
}
