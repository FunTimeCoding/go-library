# Three Pillars

Every service wires three infrastructure concerns in a consistent
order. Together they ensure operational visibility, panic capture,
and graceful degradation.

## The Pillars

| Pillar | Where constructed | Where threaded |
|--------|-------------------|----------------|
| Reporter | `Main()` via `reporter.New` | `Run()` param → workers, recovery middleware, model_context |
| Logger | `Run()` via `logger.New(ctx)` | Workers, lifecycle |
| Recovery | `Run()` via middleware + worker patterns | HTTP servers, worker loops |

## Wiring Order

`Main()` creates the reporter unconditionally. Empty locator produces
noop behavior - no branching, no nil. `Run()` receives the reporter
as `face.Reporter` (not on the option struct - option structs hold
configuration, not constructed dependencies). `Run()` creates the
logger and wires everything into lifecycle.

```go
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    r := reporter.New(constant.Identity.Name(), version).Start()
    defer func() { r.RecoverFlush(recover()) }()

    a := argument.NewInstance(constant.Identity)
    // ... register flags
    a.Parse(version, gitHash, buildDate)
    o := option.New()
    // ... populate option fields
    Run(o, r)
}
```

```go
func Run(o *option.Config, r face.Reporter) {
    l := logger.New(context.Background())
    lifecycle.New(
        l,
        lifecycle.WithWorker(worker.New(l, r)),
        lifecycle.WithServerMiddleware(
            web.AddressPort(o.Port),
            func(m *http.ServeMux) { /* routes */ },
            web.RecoveryMiddleware(r),
        ),
    ).RunUntilSignal()
}
```

## Recovery Layers

Three layers, from outermost to innermost:

1. **Entrypoint** (`Main()`): `defer func() { r.RecoverFlush(recover()) }()`
   catches panics from the entire program. See `entrypoint.md`.

2. **HTTP middleware** (`web.RecoveryMiddleware`): wraps the HTTP mux.
   Panics from handlers are caught, reported via `r.Recover(v)`, and
   converted to 500 responses. Wired via `WithServerMiddleware` or
   `WithProtectedServerMiddleware`. See `lifecycle.md`.

3. **Worker loops** (`withRecovery`): each worker wraps per-iteration
   work in a recovery defer. Panics are reported via `r.Recover(v)`
   and the worker continues. See `error-handling.md`.

MCP handlers do not need per-handler recover defers - the mcp-go
framework handles recovery internally.

## Server Type Selection

| Server type | Use when |
|-------------|----------|
| `WithServerMiddleware` | Streaming, MCP, long-running requests |
| `WithProtectedServerMiddleware` | Plain REST (adds 10s read/write timeout) |

Mixed servers (REST + MCP on same mux) use `WithServerMiddleware` -
the streaming endpoint governs the timeout choice.

Both variants accept a middleware parameter for `web.RecoveryMiddleware`.

## What Each Component Receives

| Component | Logger | Reporter | Why |
|-----------|--------|----------|-----|
| Lifecycle workers | Yes (constructor param) | Yes (for withRecovery) | Workers need both for recovery logging |
| HTTP route handlers | Via recovery middleware | Via recovery middleware | Middleware catches panics |
| MCP tool handlers | Not directly | Via Server struct + captureFail | Tier 2 errors use captureFail → response.CaptureFail |
| Startup one-shot tasks | Yes (from Run) | Yes (for withRecovery if looping) | Same recovery pattern as workers |
