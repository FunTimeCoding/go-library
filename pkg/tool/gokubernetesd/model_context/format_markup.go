package model_context

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/mark/response"
	"github.com/funtimecoding/go-library/pkg/strings/join"
	"github.com/mark3labs/mcp-go/mcp"
	"sigs.k8s.io/yaml"
)

func formatMarkup(
	object map[string]interface{},
	filtered []string,
) (*mcp.CallToolResult, error) {
	b, e := yaml.Marshal(object)

	if e != nil {
		return response.Fail("failed to marshal yaml: %s", e)
	}

	if len(filtered) == 0 {
		return response.Success(string(b))
	}

	return response.Success(
		fmt.Sprintf(
			"# filtered: %s\n%s",
			join.CommaSpace(filtered),
			b,
		),
	)
}
