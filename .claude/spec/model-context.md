# Model Context Protocol Integration

Pattern for exposing MCP tools from a service daemon, mounted on the same HTTP server as the REST API.

## When to Use

Add a `model_context/` subpackage when a daemon needs to expose MCP tools alongside its REST API. Tools typically mirror the read-oriented operations already available via REST endpoints.

## Package Structure

```
pkg/tool/go<tool>d/
└── model_context/
    ├── server.go              # Server struct (holds MCPServer + dependencies)
    ├── new.go                 # New(s *store.Store, ...) *Server
    ├── nested.go              # Nested() *server.MCPServer - for mounting
    ├── register.go            # register() - wires all tools to the MCPServer
    ├── <tool_name>.go         # One file per tool function (get_alerts.go, etc.)
    └── model_context_test.go  # Stub test
```

`model_context/` is a sibling of `route/`. Both implement the same domain operations - `route/` over REST, `model_context/` over MCP.

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

`register.go` - one `AddTool` call per tool, function in its own file:
```go
func (s *Server) register() {
    s.server.AddTool(
        mcp.NewTool(
            "get_alerts",
            mcp.WithDescription("..."),
            mcp.WithString("name", mcp.Required()),
        ),
        s.getAlerts,
    )
}
```

Tool function files (`get_alerts.go`, etc.):
```go
func (s *Server) getAlerts(
    _ context.Context,
    r mcp.CallToolRequest,
) (*mcp.CallToolResult, error) {
    name, f := r.RequireString("name")

    if f != nil {
        return mcp.NewToolResultError(fmt.Sprintf("name is required: %v", f)), nil
    }

    records := s.store.ByName(name)

    return mcp.NewToolResultText(
        fmt.Sprintf("Found %d:\n%s", len(records), notation.MarshallIndent(records)),
    ), nil
}
```

- Required params: `r.RequireString(...)` - return `mcp.NewToolResultError(...)` on failure, not a Go error
- Optional params: `r.GetString("key", "")` / `r.GetFloat("key", 0)`
- Results: `mcp.NewToolResultText(...)` for success
- Error handling is two-tier — input validation vs infrastructure failures. See `error-handling.md`.

## Wiring into run.go

Mount alongside REST on the same mux - REST routes (`/api/v1/...`) and MCP routes (`/mcp`, `/sse`, `/message`) don't conflict:

```go
import (
    generative "github.com/funtimecoding/go-library/pkg/generative/model_context/server"
    "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/model_context"
    generated "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/server"
)

lifecycle.WithServer(
    web.Listen,
    func(m *http.ServeMux) {
        generated.HandlerFromMux(route.New(s, p), m)
        generative.New(model_context.New(s, p).Nested()).Setup(m)
    },
)
```

- `generative` is the standard alias for `pkg/generative/model_context/server` (the HTTP transport infrastructure)
- `generated` is the standard alias for the oapi-codegen server package (`pkg/tool/go<tool>d/server`)
- The local `model_context` package is imported without alias

## What Not To Do

- Don't create a separate lifecycle server for MCP - one port, one mux
- Don't name it `mcp/` - use `model_context/` (no acronyms in package names)
- Don't put HTTP transport concerns in `model_context/` - that belongs in `generative.New(...).Setup(m)`
