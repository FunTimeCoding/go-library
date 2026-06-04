//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/integration_test/service_tester"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/service"
	"testing"
)

func TestMoveUpNotifies(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Sync(
		[]service.DiscoveredFile{
			{Name: "alpha", Path: "alpha.md", ContentHash: "a", Content: "a"},
			{Name: "beta", Path: "beta.md", ContentHash: "b", Content: "b"},
		},
	)
	s.Notifier.Reset()
	s.Service.MoveUp(s.Service.Seeds()[1].Identifier)
	assert.True(t, s.Notifier.Notified > 0)
	assert.String(t, "beta", s.Service.Seeds()[0].Name)
}

func TestSetPositionNotifies(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Sync(
		[]service.DiscoveredFile{
			{Name: "alpha", Path: "alpha.md", ContentHash: "a", Content: "a"},
			{Name: "beta", Path: "beta.md", ContentHash: "b", Content: "b"},
			{
				Name:        "charlie",
				Path:        "charlie.md",
				ContentHash: "c",
				Content:     "c",
			},
		},
	)
	s.Notifier.Reset()
	s.Service.SetPosition(s.Service.Seeds()[2].Identifier, 1)
	assert.True(t, s.Notifier.Notified > 0)
	assert.String(t, "charlie", s.Service.Seeds()[0].Name)
}

func TestReorderNotifies(t *testing.T) {
	s := service_tester.New(t)
	s.Service.Sync(
		[]service.DiscoveredFile{
			{Name: "alpha", Path: "alpha.md", ContentHash: "a", Content: "a"},
			{Name: "beta", Path: "beta.md", ContentHash: "b", Content: "b"},
		},
	)
	seeds := s.Service.Seeds()
	s.Notifier.Reset()
	s.Service.Reorder([]uint{seeds[1].Identifier, seeds[0].Identifier})
	assert.True(t, s.Notifier.Notified > 0)
	assert.String(t, "beta", s.Service.Seeds()[0].Name)
}
