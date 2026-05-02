package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabmcp/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (s *Server) GetFileContents(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetFileContents,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Path == "" {
		return response.Fail("path is required")
	}

	o := &gitlab.GetFileOptions{}

	if a.Reference != "" {
		o.Ref = &a.Reference
	}

	v, _, e := s.client.RepositoryFiles.GetFile(a.Project, a.Path, o)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return response.SuccessAny(v)
}
