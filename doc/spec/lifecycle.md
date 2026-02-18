# Lifecycle Spec

Reusable application lifecycle manager for service-style applications that serve HTTP routes and run background workers. Handles startup ordering and reverse-order shutdown.

## Package

`pkg/lifecycle/`

## Core Concept

Components (servers and workers) are registered via functional options. They start in registration order and stop in reverse registration order. The caller decides when to block (typically `system.KillSignalBlock()`).

## File Layout

```
pkg/lifecycle/
  constant.go               — Option type
  lifecycle.go              — Lifecycle struct
  new.go                    — constructor
  run.go                    — Run() starts all components in order
  stop.go                   — Stop() stops all in reverse order
  with_server.go            — WithServer option
  with_server_middleware.go — WithServerMiddleware option
  with_verbose.go           — WithVerbose option
  with_worker.go            — WithWorker option
  component/
    component.go            — internal component adapter
    start.go
    stop.go
  server/
    server.go               — server struct (internal)
    start.go                — sets up mux, creates http.Server, serves
    stop.go                 — graceful HTTP shutdown
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

## WithServer

```go
lifecycle.WithServer(address, func(m *http.ServeMux) {
    // register routes on m
})
```

The lifecycle owns the `*http.ServeMux` and `*http.Server`. The setup callback only registers routes. This means route registrars don't need to know about HTTP server lifecycle — they just receive a mux.

## WithServerMiddleware

```go
lifecycle.WithServerMiddleware(address, func(m *http.ServeMux) {
    // register routes on m
}, middlewareFunc)
```

Wraps the HTTP handler with middleware (e.g. `metric.MiddlewareHandler` for Prometheus request metrics).

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

- **Sentry recover** — `recover()` only works in direct defer, can't be wrapped. Use `defer func() { s.RecoverFlush(recover()) }()` before `lifecycle.Run()`.
- **App-specific setup** — database connections, client construction, configuration parsing. All happen before `lifecycle.New()`.
- **Signal blocking** — the caller calls `system.KillSignalBlock()` after `lifecycle.Run()`.
- **Reporter start** — if other components need the sentry hub at construction time, the reporter must `Start()` before `lifecycle.New()` so `Hub()` is available.

## Usage Pattern

```go
func Run(o *option.Config) {
    // 1. Pre-lifecycle setup (sentry, clients, config)
    s := reporter.New(...)
    s.Start()
    defer func() { s.RecoverFlush(recover()) }()

    // 2. Construct components
    e := metric.New(0, o.Verbose, nil)

    // 3. Build lifecycle in desired start order
    b := lifecycle.New(
        lifecycle.WithVerbose(o.Verbose),
        lifecycle.WithWorker(e),
        lifecycle.WithWorker(ticker.New(5*time.Minute, func() { ... })),
        lifecycle.WithServer(":8080", func(m *http.ServeMux) {
            m.HandleFunc("/health", healthHandler)
        }),
    )

    // 4. Run, block, stop
    b.Run()
    defer b.Stop()
    system.KillSignalBlock()
}
```

## Adapting Existing Types

To make a type usable as a `lifecycle.Worker`:

- **Move runtime parameters into constructor.** If `Start(hourly bool)` takes args, move them to `New(..., hourly)` so `Start()` takes none.
- **Move callbacks into options.** If `Start(handler)` takes a callback, add a `WithSubscriber(handler)` option so `Start()` reads from the struct.
- **HTTP servers become route registrars.** If a server type owns its own `*http.Server`, extract a `Setup(m *http.ServeMux)` method and let lifecycle own the HTTP serving via `WithServer`.

## Go-Library Components with Start()/Stop()

These types implement the Worker interface:

| Package                      | Type       | What it does                     |
|------------------------------|------------|----------------------------------|
| `pkg/metric`                 | `Server`   | Prometheus metrics HTTP endpoint |
| `pkg/ticker`                 | `Ticker`   | Periodic function execution      |
| `pkg/errors/sentry/reporter` | `Reporter` | Sentry hub lifecycle + flush     |
