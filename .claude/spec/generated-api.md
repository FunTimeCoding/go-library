# Generated API

Pattern for typed HTTP APIs using `oapi-codegen`. Use when a service
exposes an HTTP API that needs a generated client (or when type safety
on the server side is worth the codegen overhead).

## Package Structure

```
pkg/tool/go<tool>d/
├── generated/
│   ├── client/
│   │   ├── config.yaml    # oapi-codegen client config
│   │   └── generated.go   # Generated: Client, typed request/response methods
│   └── server/
│       ├── openapi.yaml   # Hand-written spec
│       ├── config.yaml    # oapi-codegen server config
│       └── generated.go   # Generated: ServerInterface, HandlerFromMux, types
├── client/                # Domain wrapper for CLI (imports generated/client/)
│   ├── client.go          # Client struct (wraps generated client)
│   ├── new.go             # New(host string) *Client
│   └── <operation>.go     # One file per operation
├── server/                # REST implementation (imports generated/server/)
│   ├── server.go          # Server struct (holds dependencies)
│   ├── new.go             # New(dep) *Server
│   └── <operation>.go     # One file per endpoint (post_deploy.go, get_status.go)
├── model_context/         # MCP implementation
├── convert/               # Type filtering shared by model_context/ and server/
├── constant/
├── option/
├── main.go
└── run.go
```

Everything under `generated/` is machine output - don't edit it. `client/`
and `server/` are hand-written code that consumes the generated types.

## Configs

`generated/server/config.yaml`:
```yaml
package: server
generate:
  std-http-server: true
  models: true
  embedded-spec: true
output: generated.go
```

`generated/client/config.yaml`:
```yaml
package: client
generate:
  client: true
  models: true
output: generated.go
```

## Generation

Run via taskfile:
```yaml
generate:
  cmds:
  - cd pkg/<service>/generated/server && oapi-codegen --config config.yaml openapi.yaml
  - cd pkg/<service>/generated/client && oapi-codegen --config config.yaml ../server/openapi.yaml
```

Add generated files to golint skip list:
```yaml
lint:
  cmds: ['golint --fix --skip tmp,pkg/<service>/generated/server/generated.go,pkg/<service>/generated/client/generated.go']
```

## Implementing the Server

`server/server.go` - plain struct, no embedding:
```go
type Server struct {
    store *store.Store
}
```

`server/<operation>.go` - implements the generated `ServerInterface`
method:
```go
func (s *Server) PostDeploy(w http.ResponseWriter, q *http.Request) {
    var body generated.PostDeployJSONRequestBody
    errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
    web.EncodeNotation(w, generated.DeployResponse{
        Tag: s.store.TriggerTargets(body.Targets),
    })
}
```

Use `web.EncodeNotation(w, result)` for the common case (200 + JSON).
For non-200 status codes, use `web.ObjectHeader(w)` +
`w.WriteHeader(...)` + `web.Encode(w, result)` separately.

When a server operation uses converters shared with `model_context/`,
import the `convert/` package:
```go
func (s *Server) GetIssue(w http.ResponseWriter, _ *http.Request, key string) {
    web.EncodeNotation(w, convert.JiraIssue(s.store.Issue(key)))
}
```

## Client Wrapper

The `client/` package wraps the generated REST client, giving the CLI
a clean interface independent of generated types. This is the internal
REST client - it talks to our own daemon, not to a third-party service.

`client/client.go`:
```go
type Client struct {
    context context.Context
    client  *generated.ClientWithResponses
}
```

`client/new.go`:
```go
func New(host string) *Client {
    c, e := generated.NewClientWithResponses(locator.New(host).String())
    errors.PanicOnError(e)
    return &Client{context: context.Background(), client: c}
}
```

`client/<operation>.go`:
```go
func (c *Client) Alerts() string {
    result, e := c.client.GetAlerts(c.context, &generated.GetAlertsParams{})
    errors.PanicOnError(e)
    return web.ReadString(result)
}
```

For services with a third-party upstream, the external API client lives
at `pkg/<domain>/` (see `service-tool.md`). The `client/` wrapper and
the external API client serve different consumers - `client/` is for the
CLI talking to the daemon, `pkg/<domain>/` is for the daemon (and
MCP-only services) talking to the upstream.

## Wiring into run.go

The generated server package is aliased `generated`. Call
`generated.HandlerFromMux` inside `lifecycle.WithServer`:

```go
import (
    generated "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/generated/server"
)

lifecycle.WithServer(
    web.Listen,
    func(m *http.ServeMux) {
        generated.HandlerFromMux(server.New(dep), m)
    },
)
```

When MCP is also mounted, add the `model_context` import and call
`Mount(m)`:

```go
import (
    "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/model_context"
    generated "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/generated/server"
)

lifecycle.WithServerMiddleware(
    web.AddressPort(o.Port),
    func(m *http.ServeMux) {
        generated.HandlerFromMux(server.New(dep), m)
        model_context.New(dep, r).Mount(m)
    },
    web.RecoveryMiddleware(r),
)
```

REST routes (`/api/...`) and MCP routes (`/mcp`, `/sse`, `/message`)
don't conflict on the same mux.

## OpenAPI Spec Patterns

Optional arrays of objects generate `*[]*Type` when the items are
nullable. Use `nullable: true` with `allOf` wrapping:

```yaml
checklist:
  # Optional: only present on certain task types.
  type: array
  items:
    nullable: true
    allOf:
    - $ref: "#/components/schemas/ChecklistItem"
```

Without `nullable: true` on items, oapi-codegen generates
`*[]Type` (pointer to slice of values). The nullable pattern
keeps pointer convention consistent between the generated types
and the `convert/` layer.

## What Not To Do

- Don't put hand-written code in `generated/` - that's machine output
- Don't edit `generated.go` - regenerate from the spec instead
- Don't import generated types in the CLI - use the `client/` wrapper
- Don't manually register routes that the spec already defines -
  let `HandlerFromMux` do it
