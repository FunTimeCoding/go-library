package jira

import (
	"github.com/andygrunwald/go-jira"
	"log"
)

func (c *Client) MetaProject(key string) *jira.MetaProject {
	result := c.CreateMeta(key).GetProjectWithKey(key)

	if result == nil {
		log.Panicf("project not found: %s", key)
	}

	return result
}
