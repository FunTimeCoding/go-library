//go:build local

package integration_test

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/integration_test/web_interface_tester"
	"net/url"
	"testing"
)

func TestFilterWebInterface(t *testing.T) {
	o := web_interface_tester.New(t)
	defer o.Close()
	o.PostForm(
		"/add",
		url.Values{
			"action":      {"restart"},
			"user":        {"alice"},
			"system":      {"worker1"},
			"service":     {"nginx"},
			"description": {"test"},
		},
	)
	o.PostForm(
		"/add",
		url.Values{
			"action":      {"backup"},
			"user":        {"bob"},
			"system":      {"worker2"},
			"service":     {"nginx"},
			"description": {"test"},
		},
	)
	b := o.Get("/entries?system=worker1")
	assert.StringContains(t, "restart", b)
	assert.StringContains(t, "worker1", b)
	assert.StringContains(
		t,
		"backup",
		o.Get(fmt.Sprintf("/entries?user=%s", "bob")),
	)
	assert.StringContains(
		t,
		"restart",
		o.Get(fmt.Sprintf("/entries?user=%s", "alice")),
	)
	assert.StringContains(
		t,
		"No entries found",
		o.Get("/entries?system=nonexistent"),
	)
}
