package jenkins

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Jobs() []*gojenkins.Job {
	if false {
		type contextKey string
		c.context = context.WithValue(
			c.context,
			contextKey("debug"),
			true,
		)
	}

	result, e := c.client.GetAllJobs(c.context)
	errors.PanicOnError(e)

	return result
}

func (c *Client) JobsJSON() {
	exec := gojenkins.Executor{
		Raw:     new(gojenkins.ExecutorResponse),
		Jenkins: c.client,
	}
	_, e := c.client.Requester.GetJSON(
		c.context,
		"/",
		exec.Raw,
		nil,
	)
	errors.PanicOnError(e)
	jobs := make([]*gojenkins.Job, len(exec.Raw.Jobs))
	jobStrings := make([]string, len(exec.Raw.Jobs))

	for i, job := range exec.Raw.Jobs {
		res := new(gojenkins.JobResponse)
		r2, e2 := c.client.Requester.GetJSON(
			c.context,
			"/job/"+job.Name,
			res,
			nil,
		)
		errors.PanicOnError(e2)

		if false {
			// TODO: Body is already read, would have to craft entire request
			jobStrings[i] = web.ReadString(r2)
		}

		if false {
			ji, err := c.client.GetJob(c.context, job.Name)
			errors.PanicOnError(err)
			jobs[i] = ji
		}
	}

	for i, js := range jobStrings {
		fmt.Printf("Job %d: %s\n", i, js)
	}
}
