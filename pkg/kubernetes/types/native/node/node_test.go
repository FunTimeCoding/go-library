package node

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	core "k8s.io/api/core/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestNode(t *testing.T) {
	assert.True(
		t,
		New(
			&core.Node{ObjectMeta: meta.ObjectMeta{Name: strings.Alfa}},
			"",
		) != nil,
	)
}
