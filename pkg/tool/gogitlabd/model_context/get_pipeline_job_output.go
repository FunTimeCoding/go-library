package model_context

import (
	"context"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/go-library/pkg/tool/gogitlabd/model_context/argument"
	"github.com/mark3labs/mcp-go/mcp"
	"io"
)

func (s *Server) GetPipelineJobOutput(
	_ context.Context,
	_ mcp.CallToolRequest,
	a argument.GetPipelineJobOutput,
) (*mcp.CallToolResult, error) {
	if a.Project == "" {
		return response.Fail("project is required")
	}

	if a.Job == 0 {
		return response.Fail("job is required")
	}

	r, _, e := s.client.Jobs.GetTraceFile(a.Project, a.Job)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	b, e := io.ReadAll(r)

	if e != nil {
		return s.captureFail(e, constant.Unreachable)
	}

	return mcp.NewToolResultText(string(b)), nil
}
