package gomaintlogd

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/client"
	"net/http"
	"testing"
)

func TestGeneratedClient(t *testing.T) {
	s, port, l := setup(t)
	defer s.Close()
	defer l.Stop()
	base := fmt.Sprintf("http://localhost:%d", port)
	c, e := client.NewClientWithResponses(base)
	assert.FatalOnError(t, e)
	x := context.Background()

	// Create
	system := "worker1"
	service := "nginx"
	description := "nginx was unresponsive"
	create, e := c.PostEntryWithResponse(
		x,
		client.PostEntryJSONRequestBody{
			Action:      "restarted web server",
			User:        "jdoe",
			System:      &system,
			Service:     &service,
			Description: &description,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, create.StatusCode())
	assert.NotNil(t, create.JSON200)
	assert.String(t, "restarted web server", create.JSON200.Action)
	assert.String(t, "jdoe", create.JSON200.User)
	id := create.JSON200.Id

	// Read
	status, e := c.GetStatusWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, status.StatusCode())
	assert.Integer(t, 1, status.JSON200.TotalEntries)

	entries, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{})
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, entries.StatusCode())
	assert.Count(t, 1, *entries.JSON200)

	// Update
	newAction := "cleared and documented"
	update, e := c.UpdateEntryWithResponse(
		x,
		id,
		client.UpdateEntryJSONRequestBody{
			Action:      newAction,
			User:        "jdoe",
			System:      &system,
			Service:     &service,
			Description: &description,
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, update.StatusCode())
	assert.String(t, newAction, update.JSON200.Action)

	// Read after update
	entries, e = c.GetEntriesWithResponse(x, &client.GetEntriesParams{})
	assert.FatalOnError(t, e)
	assert.String(t, newAction, (*entries.JSON200)[0].Action)

	// Delete
	del, e := c.DeleteEntryWithResponse(x, id)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusNoContent, del.StatusCode())

	// Read after delete
	status, e = c.GetStatusWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, 0, status.JSON200.TotalEntries)
}
