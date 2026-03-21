package gomaintlogd

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"net/http"
	"net/url"
	"strings"
	"testing"
)

func TestWebEndpoint(t *testing.T) {
	s, port, l := setup(t)
	defer s.Close()
	defer l.Stop()
	base := fmt.Sprintf("http://localhost:%d", port)
	// Pages render
	assert.HTTPStatus(t, base+"/", http.StatusOK)
	assert.HTTPStatus(t, base+"/entries", http.StatusOK)
	assert.HTTPStatus(t, base+"/add", http.StatusOK)
	// Dashboard shows empty state
	body := getBody(t, base+"/")
	assert.StringContains(t, "No entries found", body)
	// Add entry via form POST
	form := url.Values{
		"action":      {"restarted web server"},
		"user":        {"jdoe"},
		"system":      {"worker1"},
		"service":     {"nginx"},
		"description": {"nginx was unresponsive, restarted"},
	}
	r, e := http.Post(
		base+"/add",
		"application/x-www-form-urlencoded",
		strings.NewReader(form.Encode()),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, r.StatusCode)
	addBody := readBody(t, r)
	assert.StringContains(t, "Entry added", addBody)
	assert.StringContains(t, "restarted web server", addBody)
	// Dashboard now shows the entry
	body = getBody(t, base+"/")
	assert.StringContains(t, "restarted web server", body)
	assert.StringContains(t, "jdoe", body)
	// Entries page shows the entry
	body = getBody(t, base+"/entries")
	assert.StringContains(t, "worker1", body)
	// Detail fragment
	body = getBody(t, base+"/entry/detail?id=1")
	assert.StringContains(t, "nginx was unresponsive, restarted", body)
	assert.StringContains(t, "Edit", body)
	assert.StringContains(t, "Delete", body)
	// Edit form fragment
	body = getBody(t, base+"/entry/edit?id=1")
	assert.StringContains(t, "restarted web server", body)
	assert.StringContains(t, "Save", body)
	// Submit edit via form POST
	editForm := url.Values{
		"action":      {"cleared and documented"},
		"user":        {"jdoe"},
		"system":      {"worker1"},
		"service":     {"nginx"},
		"description": {"nginx was unresponsive, restarted and documented"},
		"timestamp":   {"2026-03-19T10:00"},
	}
	r, e = http.Post(
		base+"/entry/edit?id=1",
		"application/x-www-form-urlencoded",
		strings.NewReader(editForm.Encode()),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, r.StatusCode)
	editBody := readBody(t, r)
	assert.StringContains(t, "cleared and documented", editBody)
	// Delete via POST
	r, e = http.Post(base+"/entry/delete?id=1", "", nil)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, r.StatusCode)
	// Dashboard shows empty again
	body = getBody(t, base+"/")
	assert.StringContains(t, "No entries found", body)
}
