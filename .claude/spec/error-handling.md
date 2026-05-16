# Error Handling

## Principle

Minimize `error` returns and `if e != nil` comparisons. Prefer `errors.PanicOnError`
and rely on recovery chains to capture unexpected failures. Explicit error
handling is reserved for cases where it is structurally required.

## Recovery Chain

Three recovery layers wrap all program execution:

**Entrypoint** (`Main()`): `defer func() { r.RecoverFlush(recover()) }()` captures
panics from the entire program. See `entrypoint.md`.

**HTTP middleware** (`web.RecoveryMiddleware`): wraps the HTTP mux. Panics from web
handlers are caught, reported via `r.Recover(v)`, and converted to 500 responses.

**Worker goroutines**: each worker's scheduled function opens with an explicit deferred
recover. Panics are reported via the reporter and the worker continues running (the
loop is not killed).

```go
defer func() {
    if v := recover(); v != nil {
        w.reporter.Recover(v)
        w.logger.Plain("worker recovered from panic: %v", v)
    }
}()
```

The reporter is never nil - `Main()` always creates one. Empty locator
produces noop behavior. No nil-guards needed.

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

// MCP caller - captures to reporter, returns structured error
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
to Sentry and return a structured error with the event ID. Two layers:

`captureFail` is the primitive - a private method on the Server struct
that wraps `response.CaptureFail`. Takes the error and a model-facing
message. The error goes to sentry, the message goes to the model.

```go
func (s *Server) captureFail(
    e error,
    message string,
) (*mcp.CallToolResult, error) {
    return response.CaptureFail(s.reporter, e, message)
}
```

`captureDetail` is the per-service wrapper that whitelists known error
types. It checks for API-specific typed errors, relays their message
to the model if recognized, and falls back to `constant.UnexpectedError`
for anything unknown. Lives in its own file (`capture_detail.go`).

```go
func (s *Server) captureDetail(e error) (*mcp.CallToolResult, error) {
    if d, okay := errors.AsType[*detail_error.Detail](e); okay {
        return s.captureFail(e, d.Detail)
    }

    return s.captureFail(e, constant.UnexpectedError)
}
```

Each service's `captureDetail` checks for different error types
depending on its API:

- Sentry, Habitica, Jellyfin, Confluence: `*detail_error.Detail`
  from `parseDetail` in the HTTP client
- GitLab: `*gitlab.ErrorResponse` → `e.Message`
- Mattermost: `*model.AppError` → `e.Message`
- NetBox: `*netbox.GenericOpenAPIError` → `common.ExtractMessage`
  parses `detail` string, then field-level validation errors
  (`__all__` and per-field arrays). Shared between REST and MCP
  surfaces via `gonetboxd/common/`.
- Jira/Confluence (goatlassiand): both `*jira.Error` → `ErrorMessages[0]`
  and `*detail_error.Detail`
- Proxmox: sentinel errors via `errors.Is` (`ErrNotFound`,
  `ErrNotAuthorized`, `ErrTimeout`)
- GORM services: `gorm.ErrRecordNotFound` via `errors.Is`

Most handlers call `captureDetail(e)` directly:

```go
result, e := s.client.SearchIssues(...)

if e != nil {
    return s.captureDetail(e)
}
```

Use `captureFail(e, message)` directly only when the handler knows
a specific message that `captureDetail` can't derive - e.g. after
a multi-step operation where the context matters.

**`detail_error.Detail`** (`pkg/web/detail_error/`): a shared typed
error carrying a `Detail` string and `Status` string. HTTP clients
return this from `parseDetail` when the API response contains a
known error message field. The `Detail` field is what gets relayed
to the model.

**`constant.UnexpectedError`** (`pkg/constant/`): the shared fallback
message - `"unexpected error"`. Honest about not knowing what went
wrong. The model gets this alongside the sentry event ID and can
look up the full error if needed. Never lie about the cause - don't
use "API unreachable" or "database unreachable" as catch-all messages.

