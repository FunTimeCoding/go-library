# Error Handling

## Principle

Minimize `error` returns and `if e != nil` comparisons. Prefer `errors.PanicOnError`
and rely on Sentry recovery chains to capture unexpected failures. Explicit error
handling is reserved for cases where it is structurally required.

## Recovery Chain

Three recovery layers wrap all program execution:

**Entrypoint** (`Main()`): `defer func() { r.RecoverFlush(recover()) }()` captures
panics from the entire program. See `entrypoint.md`.

**HTTP middleware** (`web.RecoveryMiddleware`): wraps the HTTP mux. Panics from web
handlers are caught, reported via `hub.Recover(v)`, and converted to 500 responses.

**Worker goroutines**: each worker's scheduled function opens with an explicit deferred
recover. Panics are reported to Sentry and the worker continues running (the loop is
not killed).

```go
defer func() {
    if v := recover(); v != nil {
        if w.hub != nil {
            w.hub.Recover(v)
        }

        w.logger.Plain("worker recovered from panic: %v", v)
    }
}()
```

All three layers nil-guard the hub: `if hub != nil { ... }` because `SENTRY_LOCATOR`
may be unset in development.

## When to Use Each Strategy

### Default: `PanicOnError`

Use everywhere covered by a recovery layer - web handlers, workers, REST route
handlers. Do not return the error to the caller. Do not write `if e != nil`.

```go
errors.PanicOnError(s.store.Save(record))
```

### Flow control exception: return `error`

Use when the error outcome changes what happens next - not just "something went wrong"
but "this specific failure path has distinct handling." The canonical example is a
job worker marking status transitions: a failed status update means the job must be
skipped, which requires the caller to act differently.

```go
if e := w.store.UpdateStatus(job.ID, "processing"); e != nil {
    w.logger.Plain("failed to mark job %d processing: %v", job.ID, e)

    return
}
```

### Multi-context exception: return error for caller to decide

When a function is called from both HTTP handlers (which PanicOnError) and MCP
handlers (which translate to tool results), it returns `error`. Each caller
handles it according to its context.

```go
// Shared domain function - returns error
func ProcessRecord(path string) (*Result, error) {
    // ... filesystem or network operations that can fail
}

// HTTP caller - panics, caught by recovery middleware
result, e := store.ProcessRecord(path)
errors.PanicOnError(e)

// MCP caller - captures to Sentry, returns structured error
result, e := store.ProcessRecord(path)

if e != nil {
    return s.captureFail(e, "process record failed")
}
```

This is distinct from the flow control exception - the error doesn't change
what happens next (both callers abort on error). The difference is how the
error is surfaced.

### MCP exception: translate to tool result

MCP handlers have no recovery middleware. Panics would crash the handler without
producing a response. All errors must be translated to tool results.

Do not add per-handler recover defers to MCP tools - the mcp-go framework
handles recovery internally.

Two tiers:

**Tier 1 - Input validation** (bad params from the model): use `response.Fail`. No
Sentry - these are model mistakes, not infrastructure failures.

```go
id, f := r.RequireString(parameter.Identifier)

if f != nil {
    return response.Fail("identifier is required: %v", f)
}
```

`response.Fail` wraps `mcp.NewToolResultError` with `fmt.Sprintf` and returns the
standard `(*mcp.CallToolResult, error)` tuple. Use it for all input validation.

**Tier 2 - Infrastructure failure** (store, DB, external call): capture
to Sentry and return a structured error with the Sentry event ID. Use
`captureFail` - a private method on the Server/Tool struct that wraps
`response.CaptureFail`.

```go
result, e := s.store.UpdateRecord(id)

if e != nil {
    return s.captureFail(e, "update record failed")
}
```

`captureFail` lives in its own file (`capture_fail.go`):
```go
func (s *Server) captureFail(
    e error,
    message string,
) (*mcp.CallToolResult, error) {
    return response.CaptureFail(s.hub, e, message)
}
```

