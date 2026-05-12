# Service Tool Spec

Pattern for long-running service tools that use lifecycle for HTTP servers and background workers.

## Package Structure

```
cmd/go<tool>/
└── main.go                         # Linker vars, delegates to Main()

pkg/tool/go<tool>/
├── main.go                         # Main(): register flags, parse, build option, call Run()
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
│   └── <operation>.go              # One file per operation (save.go, resolve.go, by_name.go)
├── worker/                         # Background worker (if needed)
│   ├── worker.go                   # Worker struct
│   ├── new.go                      # Constructor
│   ├── start.go                    # Start() - goroutine with ticker
│   ├── stop.go                     # Stop()
│   └── poll.go                     # Poll() - single cycle logic
└── server/                         # REST implementation (or manual routes)
    ├── <operation>.go              # Route function returning http.HandlerFunc
    ├── response.go                 # Response struct (if needed)
    └── constant.go                 # Server-specific constants
```

## Entry Point

See `entrypoint.md` for linker variables, `Main()`, and sentry setup.

`Main()` registers the `--port` flag, builds the option struct, and calls `Run()`. Sentry and argument parsing are handled per the entrypoint convention.

```go
a := argument.NewInstance(constant.Identity)
a.Integer(argument.Port, web.ListenPort, web.PortUsage)
a.Parse(version, gitHash, buildDate)
o := option.New()
o.Port = a.RequiredInteger(argument.Port)
Run(o)
```

- `web.ListenPort` (8080) is the default; override with `--port`
- `a.RequiredInteger` reads the flag value from the scoped instance
- Port lives in the option struct as `int`, not as a formatted address string

## Run Function

`Run(o, r)` constructs components and wires lifecycle. It receives the
reporter from `Main()` as `face.Reporter` - not on the option struct.
See `three-pillars.md` for the full wiring rationale.

```go
func Run(o *option.Log, r face.Reporter) {
    l := logger.New(context.Background())
    s := store.New(o.DatabasePath)
    defer s.Close()
    lifecycle.New(
        l,
        lifecycle.WithWorker(worker.New(client, s, l, r, 1*time.Minute)),
        lifecycle.WithServerMiddleware(
            web.AddressPort(o.Port),
            func(m *http.ServeMux) {
                m.HandleFunc("/api/alerts", server.Alerts(s))
            },
            web.RecoveryMiddleware(r),
        ),
    ).RunUntilSignal()
}
```

Key conventions:
- Logger constructed first, threaded to workers and lifecycle
- Reporter threaded to workers (for `withRecovery`) and recovery middleware
- Use `WithServerMiddleware` (streaming/MCP) or `WithProtectedServerMiddleware` (plain REST) with `web.RecoveryMiddleware(r)`
- Server address uses `web.AddressPort(o.Port)` to format the port as `":8080"`
- Routes registered in the `func(*http.ServeMux)` callback
- `RunUntilSignal()` handles run, signal block, and reverse-order stop
- Store closed via `defer` before lifecycle starts
- Avoid declaring intermediate variables for lifecycle or generative server when only used once

## Server Functions

When routes are defined by an OpenAPI spec, use the Server struct
pattern - see `generated-api.md`.

For manually registered routes without a spec, route functions live in
the `server/` package. Each is a function returning `http.HandlerFunc`:

```go
func Alerts(s *store.Store) http.HandlerFunc {
    return func(
        w http.ResponseWriter,
        r *http.Request,
    ) {
        // ...
        web.EncodeNotation(w, result)
    }
}
```

- `web.EncodeNotation(w, result)` - sets Content-Type JSON header and encodes result as JSON (the common case)
- `web.Encode(w, result)` - encodes result as JSON without setting headers (use when headers are set separately, e.g. with a non-200 status code)
- `web.ObjectHeader(w)` - sets Content-Type JSON header only (use with `web.Encode` when you need a custom status code between header and body)
- Use `argument.*` constants for query parameter names
- Response structs in `response.go`, constants in `constant.go`

## HTML Web Package

When a service tool serves an HTML UI (using gomponents), routes are grouped under a `web/` package rather than `server/`. Use this pattern instead of `server/` when handlers need shared state and render HTML rather than JSON.

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
├── layout.go           # layout() - full page shell
├── navigation_link.go  # navigationLink() - navigation component
└── constant/
    └── constant.go     # inlineCSS and other web-layer constants
```

Key conventions:
- `Server` struct holds dependencies (store, worker, etc.)
- `Mount()` registers all routes using method values: `m.HandleFunc("GET /alerts", s.alerts)`
- Handler methods named after the route, no `handle` prefix: `alerts()`, `dashboard()`, `addSubmit()`
- Standalone HTML builders named after the component they produce: `alertsTable()`, `addForm()`, `detailRow()`
- Handler and builder names never collide because the handler is named after the route (`add`), not the component (`addForm`)
- Methods that access server state (store, worker) stay as methods; pure renderers are standalone functions

## Workers

Workers implement `face.Worker` (`Start()` + `Stop()`). See `lifecycle.md` for details.

## Daemon / CLI Split

Long-running services come in pairs:

| Binary | Package | Role |
|--------|---------|------|
| `go<tool>d` | `pkg/tool/go<tool>d/` | Daemon - lifecycle, store, HTTP server |
| `go<tool>` | `pkg/tool/go<tool>/` | CLI - calls the daemon's REST API, prints output |

The CLI tool is a thin wrapper around the domain client library (see below). After the standard entrypoint setup (see `entrypoint.md`), it calls a single method on the domain client.

## External API Client

When a daemon proxies a third-party service (Habitica, Mattermost,
Jellyfin, Sentry, etc.), the client library for that service
lives at `pkg/<domain>/`. This is the external API client - it talks
directly to the upstream service, not to our daemon.

```
pkg/<domain>/
├── client.go                       # Client struct (raw HTTP or SDK wrapper)
├── new.go                          # New(host, token string) *Client
├── new_environment.go              # NewEnvironment() - reads env vars
├── <operation>.go                  # One file per operation
└── constant/
    └── constant.go                 # HostEnvironment, TokenEnvironment
```

The daemon imports this client library. MCP-only services (without
REST) import it directly into their `model_context/` handlers. Daemons
with REST import it into both `model_context/` and `server/`.

**`pkg/<domain>/` is client-only.** All server-side concerns (store,
domain types, `model_context/`) belong in `pkg/tool/go<tool>d/`, not
in the shared library.

## Internal REST Client

When a daemon exposes a REST API (via OpenAPI codegen), the CLI needs
a clean interface to call it. This wrapper lives inside the daemon's
tree at `<toold>/client/`, not at `pkg/<domain>/`. See
`generated-api.md` for the full pattern.

The distinction:
- `pkg/<domain>/` - external API client, talks to a third-party service
- `<toold>/client/` - internal REST client, talks to our own daemon

## MCP Integration

When a daemon also exposes MCP tools, add a `model_context/` subpackage
and mount it on the same mux as the REST API. See `model-context.md`.

`model_context/` is the standard package name - not `tool/` or
`toolset/`.
