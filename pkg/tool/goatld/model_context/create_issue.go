package model_context

import (
	"context"
	"encoding/json"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/trivago/tgo/tcontainer"
)

func (s *Server) createIssue(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	project, f := r.RequireString(constant.Project)

	if f != nil {
		return response.Fail("project is required: %v", f)
	}

	issueType, g := r.RequireString(constant.IssueType)

	if g != nil {
		return response.Fail("issue_type is required: %v", g)
	}

	summary, h := r.RequireString(constant.Summary)

	if h != nil {
		return response.Fail("summary is required: %v", h)
	}

	description := r.GetString("description", "")
	assigneeName := r.GetString(constant.Assignee, "")
	labelsRaw := r.GetString(constant.Labels, "")
	fieldsRaw := r.GetString(constant.AdditionalFields, "")
	raw := issue.RawStub()
	raw.Fields.Unknowns = make(tcontainer.MarshalMap)
	raw.Fields.Reporter = s.jira.User()
	raw.Fields.Project.Key = project
	raw.Fields.Type.Name = issueType
	raw.Fields.Summary = summary
	raw.Fields.Description = description

	if labelsRaw != "" {
		var labels []string

		if i := json.Unmarshal(
			[]byte(labelsRaw),
			&labels,
		); i != nil {
			return response.Fail(
				"labels must be a JSON array of strings: %v",
				i,
			)
		}

		raw.Fields.Labels = labels
	}

	if fieldsRaw != "" {
		var fields map[string]any

		if i := json.Unmarshal(
			[]byte(fieldsRaw),
			&fields,
		); i != nil {
			return response.Fail(
				"additional_fields must be a JSON object: %v",
				i,
			)
		}

		fieldMap := s.jira.FieldMap()

		for name, value := range fields {
			field := fieldMap.ByName(name)

			if field == nil {
				return response.Fail(
					"unknown field: %s",
					name,
				)
			}

			raw.Fields.Unknowns.Set(field.Key, value)
		}
	}

	result, resp, i := s.jira.Nested().Issue.CreateWithContext(
		c,
		raw,
	)

	if i != nil {
		if resp != nil && resp.Body != nil {
			return response.Fail(
				"create failed: %s",
				string(system.ReadAll(resp.Body)),
			)
		}

		return response.Fail("create failed: %v", i)
	}

	if assigneeName != "" {
		user, fail, j := s.resolveAssignee(c, assigneeName)

		if fail != nil {
			return fail, j
		}

		assignResp, k := s.jira.Nested().Issue.UpdateAssigneeWithContext(
			c,
			result.Key,
			user,
		)

		if k != nil {
			if assignResp != nil && assignResp.Body != nil {
				return response.Fail(
					"created %s but assign failed: %s",
					result.Key,
					string(system.ReadAll(assignResp.Body)),
				)
			}

			return response.Fail(
				"created %s but assign failed: %v",
				result.Key,
				k,
			)
		}
	}

	return response.SuccessAny(
		convert.JiraIssue(s.jira.Issue(result.Key)),
	)
}
