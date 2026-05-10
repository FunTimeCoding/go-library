package basic

import (
	"encoding/json"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"
	"github.com/funtimecoding/go-library/pkg/provision/salt/constant"
)

func (c *Client) LookupJob(jid string) (*response.Job, error) {
	b, e := c.Get(fmt.Sprintf("%s/%s", constant.JobsPath, jid))

	if e != nil {
		return nil, e
	}

	var r response.JobDetail

	if f := json.Unmarshal(b, &r); f != nil {
		return nil, f
	}

	if len(r.Details) == 0 {
		return nil, fmt.Errorf("job %s not found", jid)
	}

	job := r.Details[0]

	if job.Error != "" {
		return nil, fmt.Errorf("job %s: %s", jid, job.Error)
	}

	return &job, nil
}
