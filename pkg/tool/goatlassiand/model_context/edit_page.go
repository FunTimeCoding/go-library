package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/page"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/goatlassiand/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) editPage(
	_ context.Context,
	r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
	identifier, f := r.RequireString(parameter.Identifier)

	if f != nil {
		return response.Fail("identifier is required: %v", f)
	}

	oldText, g := r.RequireString(constant.OldText)

	if g != nil {
		return response.Fail("old_text is required: %v", g)
	}

	newText, h := r.RequireString(constant.NewText)

	if h != nil {
		return response.Fail("new_text is required: %v", h)
	}

	current, i := s.confluence.Page(identifier)

	if i != nil {
		return s.captureFail(i, "page not found")
	}

	markdown := page.ToMarkdown(current.Raw.Body.Storage.Value)
	newMarkdown, k := replaceUnique(markdown, oldText, newText)

	if k != nil {
		return response.Fail("%s", k)
	}

	title := r.GetString(parameter.Title, current.Name)
	message := r.GetString(parameter.Message, "")
	result, j := s.confluence.UpdatePageAt(
		identifier,
		title,
		newMarkdown,
		current.Raw.Version.Number+1,
		message,
	)

	if j != nil {
		return s.captureFail(j, "page not updated")
	}

	return response.SuccessAny(result)
}
