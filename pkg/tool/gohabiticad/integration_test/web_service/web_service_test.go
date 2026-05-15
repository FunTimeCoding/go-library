//go:build local

package web_service

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/assert"
	"github.com/funtimecoding/go-library/pkg/habitica/response"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/generated/client"
	"github.com/funtimecoding/go-library/pkg/tool/gohabiticad/integration_test/web_service_tester"
	"net/http"
	"testing"
)

func TestWebService(t *testing.T) {
	o := web_service_tester.New(t)
	defer o.Close()
	c := o.Client
	x := context.Background()
	o.MockClient.AddTask(response.Task{
		ID:   "task-1",
		Text: "Deploy fleet",
		Type: "daily",
	})
	o.MockClient.AddTag(response.Tag{ID: "tag-1", Name: "deploy"})
	tasks, e := c.GetTasksWithResponse(x, &client.GetTasksParams{})
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, tasks.StatusCode())
	assert.NotNil(t, tasks.JSON200)
	assert.Count(t, 1, *tasks.JSON200)
	tags, e := c.GetTagsWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, tags.StatusCode())
	assert.Count(t, 1, *tags.JSON200)
	stats, e := c.GetStatsWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, stats.StatusCode())
	assert.NotNil(t, stats.JSON200)
	score, e := c.ScoreTaskWithResponse(
		x,
		"task-1",
		client.Up,
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, score.StatusCode())
	gear, e := c.GetGearWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, gear.StatusCode())
	cron, e := c.RunCronWithResponse(x)
	assert.FatalOnError(t, e)
	assert.Integer(t, http.StatusOK, cron.StatusCode())
}
