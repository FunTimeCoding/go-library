package mark

import (
	"context"
	"fmt"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/constant"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/example/mark/option"
	"github.com/funtimecoding/go-library/pkg/generative/model_context/server"
	"github.com/funtimecoding/go-library/pkg/project"
	"github.com/funtimecoding/go-library/pkg/system"
	web "github.com/funtimecoding/go-library/pkg/web/constant"
	"github.com/mark3labs/mcp-go/mcp"
	markServer "github.com/mark3labs/mcp-go/server"
	"os"
)

func Run(o *option.Mark) {
	s := markServer.NewMCPServer(
		"Demo",
		constant.DefaultVersion,
		markServer.WithToolCapabilities(false),
	)
	s.AddTool(
		mcp.NewTool(
			constant.GreetTool,
			mcp.WithDescription("Say hello to someone"),
			mcp.WithString(
				constant.NameParameter,
				mcp.Required(),
				mcp.Description("Name of the person to greet"),
			),
		),
		func(
			_ context.Context,
			r mcp.CallToolRequest,
		) (*mcp.CallToolResult, error) {
			name, e := r.RequireString(constant.NameParameter)

			if e != nil {
				return mcp.NewToolResultError(e.Error()), nil
			}

			return mcp.NewToolResultText(
				fmt.Sprintf("Hello %s", name),
			), nil
		},
	)
	s.AddResource(
		mcp.NewResource(
			ReadmeDocument,
			"Project README",
			mcp.WithResourceDescription(
				"The project's README file",
			),
			mcp.WithMIMEType(web.Markdown),
		),
		func(
			_ context.Context,
			_ mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			content, e := os.ReadFile(project.ReadmeFile)

			if e != nil {
				return nil, e
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      ReadmeDocument,
					MIMEType: web.Markdown,
					Text:     string(content),
				},
			}, nil
		},
	)
	s.AddResourceTemplate(
		mcp.NewResourceTemplate(
			"user://{id}/profile",
			"User Profile",
			mcp.WithTemplateDescription(
				"Returns user profile information",
			),
			mcp.WithTemplateMIMEType(web.Object),
		),
		func(
			_ context.Context,
			r mcp.ReadResourceRequest,
		) ([]mcp.ResourceContents, error) {
			p, e := profile(parseLocator(r.Params.URI))

			if e != nil {
				return nil, e
			}

			return []mcp.ResourceContents{
				mcp.TextResourceContents{
					URI:      r.Params.URI,
					MIMEType: web.Object,
					Text:     p,
				},
			}, nil
		},
	)

	if o.Local {
		server.New(s).ServeLocal()
		system.KillSignalBlock()
	} else {
		v := server.New(s)
		v.Start()
		system.KillSignalBlock()
		v.Stop()
	}
}
