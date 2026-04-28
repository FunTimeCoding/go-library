# Three Pillars

Every service wires three infrastructure concerns in a consistent
order. Together they ensure operational visibility, panic capture,
and graceful degradation.

## The Pillars

| Pillar | Where constructed | Where threaded |
|--------|-------------------|----------------|
| Sentry hub | `Main()` via reporter | `Run()` param → workers, recovery middleware |
| Logger | `Run()` via `logger.New(ctx)` | Workers, lifecycle |
| Recovery | `Run()` via middleware + worker patterns | HTTP servers, worker loops |

## Wiring Order

`Main()` creates the reporter and extracts the hub. `Run()` receives
the hub as a parameter (not on the option struct — option structs hold
configuration, not constructed dependencies). `Run()` creates the
logger and wires everything into lifecycle.

```go
func Main(version, gitHash, buildDate string) {
    var h *sentry.Hub

    if c := environment.Optional(sentryConstant.LocatorEnvironment); c != "" {
        r := reporter.New("goexample", c, "", version)
        r.Start()
        defer func() { r.RecoverFlush(recover()) }()
        h = r.Hub()
    }

    monitor.ParseBind(version, gitHash, buildDate)
    o := option.New()
    // ... populate option fields
    Run(o, h)
}
```

```go
func Run(o *option.Config, h *sentry.Hub) {
    l := logger.New(context.Background())
    lifecycle.New(
        lifecycle.WithLogger(l),
        lifecycle.WithWorker(worker.New(l, h)),
        lifecycle.WithServerMiddleware(
            web.AddressPort(o.Port),
            func(m *http.ServeMux) { /* routes */ },
            web.RecoveryMiddleware(h),
        ),
    ).RunUntilSignal()
}
```

## Recovery Layers

Three layers, from outermost to innermost:

1. **Entrypoint** (`Main()`): `defer func() { r.RecoverFlush(recover()) }()`
   catches panics from the entire program. See `entrypoint.md`.

2. **HTTP middleware** (`web.RecoveryMiddleware`): wraps the HTTP mux.
   Panics from handlers are caught, reported via `hub.Recover(v)`, and
   converted to 500 responses. Wired via `WithServerMiddleware` or
   `WithProtectedServerMiddleware`. See `lifecycle.md`.

3. **Worker loops** (`withRecovery`): each worker wraps per-iteration
   work in a recovery defer. Panics are reported to Sentry and the
   worker continues. See `error-handling.md`.

MCP handlers do not need per-handler recover defers — the mcp-go
framework handles recovery internally.

## Server Type Selection

| Server type | Use when |
|-------------|----------|
| `WithServerMiddleware` | Streaming, MCP, long-running requests |
| `WithProtectedServerMiddleware` | Plain REST (adds 10s read/write timeout) |

Mixed servers (REST + MCP on same mux) use `WithServerMiddleware` —
the streaming endpoint governs the timeout choice.

Both variants accept a middleware parameter for `web.RecoveryMiddleware`.

## What Each Component Receives

| Component | Logger | Hub | Why |
|-----------|--------|-----|-----|
| Lifecycle workers | Yes (constructor param) | Yes (for withRecovery) | Workers need both for recovery logging |
| HTTP route handlers | Via recovery middleware | Via recovery middleware | Middleware catches panics |
| MCP tool handlers | Not directly | Via tool struct (deferred) | Tier 2 errors use hub.CaptureException |
| Startup one-shot tasks | Yes (from Run) | Yes (for withRecovery if looping) | Same recovery pattern as workers |
