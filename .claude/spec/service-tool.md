# Service Tool Spec

Pattern for long-running service tools that use lifecycle for HTTP servers and background workers.

## Package Structure

```
cmd/go<tool>/
└── main.go                         # Linker vars, delegates to Main()

pkg/tool/go<tool>/
├── main.go                         # Main(): register flags, parse, build option, call Run()
├── main_test.go                    # Stub test
├── run.go                          # Run(o): wiring, lifecycle, signal block
├── option/
│   ├── <name>.go                   # Option struct (named after tool/domain, not "Option")
│   ├── new.go                      # Factory: New() *<Name>
│   └── <name>_test.go              # assert.NotNil(t, New())
├── store/                          # Persistence (if needed)
│   ├── store.go                    # Store struct
│   ├── constant.go                 # Bucket name
│   ├── record.go                   # Record type
│   ├── new.go                      # Constructor: opens db, ensures bucket
│   ├── close.go                    # Close()
│   ├── <operation>.go              # One file per operation (save.go, resolve.go, by_name.go)
│   └── store_test.go               # Stub test
├── poller/                         # Background worker (if needed)
│   ├── poller.go                   # Poller struct
│   ├── new.go                      # Constructor
│   ├── start.go                    # Start() - goroutine with ticker
│   ├── stop.go                     # Stop()
│   ├── poll.go                     # Poll() - single cycle logic
│   └── poller_test.go              # assert.NotNil(t, New(...))
└── route/                          # HTTP route functions
    ├── <route>.go                  # Route function returning http.HandlerFunc
    ├── response.go                 # Response struct (if needed)
    ├── constant.go                 # Route-specific constants
    └── route_test.go               # Stub test
```

## Entry Point

See `entrypoint.md` for linker variables, `Main()`, and sentry setup.

`Main()` registers flags, builds the option struct, and calls `Run()`. Sentry and `monitor.ParseBind` are handled per the entrypoint convention.

## Run Function

`Run(o)` constructs components and wires lifecycle:

```go
func Run(o *option.Log) {
    s := store.New(o.DatabasePath)
    defer s.Close()
    l := lifecycle.New(
        lifecycle.WithWorker(
            poller.New(client, s, 1*time.Minute),
        ),
        lifecycle.WithServer(
            constant.Listen,
            func(m *http.ServeMux) {
                m.HandleFunc("/api/v1/alerts", route.Alerts(s))
            },
        ),
    )
    l.RunUntilSignal()
}
```

Key conventions:
- Server address uses `web/constant.Listen` (`:8080`)
- Routes registered in the `func(*http.ServeMux)` callback
- `RunUntilSignal()` handles run, signal block, and reverse-order stop
- Store closed via `defer` before lifecycle starts

## Route Functions

This pattern is for manually registered routes. When routes are defined by an OpenAPI spec, use the Router struct pattern instead - see `openapi.md`.

Route functions live in the `route/` package. Each is a function returning `http.HandlerFunc`:

```go
func Alerts(s *store.Store) http.HandlerFunc {
    return func(
        w http.ResponseWriter,
        r *http.Request,
    ) {
        // ...
        w.Header().Set(constant.ContentType, constant.Object)
        errors.PanicOnError(json.NewEncoder(w).Encode(result))
    }
}
```

- Use `web/constant.ContentType` and `web/constant.Object` for JSON responses
- Use `argument.*` constants for query parameter names
- Response structs in `response.go`, constants in `constant.go`

## HTML Web Package

When a service tool serves an HTML UI (using gomponents), routes are grouped under a `web/` package rather than `route/`. Use this pattern instead of `route/` when handlers need shared state and render HTML rather than JSON.

```
pkg/tool/go<tool>d/web/
├── server.go           # type Server struct (dependencies only, no functions)
├── new.go              # NewServer(deps) *Server
├── mount.go            # (s *Server) Mount(m *http.ServeMux)
├── is_htmx.go          # (s *Server) isHTMX(r *http.Request) bool
├── render_page.go      # renderPage(w, page g.Node)
├── render_fragment.go  # renderFragment(w, fragment g.Node)
├── <route>.go          # handler method named after route: alerts(), dashboard()
├── <component>.go      # HTML builder named after component: alerts_table.go
├── layout.go           # layout() — full page shell
├── nav_link.go         # navLink() — navigation component
└── constant/
    └── constant.go     # inlineCSS and other web-layer constants
```

Key conventions:
- `Server` struct holds dependencies (store, poller, etc.)
- `Mount()` registers all routes using method values: `m.HandleFunc("GET /alerts", s.alerts)`
- Handler methods named after the route, no `handle` prefix: `alerts()`, `dashboard()`, `addSubmit()`
- Standalone HTML builders named after the component they produce: `alertsTable()`, `addForm()`, `detailRow()`
- Handler and builder names never collide because the handler is named after the route (`add`), not the component (`addForm`)
- Methods that access server state (store, poller) stay as methods; pure renderers are standalone functions

## Workers

Workers implement `face.Worker` (`Start()` + `Stop()`). See `lifecycle.md` for details.

## Daemon / CLI Split

Long-running services come in pairs:

| Binary | Package | Role |
|--------|---------|------|
| `go<tool>d` | `pkg/tool/go<tool>d/` | Daemon - lifecycle, store, HTTP server |
| `go<tool>` | `pkg/tool/go<tool>/` | CLI - calls the daemon's REST API, prints output |

The CLI tool is a thin wrapper around the domain client library (see below). After the standard entrypoint setup (see `entrypoint.md`), it calls a single method on the domain client.

## Domain Client Library

The REST client is exposed via a domain library at `pkg/<domain>/`, analogous to how `pkg/alert_log/` wraps the goalertlogd client. This keeps consumer code independent of the generated client types.

**`pkg/<domain>/` is client-only.** All server-side concerns (store, domain types, `model_context/`) belong in `pkg/tool/go<tool>d/`, not in the shared library.

```
pkg/<domain>/
├── client.go                       # Client struct (wraps generated client)
├── new.go                          # New(host string) *Client
├── new_environment.go              # NewEnvironment() - reads HOST env var
├── <operation>.go                  # One file per operation (entries.go, alerts.go)
└── constant/
    └── constant.go                 # HostEnvironment = "<DOMAIN>_HOST"
```

`New` wraps the generated client with `locator.New(host).String()`:
```go
func New(host string) *Client {
    c, e := client.NewClientWithResponses(locator.New(host).String())
    errors.PanicOnError(e)
    return &Client{context: context.Background(), client: c}
}
```

Operations call the plain (non-`WithResponse`) client method and return the raw string via `web.ReadString`:
```go
func (c *Client) Entries() string {
    result, e := c.client.GetEntries(c.context, &client.GetEntriesParams{})
    errors.PanicOnError(e)
    return web.ReadString(result)
}
```

## MCP Integration

When a daemon also exposes MCP tools, add a `model_context/` subpackage and mount it on the same mux as the REST API. See `model-context.md`.
