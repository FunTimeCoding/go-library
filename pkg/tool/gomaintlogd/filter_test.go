package gomaintlogd

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/client"
	"net/http"
	"testing"
)

func TestFilterViaREST(t *testing.T) {
	s, port, l := setup(t)
	defer s.Close()
	defer l.Stop()
	base := fmt.Sprintf("http://localhost:%d", port)
	c, e := client.NewClientWithResponses(base)
	assert.FatalOnError(t, e)
	x := context.Background()

	// Seed: alice/worker1/nginx, bob/worker2/postgres, alice/worker2/nginx
	system1 := "worker1"
	system2 := "worker2"
	service1 := "nginx"
	service2 := "postgres"
	user1 := "alice"
	user2 := "bob"
	desc := "test"

	for _, entry := range []struct {
		action  string
		user    string
		system  *string
		service *string
	}{
		{"restart", user1, &system1, &service1},
		{"backup", user2, &system2, &service2},
		{"deploy", user1, &system2, &service1},
	} {
		_, e = c.PostEntryWithResponse(x, client.PostEntryJSONRequestBody{
			Action:      entry.action,
			User:        entry.user,
			System:      entry.system,
			Service:     entry.service,
			Description: &desc,
		})
		assert.FatalOnError(t, e)
	}

	// No filter → all 3
	all, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{})
	assert.FatalOnError(t, e)
	assert.Count(t, 3, *all.JSON200)

	// Filter by system
	bySys1, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{System: &system1})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *bySys1.JSON200)
	assert.String(t, "restart", (*bySys1.JSON200)[0].Action)

	bySys2, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{System: &system2})
	assert.FatalOnError(t, e)
	assert.Count(t, 2, *bySys2.JSON200)

	// Filter by user
	byAlice, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{User: &user1})
	assert.FatalOnError(t, e)
	assert.Count(t, 2, *byAlice.JSON200)

	byBob, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{User: &user2})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *byBob.JSON200)
	assert.String(t, "backup", (*byBob.JSON200)[0].Action)

	// Filter by service
	bySvc1, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{Service: &service1})
	assert.FatalOnError(t, e)
	assert.Count(t, 2, *bySvc1.JSON200)

	bySvc2, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{Service: &service2})
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *bySvc2.JSON200)

	// Combined: alice on worker2 → only "deploy"
	byAliceSys2, e := c.GetEntriesWithResponse(
		x,
		&client.GetEntriesParams{User: &user1, System: &system2},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 1, *byAliceSys2.JSON200)
	assert.String(t, "deploy", (*byAliceSys2.JSON200)[0].Action)

	// No match
	nobody := "nobody"
	byNobody, e := c.GetEntriesWithResponse(x, &client.GetEntriesParams{User: &nobody})
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, byNobody.StatusCode())
	assert.Count(t, 0, *byNobody.JSON200)
}

func TestFilterViaWeb(t *testing.T) {
	s, port, l := setup(t)
	defer s.Close()
	defer l.Stop()
	base := fmt.Sprintf("http://localhost:%d", port)
	c, e := client.NewClientWithResponses(base)
	assert.FatalOnError(t, e)
	x := context.Background()

	system1 := "worker1"
	system2 := "worker2"
	user1 := "alice"
	user2 := "bob"
	svc := "nginx"
	desc := "test"

	_, e = c.PostEntryWithResponse(x, client.PostEntryJSONRequestBody{
		Action: "restart", User: user1, System: &system1, Service: &svc, Description: &desc,
	})
	assert.FatalOnError(t, e)

	_, e = c.PostEntryWithResponse(x, client.PostEntryJSONRequestBody{
		Action: "backup", User: user2, System: &system2, Service: &svc, Description: &desc,
	})
	assert.FatalOnError(t, e)

	// Filter by system=worker1 → "restart" visible, "backup" not
	body := getBody(t, base+"/entries?system="+system1)
	assert.StringContains(t, "restart", body)
	assert.StringContains(t, "worker1", body) // form value reflected

	// Filter by user=bob → "backup" visible
	body = getBody(t, base+"/entries?user="+user2)
	assert.StringContains(t, "backup", body)

	// Filter by user=alice → "restart" visible
	body = getBody(t, base+"/entries?user="+user1)
	assert.StringContains(t, "restart", body)

	// No match → "No entries found"
	body = getBody(t, base+"/entries?system=nonexistent")
	assert.StringContains(t, "No entries found", body)
}
