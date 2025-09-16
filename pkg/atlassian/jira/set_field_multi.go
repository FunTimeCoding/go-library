package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/custom_field_value"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"log"
	"slices"
)

func (c *Client) SetFieldMulti(
	projectKey string,
	issueType string,
	i *jira.Issue,
	fieldName string,
	valueNames []string,
) {
	valid := c.CustomFieldValues(projectKey, issueType, fieldName)

	if len(valid) == 0 {
		log.Panicf(
			"project %s type %s field %s has no values",
			projectKey,
			issueType,
			fieldName,
		)
	}

	var validNames []string

	for _, u := range valid {
		validNames = append(validNames, u.Value)
	}

	var notFound []string

	for _, n := range valueNames {
		if !slices.Contains(validNames, n) {
			notFound = append(notFound, n)
		}
	}

	if len(notFound) > 0 {
		log.Panicf(
			"field %s value(s) not found: %s, valid: %s",
			fieldName,
			join.Comma(notFound),
			join.Comma(validNames),
		)
	}

	var result []custom_field_value.Value

	for _, u := range valid {
		if slices.Contains(valueNames, u.Value) {
			result = append(result, u)
		}
	}

	c.SetField(i, fieldName, result)
}
