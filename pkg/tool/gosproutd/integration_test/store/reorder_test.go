//go:build local

package store

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/integration_test/store_tester"
	"testing"
)

func TestReorderReverses(t *testing.T) {
	s := store_tester.New(t)
	threeSeeds(s)
	seeds := s.Store.Seeds()
	s.Store.Reorder(
		[]uint{
			seeds[2].Identifier,
			seeds[1].Identifier,
			seeds[0].Identifier,
		},
	)
	after := s.Store.Seeds()
	assert.String(t, "charlie", after[0].Name)
	assert.String(t, "beta", after[1].Name)
	assert.String(t, "alpha", after[2].Name)
	assert.Integer(t, 1, after[0].Position)
	assert.Integer(t, 2, after[1].Position)
	assert.Integer(t, 3, after[2].Position)
}
