//go:build local

package service

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/strings"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/integration_test/service_tester"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/service"
	"testing"
)

func TestTopNodes(t *testing.T) {
	s := service_tester.New(t)
	s.AddNode(strings.Alfa, 2, 8)
	s.AddNodeMetrics(strings.Alfa, "500m", "4Gi")
	result, e := s.Service.TopNodes(context.Background(), "test")
	assert.Nil(t, e)
	assert.Count(t, 1, result)
	assert.String(t, strings.Alfa, result[0].Name)
	assert.String(t, "500m", result[0].CPU)
	assert.String(t, "25%", result[0].CPUPercent)
	assert.String(t, "50%", result[0].MemoryPercent)
}

func TestTopPods(t *testing.T) {
	s := service_tester.New(t)
	s.AddPodMetrics(
		"default", "nginx", []map[string]string{
			{"name": "nginx", "cpu": "100m", "memory": "64Mi"},
		},
	)
	s.AddPodMetrics(
		"default", "redis", []map[string]string{
			{"name": "redis", "cpu": "200m", "memory": "128Mi"},
		},
	)
	result, e := s.Service.TopPods(
		context.Background(), "test", service.TopQuery{Namespace: "default"},
	)
	assert.Nil(t, e)
	assert.Count(t, 2, result)
	assert.String(t, "redis", result[0].Name)
	assert.String(t, "nginx", result[1].Name)
}

func TestTopPodContainers(t *testing.T) {
	s := service_tester.New(t)
	s.AddPodMetrics(
		"default", "multi", []map[string]string{
			{"name": "app", "cpu": "300m", "memory": "256Mi"},
			{"name": "sidecar", "cpu": "50m", "memory": "32Mi"},
		},
	)
	result, e := s.Service.TopPodContainers(
		context.Background(),
		"test",
		service.TopQuery{Namespace: "default"},
	)
	assert.Nil(t, e)
	assert.Count(t, 2, result)
	assert.String(t, "app", result[0].Container)
	assert.String(t, "sidecar", result[1].Container)
}
