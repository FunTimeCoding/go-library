package model_context

import (
	"github.com/funtimecoding/go-library/pkg/generative/model_context/parameter"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/constant"
	"github.com/mark3labs/mcp-go/mcp"
)

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraSearch,
			mcp.WithDescription("Search Jira issues using JQL"),
			mcp.WithString(
				parameter.Query,
				mcp.Required(),
				mcp.Description("JQL query string"),
			),
			mcp.WithNumber(
				parameter.Limit,
				mcp.Required(),
				mcp.Description("Maximum number of results (e.g. 25)"),
			),
		),
		s.searchIssues,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetIssue,
			mcp.WithDescription("Get a Jira issue by key"),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key (e.g. PROJ-123)"),
			),
		),
		s.getIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraSearchProjects,
			mcp.WithDescription("Search Jira projects by key or name"),
			mcp.WithString(
				parameter.Query,
				mcp.Description("Filter by key (exact, case-insensitive) or name (case-insensitive contains). Omit to list all."),
			),
			mcp.WithNumber(
				parameter.Limit,
				mcp.Required(),
				mcp.Description("Maximum number of results (e.g. 25)"),
			),
			mcp.WithBoolean(
				constant.IncludeDescriptions,
				mcp.Description("Fetch full project details including description (slower, one API call per result). Defaults to false."),
			),
		),
		s.searchProjects,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceSearch,
			mcp.WithDescription("Search Confluence pages using CQL or plain text"),
			mcp.WithString(
				parameter.Query,
				mcp.Required(),
				mcp.Description("CQL query string or plain text"),
			),
		),
		s.searchPages,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceGetPage,
			mcp.WithDescription("Get a Confluence page by ID with body content as markdown"),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
		),
		s.getPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceCreatePage,
			mcp.WithDescription("Create a new Confluence page with markdown content"),
			mcp.WithString(
				constant.SpaceIdentifier,
				mcp.Required(),
				mcp.Description("Space ID"),
			),
			mcp.WithString(
				constant.ParentIdentifier,
				mcp.Required(),
				mcp.Description("Parent page ID"),
			),
			mcp.WithString(
				parameter.Title,
				mcp.Required(),
				mcp.Description("Page title"),
			),
			mcp.WithString(
				parameter.Body,
				mcp.Required(),
				mcp.Description("Page content in markdown"),
			),
		),
		s.createPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceUpdatePage,
			mcp.WithDescription("Update a Confluence page with markdown content. Gets the current version automatically."),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				parameter.Title,
				mcp.Required(),
				mcp.Description("Page title"),
			),
			mcp.WithString(
				parameter.Body,
				mcp.Required(),
				mcp.Description("Full page content in markdown"),
			),
			mcp.WithString(
				parameter.Message,
				mcp.Description("Version comment"),
			),
		),
		s.updatePage,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceListSpaces,
			mcp.WithDescription("List all visible Confluence spaces"),
		),
		s.listSpaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceGetPageChildren,
			mcp.WithDescription("List child pages of a Confluence page"),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Parent page ID"),
			),
		),
		s.getPageChildren,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.ConfluenceAddComment,
			mcp.WithDescription("Add a comment to a Confluence page"),
			mcp.WithString(
				parameter.Identifier,
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				parameter.Body,
				mcp.Required(),
				mcp.Description("Comment text"),
			),
		),
		s.addPageComment,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraGetTransitions,
			mcp.WithDescription("List available transitions for a Jira issue"),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
		),
		s.getTransitions,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraTransitionIssue,
			mcp.WithDescription("Transition a Jira issue to a new status"),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				constant.TransitionIdentifier,
				mcp.Required(),
				mcp.Description("Transition ID from jira_get_transitions"),
			),
		),
		s.transitionIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			constant.JiraAddComment,
			mcp.WithDescription("Add a comment to a Jira issue"),
			mcp.WithString(
				parameter.Key,
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				parameter.Body,
				mcp.Required(),
				mcp.Description("Comment text"),
			),
		),
		s.addIssueComment,
	)
}
