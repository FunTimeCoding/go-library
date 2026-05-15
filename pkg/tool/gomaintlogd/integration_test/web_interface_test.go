//go:build local

package integration_test

import (
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/integration_test/web_interface_tester"
	"net/http"
	"net/url"
	"testing"
)

func TestWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	o.AssertStatus("/", http.StatusOK)
	o.AssertStatus("/entries", http.StatusOK)
	o.AssertStatus("/add", http.StatusOK)
	assert.StringContains(t, "No entries found", o.Get("/"))
	addBody := o.PostForm(
		"/add",
		url.Values{
			"action":      {"restarted web server"},
			"user":        {"jdoe"},
			"system":      {"worker1"},
			"service":     {"nginx"},
			"description": {"nginx was unresponsive, restarted"},
		},
	)
	assert.StringContains(t, "Entry added", addBody)
	assert.StringContains(t, "restarted web server", addBody)
	assert.StringContains(t, "restarted web server", o.Get("/"))
	assert.StringContains(t, "jdoe", o.Get("/"))
	assert.StringContains(t, "worker1", o.Get("/entries"))
	detail := o.Get("/entry/detail?id=1")
	assert.StringContains(t, "nginx was unresponsive, restarted", detail)
	assert.StringContains(t, "Edit", detail)
	assert.StringContains(t, "Delete", detail)
	edit := o.Get("/entry/edit?id=1")
	assert.StringContains(t, "restarted web server", edit)
	assert.StringContains(t, "Save", edit)
	editBody := o.PostForm(
		"/entry/edit?id=1",
		url.Values{
			"action":      {"cleared and documented"},
			"user":        {"jdoe"},
			"system":      {"worker1"},
			"service":     {"nginx"},
			"description": {"nginx was unresponsive, restarted and documented"},
			"timestamp":   {"2026-03-19T10:00"},
		},
	)
	assert.StringContains(t, "cleared and documented", editBody)
	o.PostForm("/entry/delete?id=1", nil)
	assert.StringContains(t, "No entries found", o.Get("/"))
}
