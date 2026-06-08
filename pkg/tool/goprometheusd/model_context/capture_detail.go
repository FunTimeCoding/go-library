package model_context

import (
	"errors"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/validation"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if v, okay := errors.AsType[*validation.Detail](e); okay {
		return response.Fail("%s", v.Message)
	}

	return s.captureFail(e, constant.UnexpectedError)
}
