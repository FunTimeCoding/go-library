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
│   ├── start.go                    # Start() — goroutine with ticker
│   ├── stop.go                     # Stop()
│   ├── poll.go                     # Poll() — single cycle logic
│   └── poller_test.go              # assert.NotNil(t, New(...))
└── route/                          # HTTP route handlers
    ├── <route>.go                  # Handler function returning http.HandlerFunc
    ├── response.go                 # Response struct (if needed)
    ├── constant.go                 # Route-specific constants
    └── route_test.go               # Stub test
```

## Entry Point

`cmd/go<tool>/main.go` declares linker variables and delegates:

```go
package main

import "github.com/funtimecoding/go-library/pkg/tool/go<tool>"

var (
    Version   string
    GitHash   string
    BuildDate string
)

func main() {
    go<tool>.Main(Version, GitHash, BuildDate)
}
```

## Main Function

`Main()` registers flags, parses with `monitor.ParseBind`, builds the option struct, and calls `Run()`:

```go
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    pflag.String(argument.Path, "tmp/go<tool>.db", "Database path")
    monitor.ParseBind(version, gitHash, buildDate)
    o := option.New()
    o.DatabasePath = viper.GetString(argument.Path)
    Run(o)
}
```

`monitor.ParseBind` adds `--version` flag, calls `argument.ParseBind()`, and exits on `--version`.

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
    l.Run()
    system.KillSignalBlock()
    l.Stop()
}
```

Key conventions:
- Server address uses `web/constant.Listen` (`:8080`)
- Routes registered in the `func(*http.ServeMux)` callback
- `system.KillSignalBlock()` between `Run()` and `Stop()`
- Store closed via `defer` before lifecycle starts

## Sentry Integration

Optional error reporting via `SENTRY_LOCATOR` environment variable. When set, unhandled panics in `Run()` are reported to Sentry before the process exits:

```go
func Run(o *option.Log) {
    g := logger.New(context.Background())
    locator := environment.Optional(constant.LocatorEnvironment)

    if locator != "" {
        r := reporter.New("go<tool>", locator, "", o.Version)
        r.Start()
        defer func() { r.RecoverFlush(recover()) }()
    }

    s := store.New(o.DatabasePath)
    defer s.Close()
    // ...
}
```

- Sentry setup comes before store/lifecycle so the defer runs last (captures panics from all later code)
- `environment.Optional` returns empty string when the variable is unset — no error, no Sentry
- `reporter.RecoverFlush(recover())` captures the panic value, reports it, and flushes before exit
- The option struct carries `Version` so it can be passed to the reporter
- Workers should use their own `defer/recover` to log errors via the structured logger rather than crashing the process

## Route Handlers

Handlers live in the `route/` package. Each handler is a function returning `http.HandlerFunc`:

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

## Workers

Workers implement `face.Worker` (`Start()` + `Stop()`). See `lifecycle.md` for details.
