package event

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	events "k8s.io/api/events/v1"
	meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"testing"
)

func TestEvent(t *testing.T) {
	assert.True(
		t,
		New(
			&events.Event{ObjectMeta: meta.ObjectMeta{Name: strings.Alfa}},
			"",
		) != nil,
	)
}
