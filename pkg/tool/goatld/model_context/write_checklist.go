package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/system"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/convert"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/trivago/tgo/tcontainer"
)

func (s *Server) writeChecklist(
	c context.Context,
	key string,
	items []convert.ChecklistItem,
) (*mcp.CallToolResult, error) {
	field := s.jira.FieldMap().ByName(constant.ChecklistField)

	if field == nil {
		return response.Fail("checklist field not found")
	}

	raw := issue.Raw(key)
	raw.Fields.Unknowns = make(tcontainer.MarshalMap)
	raw.Fields.Unknowns.Set(field.Key, formatChecklist(items))
	_, resp, e := s.jira.Nested().Issue.UpdateWithContext(c, raw)

	if e != nil {
		if resp != nil && resp.Body != nil {
			return response.Fail(
				"checklist update failed: %s",
				string(system.ReadAll(resp.Body)),
			)
		}

		return response.Fail("checklist update failed: %v", e)
	}

	return nil, nil
}
