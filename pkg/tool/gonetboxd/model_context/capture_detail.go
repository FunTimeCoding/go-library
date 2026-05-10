package model_context

import (
	"encoding/json"
	"errors"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/basic/response"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/netbox-community/go-netbox/v4"
)

func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
	if f, okay := errors.AsType[netbox.GenericOpenAPIError](e); okay {
		var r response.Error

		if json.Unmarshal(f.Body(), &r) == nil && r.Detail != "" {
			return s.captureFail(e, r.Detail)
		}
	}

	return s.captureFail(e, constant.UnexpectedError)
}
