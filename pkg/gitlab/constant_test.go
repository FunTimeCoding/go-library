package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "Private-Token", constant.PrivateTokenHeader)
	assert.String(t, "Job-Token", constant.JobTokenHeader)
	assert.String(t, "CI_API_V4_URL", constant.InterfaceLocator)
	assert.String(t, "CI_PROJECT_ID", constant.ProjectIdentifier)
	assert.String(t, "CI_COMMIT_TAG", constant.CommitTag)
	assert.String(t, "CI_JOB_TOKEN", constant.JobToken)
	assert.String(t, "asc", constant.Ascending)
	assert.String(t, "failed", constant.Failed)
	assert.String(t, "success", constant.Success)
}
