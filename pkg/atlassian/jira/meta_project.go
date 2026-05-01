package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

func (c *Client) MetaProject(key string) (*jira.MetaProject, error) {
	meta, e := c.CreateMeta(key)

	if e != nil {
		return nil, e
	}

	result := meta.GetProjectWithKey(key)

	if result == nil {
		return nil, fmt.Errorf("project not found: %s", key)
	}

	return result, nil
}
