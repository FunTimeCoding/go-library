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

Use everywhere covered by a recovery layer — web handlers, workers, REST route
handlers. Do not return the error to the caller. Do not write `if e != nil`.

```go
errors.PanicOnError(s.store.Save(record))
```

### Flow control exception: return `error`

Use when the error outcome changes what happens next — not just "something went wrong"
but "this specific failure path has distinct handling." The canonical example is a
job worker marking status transitions: a failed status update means the job must be
skipped, which requires the caller to act differently.

```go
if e := w.store.UpdateStatus(job.ID, "processing"); e != nil {
    w.logger.Plain("failed to mark job %d processing: %v", job.ID, e)

    return
}
```

### MCP exception: translate to tool result

MCP handlers have no recovery middleware. Panics would crash the handler without
producing a response. All errors must be translated to `mcp.NewToolResultError`.

Two tiers:

**Input validation** (bad params from the model): return a tool error. No Sentry —
these are model mistakes, not infrastructure failures.

```go
id, f := r.RequireString("id")

if f != nil {
    return mcp.NewToolResultError("id is required"), nil
}
```

**Infrastructure failure** (store, DB, external call): return a generic tool error
AND capture to Sentry. The message to the model is generic — do not leak internal
error details that would confuse the model. The error detail goes to Sentry for
retroactive analysis.

```go
if e := s.store.UpdateRecord(id); e != nil {
    if s.hub != nil {
        s.hub.CaptureException(e)
    }

    return mcp.NewToolResultError("failed to update record"), nil
}
```

The hub must be threaded into the MCP server and nil-guarded (same pattern as web
and workers).

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

## Deviations

**vs idiomatic Go**: idiomatic Go returns `(T, error)` from nearly everything. Here,
error returns are the exception, not the rule. See `conventions.md`.

**vs Claude defaults**: Claude will tend to add `if e != nil` handling everywhere and
thread errors up the call stack. Resist this. When in doubt, `PanicOnError`.
