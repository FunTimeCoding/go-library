package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "Private-Token", PrivateTokenHeader)
	assert.String(t, "CI_API_V4_URL", InterfaceLocator)
	assert.String(t, "CI_PROJECT_ID", ProjectIdentifier)
	assert.String(t, "CI_COMMIT_TAG", CommitTag)
	assert.String(t, "CI_JOB_TOKEN", JobToken)
}
