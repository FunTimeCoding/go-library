//go:build local

package service

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/integration_test/service_tester"
	"testing"
)

func TestServiceForgetMemory(t *testing.T) {
	o := service_tester.New(t)
	m, e := o.Service.CreateMemory(
		"to forget",
		"content",
		"desc",
		"feedback",
		"test",
	)
	assert.FatalOnError(t, e)
	assert.FatalOnError(t, o.Service.ForgetMemory(m.Identifier, "test"))
	assert.Count(t, 1, o.Indexer.Deleted)
	assert.String(t, "memory/1", o.Indexer.Deleted[0])
	active, e := o.Service.ListMemories(
		"",
		"",
		true,
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, active)
}
