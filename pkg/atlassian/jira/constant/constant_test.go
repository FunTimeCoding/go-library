package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.Integer(t, 100, SearchLimit)
	assert.String(t, "Assignee", AssigneeName)
	assert.String(t, "Attachment", AttachmentName)
	assert.String(t, "Labels", LabelsName)
	assert.String(t, "Linked Issues", LinkedIssuesName)
	assert.String(t, "Parent", ParentName)
	assert.String(t, "Reporter", ReporterName)
	assert.String(t, "Rank", RankName)
	assert.String(t, "Flagged", FlaggedName)
	assert.String(t, "Team", TeamName)
	assert.String(t, "Development", DevelopmentName)
}
