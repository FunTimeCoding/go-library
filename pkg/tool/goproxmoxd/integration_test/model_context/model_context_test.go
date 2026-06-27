//go:build local

package model_context

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/integration_test/model_context_tester"
	"testing"
)

func TestListNodes(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(constant.ListNodes, nil)
	assert.StringContains(t, "test", result)
}

func TestListMachines(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.MockClient.AddMachine("test", 100, "web-server")
	result := o.Client.MustCallTool(
		constant.ListMachines,
		map[string]any{"node": "test"},
	)
	assert.StringContains(t, "web-server", result)
}

func TestCreateMachineCloudInit(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.CreateMachine,
		map[string]any{
			"node":    "test",
			"name":    "ci-test",
			"ci_user": "admin",
		},
	)
	assert.StringContains(t, "created", result)
}

func TestCreateMachineWithDiskImport(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.CreateMachine,
		map[string]any{
			"node":        "test",
			"name":        "import-test",
			"disk_import": "local:import/debian-13.qcow2",
			"disk_size":   64,
			"ci_user":     "admin",
		},
	)
	assert.StringContains(t, "created", result)
}

func TestCreateMachineCDROM(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.CreateMachine,
		map[string]any{
			"node":  "test",
			"name":  "iso-test",
			"cdrom": "local:iso/debian.iso",
		},
	)
	assert.StringContains(t, "created", result)
}

func TestCreateMachineCDROMAndCloudInitError(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallToolError(
		constant.CreateMachine,
		map[string]any{
			"node":    "test",
			"name":    "conflict",
			"cdrom":   "local:iso/debian.iso",
			"ci_user": "admin",
		},
	)
	assert.StringContains(t, "mutually exclusive", result)
}

func TestCloneMachine(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.MockClient.AddMachine("test", 100, "template")
	result := o.Client.MustCallTool(
		constant.CloneMachine,
		map[string]any{
			"identifier": 100,
			"name":       "clone-of-template",
			"node":       "test",
		},
	)
	assert.StringContains(t, "cloned", result)
}

func TestSnippetCRUD(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.CreateSnippet,
		map[string]any{
			"name":    "test.yml",
			"content": "#cloud-config\npackages:\n  - vim\n",
		},
	)
	list := o.Client.MustCallTool(constant.ListSnippets, nil)
	assert.StringContains(t, "test.yml", list)
	get := o.Client.MustCallTool(
		constant.GetSnippet,
		map[string]any{"name": "test.yml"},
	)
	assert.StringContains(t, "cloud-config", get)
	o.Client.MustCallTool(
		constant.DeleteSnippet,
		map[string]any{"name": "test.yml"},
	)
	listAfter := o.Client.MustCallTool(constant.ListSnippets, nil)
	assert.StringNotContains(t, "test.yml", listAfter)
}

func TestListStorages(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.MockClient.AddStorage("test", "local", "dir", "iso,snippets")
	result := o.Client.MustCallTool(
		constant.ListStorages,
		map[string]any{"node": "test"},
	)
	assert.StringContains(t, "local", result)
}

func TestMachineNotFound(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallToolError(
		constant.GetMachine,
		map[string]any{"identifier": 999, "node": "test"},
	)
	assert.StringContains(t, "unable to find", result)
}
