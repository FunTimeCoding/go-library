# Model Context Protocol Integration

Pattern for exposing MCP tools from a service daemon, mounted on the same HTTP server as the REST API.

## When to Use

Add a `model_context/` subpackage when a daemon needs to expose MCP tools alongside its REST API. Tools typically mirror the read-oriented operations already available via REST endpoints.

## Package Structure

```
pkg/tool/go<tool>d/
├── constant/
│   └── constant.go            # Tool name + tool-specific parameter constants
├── convert/
│   └── <type>.go              # One file per converter function
├── model_context/
│   ├── server.go              # Server struct (holds MCPServer + dependencies)
│   ├── new.go                 # New(s *store.Store, ...) *Server
│   ├── nested.go              # Nested() *server.MCPServer - for mounting
│   ├── register.go            # register() - wires all tools to the MCPServer
│   ├── <tool_name>.go         # One file per tool function (get_alerts.go, etc.)
│   └── model_context_test.go  # Stub test
└── route/
```

`model_context/` is a sibling of `route/`. Both implement the same domain operations - `route/` over REST, `model_context/` over MCP.

## Convert Package

MCP tools must never serialize raw domain objects directly. Always convert through a `convert/` package that produces slim response types — the same types used by REST routes.

```
pkg/tool/go<tool>d/convert/
├── jira_issue.go              # JiraIssue(*issue.Issue) *server.JiraIssue
├── jira_issues.go             # JiraIssues([]*issue.Issue) []*server.JiraIssue
└── ...
```

Converter functions are exported and shared by both `route/` and `model_context/`. This prevents MCP tools from leaking raw structs (with all nested fields, changelogs, custom fields, etc.) into the LLM context.

Even MCP-only daemons without REST routes should use a `convert/` package to sanitize outputs.

## Constants

Tool name and tool-specific parameter constants live in `pkg/tool/go<tool>d/constant/constant.go`:

```go
const (
    GetAlerts    = "get_alerts"
    GetStatus    = "get_status"
)

const (
    IncludeResolved = "include_resolved"
)
```

Reusable parameter names shared across multiple MCP tools (`query`, `limit`, `key`, `body`, `identifier`, `title`, `message`) live in `pkg/generative/model_context/parameter/parameter.go` and are referenced as `parameter.Query`, `parameter.Limit`, etc.

## Server Struct

`server.go`:
```go
type Server struct {
    server *server.MCPServer
    store  *store.Store
    poller *poller.Poller  // omit if no poller
}
```

`new.go` - create the MCPServer, register tools, return:
```go
func New(s *store.Store, p *poller.Poller) *Server {
    result := &Server{
        server: server.NewMCPServer(
            "<service-name>",
            constant.DefaultVersion,
            server.WithToolCapabilities(true),
        ),
        store:  s,
        poller: p,
    }
    result.register()

    return result
}
```

`nested.go` - exposes the inner MCPServer for the HTTP transport layer:
```go
func (s *Server) Nested() *server.MCPServer {
    return s.server
}
```

## Registering Tools

`register.go` - one `AddTool` call per tool, function in its own file. Tool names and parameter names are constants, never string literals:
```go
func (s *Server) register() {
    s.server.AddTool(
        mcp.NewTool(
            constant.GetAlerts,
            mcp.WithDescription("..."),
            mcp.WithString(parameter.Name, mcp.Required()),
            mcp.WithNumber(
                parameter.Limit,
                mcp.Required(),
                mcp.Description("Maximum number of results (e.g. 25)"),
            ),
        ),
        s.getAlerts,
    )
}
```

Tool function files (`get_alerts.go`, etc.) use `response.Fail` for input errors, `response.SuccessAny` for JSON results, and `response.Success` for plain text results:
```go
func (s *Server) getAlerts(
    _ context.Context,
    r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
    name, f := r.RequireString(parameter.Name)

    if f != nil {
        return response.Fail("name is required: %v", f)
    }

    limit, g := r.RequireFloat(parameter.Limit)

    if g != nil {
        return response.Fail("limit is required: %v", g)
    }

    return response.SuccessAny(
        convert.Alerts(s.store.ByName(name, int(limit))),
    )
}
```

- Required params: `r.RequireString(...)` / `r.RequireFloat(...)` — return `response.Fail(...)` on failure
- Optional params: `r.GetString(constant.Key, "")` / `r.GetFloat(constant.Key, 0)` / `r.GetBool(constant.Key, false)`
- JSON results: `response.SuccessAny(converted)` — serializes via `notation.MarshalIndent`
- Text results: `response.Success("comment added")`
- Error results: `response.Fail("message: %v", err)` — wraps `mcp.NewToolResultError`
- Error variables progress `e`, `f`, `g`, `h`, `i` — never reuse the same letter. See `naming.md`.
- Always convert results through the `convert/` package — never serialize raw domain objects
- Error handling is two-tier — input validation vs infrastructure failures. See `error-handling.md`.

## Wiring into run.go

Mount alongside REST on the same mux - REST routes (`/api/...`) and MCP routes (`/mcp`, `/sse`, `/message`) don't conflict:

```go
import (
    generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
    "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/model_context"
    generated "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/server"
    "github.com/funtimecoding/go-library/pkg/web"
)

lifecycle.WithServer(
    web.AddressPort(o.Port),
    func(m *http.ServeMux) {
        generated.HandlerFromMux(route.New(s, p), m)
        generative.New(model_context.New(s, p).Nested()).Setup(m)
    },
)
```

- `generative` is the standard alias for `pkg/generative/model_context/server` (the HTTP transport infrastructure)
- `generated` is the standard alias for the oapi-codegen server package (`pkg/tool/go<tool>d/server`)
- The local `model_context` package is imported without alias
- `generative.New(...)` is called directly inside the closure, not assigned to a variable outside it

## What Not To Do

- Don't create a separate lifecycle server for MCP - one port, one mux
- Don't name it `mcp/` - use `model_context/` (no acronyms in package names)
- Don't put HTTP transport concerns in `model_context/` - that belongs in `generative.New(...).Setup(m)`
