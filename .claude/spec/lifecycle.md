# Lifecycle Spec

Reusable application lifecycle manager for service-style applications that serve HTTP routes and run background workers. Handles startup ordering and reverse-order shutdown.

## Package

`pkg/lifecycle/`

## Core Concept

Components (servers and workers) are registered via functional options. They start in registration order and stop in reverse registration order. `RunUntilSignal()` handles the run/block/stop sequence.

## File Layout

```
pkg/lifecycle/
  constant.go                         - Option type
  lifecycle.go                        - Lifecycle struct
  new.go                              - constructor
  run.go                              - Run() starts all components in order
  run_until_signal.go                 - RunUntilSignal() runs, blocks, stops
  stop.go                             - Stop() stops all in reverse order
  with_server.go                      - WithServer option (no timeout — streaming/MCP)
  with_server_middleware.go           - WithServerMiddleware option (no timeout)
  with_protected_server.go            - WithProtectedServer option (10s — plain REST)
  with_protected_server_middleware.go - WithProtectedServerMiddleware option (10s + middleware)
  with_verbose.go                     - WithVerbose option
  with_worker.go                      - WithWorker option
  component/
    component.go            - internal component adapter
    start.go
    stop.go
  server/
    constant.go             - readWriteTimeout for protected servers
    server.go               - server struct (internal)
    new.go                  - New constructor (no timeout)
    new_protected.go        - NewProtected constructor (sets protected flag)
    start.go                - sets up mux, creates http.Server, applies timeout if protected
    stop.go                 - graceful HTTP shutdown
```

## Worker Interface

The `face.Worker` interface from `pkg/face`:

```go
type Worker interface {
    Start()
    Stop()
}
```

Implemented by: `metric.Server`, `ticker.Ticker`, `reporter.Reporter`, and any struct with `Start()`/`Stop()`.

## WithServer / WithServerMiddleware

For streaming servers (MCP, SSE) — no read/write timeout. The lifecycle owns the `*http.ServeMux` and `*http.Server`; the setup callback only registers routes.

```go
lifecycle.WithServer(address, func(m *http.ServeMux) {
    v.Setup(m) // MCP or SSE handler
})

lifecycle.WithServerMiddleware(address, func(m *http.ServeMux) {
    // register routes on m
}, middlewareFunc)
```

## WithProtectedServer / WithProtectedServerMiddleware

For plain request/response REST servers — applies a 10s read/write timeout (slowloris protection). Use these whenever the server does not serve streaming or long-lived connections.

```go
lifecycle.WithProtectedServer(address, func(m *http.ServeMux) {
    m.HandleFunc("/health", health)
})

lifecycle.WithProtectedServerMiddleware(address, func(m *http.ServeMux) {
    // register routes on m
}, tokenMiddleware)
```

Mixed servers (REST routes + MCP on the same mux) must use `WithServer` — the streaming endpoint governs the timeout choice.

## Registration Order Matters

Components start in the order they are registered. This is intentional:

```go
// Worker initializes before servers accept traffic
lifecycle.New(
    lifecycle.WithWorker(w),   // starts first
    lifecycle.WithServer(...), // starts second
)
```

Stop happens in reverse: the server shuts down before the worker stops.

## What Stays Outside Lifecycle

- **Sentry** - lives in `Main()`, not `Run()`. See `entrypoint.md`. The `recover()` defer in `Main()` catches panics from `Run()`.
- **App-specific setup** - database connections, client construction, configuration parsing. All happen before `lifecycle.New()`.

## Usage Pattern

```go
func Run(o *option.Config) {
    // 1. Construct components
    e := metric.New(0, o.Verbose, nil)

    // 2. Build lifecycle in desired start order
    l := lifecycle.New(
        lifecycle.WithVerbose(o.Verbose),
        lifecycle.WithWorker(e),
        lifecycle.WithWorker(ticker.New(5*time.Minute, func() { ... })),
        lifecycle.WithServer(":8080", func(m *http.ServeMux) {
            m.HandleFunc("/health", health)
        }),
    )

    // 3. Run until signal, stop in reverse order
    l.RunUntilSignal()
}
```

## Adapting Existing Types

To make a type usable as a `lifecycle.Worker`:

- **Move runtime parameters into constructor.** If `Start(hourly bool)` takes args, move them to `New(..., hourly)` so `Start()` takes none.
- **Move callbacks into options.** If `Start(fn)` takes a callback, add a `WithSubscriber(fn)` option so `Start()` reads from the struct.
- **HTTP servers become route registrars.** If a server type owns its own `*http.Server`, extract a `Setup(m *http.ServeMux)` method and let lifecycle own the HTTP serving via `WithServer`.

## Go-Library Components with Start()/Stop()

These types implement the Worker interface:

| Package                      | Type       | What it does                     |
|------------------------------|------------|----------------------------------|
| `pkg/metric`                 | `Server`   | Prometheus metrics HTTP endpoint |
| `pkg/ticker`                 | `Ticker`   | Periodic function execution      |
| `pkg/errors/sentry/reporter` | `Reporter` | Sentry hub lifecycle + flush     |
