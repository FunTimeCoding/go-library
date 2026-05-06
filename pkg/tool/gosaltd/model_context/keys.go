package model_context

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/tool/gosaltd/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) keys(
	_ context.Context,
	_ mcp.CallToolRequest,
) (result *mcp.CallToolResult, returnError error) {
	defer func() {
		if v := recover(); v != nil {
			result, returnError = s.captureFail(
				fmt.Errorf("%v", v),
				constant.KeysFailed,
			)
		}
	}()

	return response.SuccessAny(s.runner.SaltClient().Keys())
}
