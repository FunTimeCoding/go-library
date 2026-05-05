package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/constant"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/convert"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) addChecklistItem(
	c context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	key, f := r.RequireString(parameter.Key)

	if f != nil {
		return response.Fail("key is required: %v", f)
	}

	text, g := r.RequireString(constant.Text)

	if g != nil {
		return response.Fail("text is required: %v", g)
	}

	items, h := s.readChecklist(key)

	if h != nil {
		return s.captureFail(h, "checklist not readable")
	}

	items = append(
		items,
		convert.ChecklistItem{
			Index: len(items) + 1,
			Text:  text,
		},
	)
	fail, e := s.writeChecklist(c, key, items)

	if fail != nil {
		return fail, e
	}

	return response.SuccessAny(items)
}
