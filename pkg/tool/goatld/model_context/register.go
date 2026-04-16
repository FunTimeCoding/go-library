package model_context

import "github.com/mark3labs/mcp-go/mcp"

func (s *Server) register() {
	s.server.AddTool(
		mcp.NewTool(
			"jira_search",
			mcp.WithDescription("Search Jira issues using JQL"),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("JQL query string"),
			),
			mcp.WithNumber(
				"limit",
				mcp.Description("Maximum number of results"),
			),
		),
		s.searchIssues,
	)
	s.server.AddTool(
		mcp.NewTool(
			"jira_get_issue",
			mcp.WithDescription("Get a Jira issue by key"),
			mcp.WithString(
				"key",
				mcp.Required(),
				mcp.Description("Issue key (e.g. PROJ-123)"),
			),
		),
		s.getIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			"jira_list_projects",
			mcp.WithDescription("List all visible Jira projects"),
		),
		s.listProjects,
	)
	s.server.AddTool(
		mcp.NewTool(
			"confluence_search",
			mcp.WithDescription("Search Confluence pages using CQL or plain text"),
			mcp.WithString(
				"query",
				mcp.Required(),
				mcp.Description("CQL query string or plain text"),
			),
		),
		s.searchPages,
	)
	s.server.AddTool(
		mcp.NewTool(
			"confluence_get_page",
			mcp.WithDescription("Get a Confluence page by ID with body content as markdown"),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Page ID"),
			),
		),
		s.getPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			"confluence_create_page",
			mcp.WithDescription("Create a new Confluence page with markdown content"),
			mcp.WithString(
				"space_identifier",
				mcp.Required(),
				mcp.Description("Space ID"),
			),
			mcp.WithString(
				"parent_identifier",
				mcp.Required(),
				mcp.Description("Parent page ID"),
			),
			mcp.WithString(
				"title",
				mcp.Required(),
				mcp.Description("Page title"),
			),
			mcp.WithString(
				"body",
				mcp.Required(),
				mcp.Description("Page content in markdown"),
			),
		),
		s.createPage,
	)
	s.server.AddTool(
		mcp.NewTool(
			"confluence_update_page",
			mcp.WithDescription("Update a Confluence page with markdown content. Gets the current version automatically."),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				"title",
				mcp.Required(),
				mcp.Description("Page title"),
			),
			mcp.WithString(
				"body",
				mcp.Required(),
				mcp.Description("Full page content in markdown"),
			),
			mcp.WithString("message", mcp.Description("Version comment")),
		),
		s.updatePage,
	)
	s.server.AddTool(
		mcp.NewTool(
			"confluence_list_spaces",
			mcp.WithDescription("List all visible Confluence spaces"),
		),
		s.listSpaces,
	)
	s.server.AddTool(
		mcp.NewTool(
			"confluence_get_page_children",
			mcp.WithDescription("List child pages of a Confluence page"),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Parent page ID"),
			),
		),
		s.getPageChildren,
	)
	s.server.AddTool(
		mcp.NewTool(
			"confluence_add_comment",
			mcp.WithDescription("Add a comment to a Confluence page"),
			mcp.WithString(
				"identifier",
				mcp.Required(),
				mcp.Description("Page ID"),
			),
			mcp.WithString(
				"body",
				mcp.Required(),
				mcp.Description("Comment text"),
			),
		),
		s.addPageComment,
	)
	s.server.AddTool(
		mcp.NewTool(
			"jira_get_transitions",
			mcp.WithDescription("List available transitions for a Jira issue"),
			mcp.WithString(
				"key",
				mcp.Required(),
				mcp.Description("Issue key"),
			),
		),
		s.getTransitions,
	)
	s.server.AddTool(
		mcp.NewTool(
			"jira_transition_issue",
			mcp.WithDescription("Transition a Jira issue to a new status"),
			mcp.WithString(
				"key",
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				"transition_identifier",
				mcp.Required(),
				mcp.Description("Transition ID from jira_get_transitions"),
			),
		),
		s.transitionIssue,
	)
	s.server.AddTool(
		mcp.NewTool(
			"jira_add_comment",
			mcp.WithDescription("Add a comment to a Jira issue"),
			mcp.WithString(
				"key",
				mcp.Required(),
				mcp.Description("Issue key"),
			),
			mcp.WithString(
				"body",
				mcp.Required(),
				mcp.Description("Comment text"),
			),
		),
		s.addIssueComment,
	)
}
