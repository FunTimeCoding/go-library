package model_context

import (
	"errors"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/web/detail_error"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if f, okay := errors.AsType[*detail_error.Detail](e); okay {
		return s.captureFail(e, f.Detail)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
