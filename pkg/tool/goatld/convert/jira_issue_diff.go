package convert

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"strings"
)

func JiraIssueDiff(
	before *issue.Issue,
	after *issue.Issue,
	noDiff bool,
	customFields []string,
) *UpdateResult {
	result := &UpdateResult{Issue: JiraIssue(after)}

	if noDiff {
		return result
	}

	if before.Summary != after.Summary {
		result.Changes = append(
			result.Changes,
			FieldChange{
				Field:  "Summary",
				Before: before.Summary,
				After:  after.Summary,
			},
		)
	}

	if before.Description != after.Description {
		result.Changes = append(
			result.Changes,
			FieldChange{
				Field:  "Description",
				Before: before.Description,
				After:  after.Description,
			},
		)
	}

	if before.Status != after.Status {
		result.Changes = append(
			result.Changes,
			FieldChange{
				Field:  "Status",
				Before: before.Status,
				After:  after.Status,
			},
		)
	}

	if before.Priority != after.Priority {
		result.Changes = append(
			result.Changes,
			FieldChange{
				Field:  "Priority",
				Before: before.Priority,
				After:  after.Priority,
			},
		)
	}

	beforeAssignee := assigneeName(before)
	afterAssignee := assigneeName(after)

	if beforeAssignee != afterAssignee {
		result.Changes = append(
			result.Changes,
			FieldChange{
				Field:  "Assignee",
				Before: beforeAssignee,
				After:  afterAssignee,
			},
		)
	}

	beforeLabels := strings.Join(before.Labels, ", ")
	afterLabels := strings.Join(after.Labels, ", ")

	if beforeLabels != afterLabels {
		result.Changes = append(
			result.Changes,
			FieldChange{
				Field:  "Labels",
				Before: beforeLabels,
				After:  afterLabels,
			},
		)
	}

	for _, name := range customFields {
		b := before.CustomValue(name)
		a := after.CustomValue(name)

		if b != a {
			if b == issue.NilValue || b == issue.UnknownValue {
				b = ""
			}

			result.Changes = append(
				result.Changes,
				FieldChange{
					Field:  name,
					Before: b,
					After:  a,
				},
			)
		}
	}

	return result
}
