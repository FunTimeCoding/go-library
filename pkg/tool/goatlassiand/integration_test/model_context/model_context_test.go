//go:build local

package model_context

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/integration_test/model_context_tester"
	"testing"
)

func TestCreateAndGetPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Test Page",
			"body":              "Hello world",
		},
	)
	assert.StringContains(t, "Test Page", result)
	assert.StringContains(t, "current", result)
	assert.StringContains(t, "Hello world", result)
	getResult := o.Client.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "Test Page", getResult)
}

func TestCreateDraftPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	result := o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft Page",
			"body":              "Draft content",
			"draft":             true,
		},
	)
	assert.StringContains(t, "Draft Page", result)
	assert.StringContains(t, "draft", result)
}

func TestGetDraftPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "My Draft",
			"body":              "Draft body",
			"draft":             true,
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1", "draft": true},
	)
	assert.StringContains(t, "My Draft", result)
	assert.StringContains(t, "draft", result)
}

func TestPublishDraftPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Publish Me",
			"body":              "Will be published",
			"draft":             true,
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", result)
	assert.StringContains(t, "Publish Me", result)
}

func TestDraftOverlay(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Published Page",
			"body":              "Original content",
		},
	)
	o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "draft"},
	)
	overlay := o.Client.MustCallTool(
		constant.ConfluenceGetDraftOverlay,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "draft", overlay)
	published := o.Client.MustCallTool(
		constant.ConfluenceGetPage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "current", published)
}

func TestPublishOverlay(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Overlay Target",
			"body":              "Original",
		},
	)
	o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "draft"},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceSetPageStatus,
		map[string]any{"identifier": "1", "status": "current"},
	)
	assert.StringContains(t, "current", result)
}

func TestDeletePublishedPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Delete Me",
			"body":              "Temporary",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1"},
	)
	assert.StringContains(t, "deleted", result)
}

func TestDeleteDraftPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Draft To Delete",
			"body":              "Will be discarded",
			"draft":             true,
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceDeletePage,
		map[string]any{"identifier": "1", "draft": true},
	)
	assert.StringContains(t, "deleted", result)
}

func TestEditPage(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Edit Target",
			"body":              "Before edit",
		},
	)
	result := o.Client.MustCallTool(
		constant.ConfluenceEditPage,
		map[string]any{
			"identifier": "1",
			"old_text":   "Before edit",
			"new_text":   "After edit",
		},
	)
	assert.StringContains(t, "After edit", result)
}

func TestEditPageNotFound(t *testing.T) {
	o := model_context_tester.New(t)
	defer o.Close()
	o.Client.MustCallTool(
		constant.ConfluenceCreatePage,
		map[string]any{
			"space_identifier":  "1",
			"parent_identifier": "0",
			"title":             "Edit Target",
			"body":              "Some content",
		},
	)
	result := o.Client.MustCallToolError(
		constant.ConfluenceEditPage,
		map[string]any{
			"identifier": "1",
			"old_text":   "nonexistent text",
			"new_text":   "replacement",
		},
	)
	assert.StringContains(t, "not found", result)
}