`response.CaptureFail` captures the exception to Sentry and returns
structured JSON with `error` and `event_identifier` fields via
`response.FailAny`. The event ID lets the model look up the
stacktrace via Sentry MCP tools and diagnose the problem in the same
conversation.

The hub must be threaded into the MCP server (same pattern as web and
workers). The hub is nil-safe - `CaptureFail` handles nil hubs
gracefully.

Design principle: any `error` value from a function call in an MCP
handler is worth capturing to Sentry. Even local file I/O errors.
`response.Fail` is only for validation where the handler constructs
the error message itself (e.g. "service is required").

## Store Method Rule

Store methods called only from web handlers or workers must use `PanicOnError`
internally and return nothing. Store methods called from MCP handlers must return
`error` so the caller can translate it.

```go
// Only called from web/worker → panic inside, void return
func (s *Store) CreateIfAbsent(v *Record) {
    errors.PanicOnError(
        s.database.Clauses(clause.OnConflict{DoNothing: true}).Create(v).Error,
    )
}

// Called from MCP handler → return error
func (s *Store) UpdateStatus(id string, status string) error {
    return s.database.Model(&Record{}).
        Where("id = ?", id).
        Update("status", status).Error
}
```

Methods called from both web and MCP return `error`. The web caller wraps with
`PanicOnError`; the MCP caller handles explicitly.

## Worker Recovery Pattern

Workers that loop (pollers, watchers, schedulers) wrap per-iteration work
in a `withRecovery` method. This catches panics from a single iteration,
reports to Sentry, and lets the loop continue.

```go
func (w *Worker) withRecovery(f func()) {
    defer func() {
        if v := recover(); v != nil {
            if w.hub != nil {
                w.hub.Recover(v)
            }

            w.logger.Plain("worker recovered from panic: %v", v)
        }
    }()
    f()
}
```

Usage in a loop:

```go
for {
    select {
    case <-ticker.C:
        w.withRecovery(w.poll)
    case <-w.done:
        return
    }
}
```

The `withRecovery` method lives on each worker struct that holds hub +
logger. See `three-pillars.md` for where this fits in the overall
recovery layer design.

## HTTP Recovery Middleware

`pkg/web/RecoveryMiddleware` is the shared HTTP recovery layer. It wraps
the mux, catches panics from any handler, reports via `hub.Recover(v)`,
and returns 500. Wired into lifecycle via `WithServerMiddleware` or
`WithProtectedServerMiddleware`. See `lifecycle.md`.

## External Process Failure

When shelling out to an external process (Python scripts, yt-dlp,
ffmpeg) via `run.New().NoPanic()`, enrich Sentry with the process
output before panicking:

```go
if r.Error != nil {
    if hub != nil {
        hub.WithScope(func(s *sentry.Scope) {
            s.SetContext("process", map[string]any{
                "output": r.OutputString,
                "stderr": r.ErrorString,
            })
            hub.CaptureException(r.Error)
        })
    }

    errors.PanicOnError(r.Error)
}
```

The `SetContext` call attaches stdout and stderr as structured data on
the Sentry event - they're available in the Sentry UI without being
stuffed into the panic message.

## Self-Healing File Operations

In batch operations (cleanup loops, file processing), file remove failures
are captured to Sentry without panicking. The file stays on disk and
gets cleaned up on the next run.

```go
if e := os.Remove(path); e != nil {
    if hub != nil {
        hub.CaptureException(e)
    }
}
```

This is the exception to the PanicOnError default - the failure is
transient, self-healing, and not worth killing the batch over. But it
should still be visible in Sentry.

## Deviations

**vs idiomatic Go**: idiomatic Go returns `(T, error)` from nearly everything. Here,
error returns are the exception, not the rule. See `conventions.md`.

**vs Claude defaults**: Claude will tend to add `if e != nil` handling everywhere and
thread errors up the call stack. Resist this. When in doubt, `PanicOnError`.
