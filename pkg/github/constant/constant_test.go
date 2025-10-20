package constant

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"testing"
)

func TestConstant(t *testing.T) {
	assert.String(t, "X-GitHub-Event", EventHeader)
	assert.String(t, "X-Hub-Signature-256", SignatureHeader)
}