**`parseDetail`** in HTTP clients: checks the response status code
and parses the error body for a known message field. Each API has
its own shape:

- Sentry: `{"detail": "..."}` (JSON)
- Confluence: `{"message": "..."}` (JSON)
- Habitica: `{"message": "..."}` (JSON)
- Jellyfin: `{"title": "..."}` (ASP.NET ProblemDetails JSON),
  or plain text body, or empty
- Salt: `<p>...</p>` in CherryPy HTML error pages

If the known field is found, returns `detail_error.New(message, status)`.
Otherwise returns `fmt.Errorf("%s", status)`.

`response.CaptureFail` captures the exception via the reporter and
returns structured JSON with `error` and `event_identifier` fields via
`response.FailAny`. The event ID lets the model look up the
stacktrace via Sentry MCP tools and diagnose the problem in the same
conversation.

The reporter is threaded into the MCP server (same pattern as web and
workers). The reporter is never nil - it always exists, even in noop
mode.

Design principle: any `error` value from a function call in an MCP
handler is worth capturing. Even local file I/O errors.
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
reports via the reporter, and lets the loop continue.

```go
func (w *Worker) withRecovery(f func()) {
    defer func() {
        if v := recover(); v != nil {
            w.reporter.Recover(v)
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

The `withRecovery` method lives on each worker struct that holds
reporter + logger. See `three-pillars.md` for where this fits in the
overall recovery layer design.

## HTTP Recovery Middleware

`pkg/web/RecoveryMiddleware` is the shared HTTP recovery layer. It wraps
the mux, catches panics from any handler, reports via `r.Recover(v)`,
and returns 500. Wired into lifecycle via `WithServerMiddleware` or
`WithProtectedServerMiddleware`. See `lifecycle.md`.

### Sentry body enrichment

`pkg/errors/sentry/start.go` installs a `BeforeSend` hook that
checks every error for `face.BodyProvider` (an interface with
`Body() []byte`). When present — e.g. `*netbox.GenericOpenAPIError`
— the response body is attached as a `response` context on the
Sentry event. This works for both `CaptureException` and `Recover`
paths (via `OriginalException` and `RecoveredException` in the
`EventHint`). No caller changes needed — any error satisfying
`BodyProvider` is automatically enriched.

### REST captureDetail variant (gonetboxd)

When a REST API serves both CLI and MCP consumers and needs to surface
meaningful error messages (not just 500), handlers use the non-Must
client methods and call `captureDetail` on error — the same pattern
as MCP handlers. The Server struct holds a `face.Reporter` for Sentry
reporting. Shared error extraction logic lives in a `common/` package
(e.g. `gonetboxd/common/ExtractMessage`). The recovery middleware
remains as a safety net for unexpected panics.

## External Process Failure

When shelling out to an external process (terraform, git, ansible,
yt-dlp, ffmpeg), use `SetReporter` on `run.Run` to enrich Sentry
with stdout and stderr automatically on failure:

```go
c := run.New()
c.Directory = directory
c.SetReporter(r.reporter, "terraform init")
c.Start("terraform", "init")
```

`SetReporter` stores the reporter and a label on the `Run` struct.
When `Start` encounters an error and `Panic` is true (the default),
it calls `CaptureWithContext` with the process output before
panicking. The Sentry event gets stdout and stderr as structured
context - visible in the Sentry UI without being stuffed into the
panic message.

For cases where the caller needs to inspect the error before
deciding what to do (e.g. parsing terraform's JSON output to
decide whether to retry with `-upgrade`), use `NoPanic()` and
handle manually:

```go
c := run.New().NoPanic()
c.Start("terraform", "init", "-json")

if c.Error != nil && needsUpgrade(c.OutputString) {
    // retry with different args
}
```

## Self-Healing File Operations

In batch operations (cleanup loops, file processing), file remove failures
are captured without panicking. The file stays on disk and gets cleaned
up on the next run.

```go
if e := os.Remove(path); e != nil {
    c.reporter.CaptureException(e)
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
