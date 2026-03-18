# Entrypoint Spec

Shared conventions for all `cmd/` programs — linker variables, `Main()`, sentry integration.

## Linker Variables

Every `cmd/<name>/main.go` declares linker variables and delegates to `Main()`:

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

These are injected by `gobuild` at compile time — see `build.md`.

## Main Function

`Main()` is the first real function in every program. It always:

1. Sets up sentry (see below)
2. Registers flags and calls `monitor.ParseBind`
3. Delegates to the tool-specific logic (`Check()`, `Run()`, etc.)

```go
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
        r := reporter.New("go<tool>", c, "", version)
        r.Start()
        defer func() { r.RecoverFlush(recover()) }()
    }

    monitor.ParseBind(version, gitHash, buildDate)
    // ... flags, option struct, delegate
}
```

`monitor.ParseBind` adds `--version` flag, calls `argument.ParseBind()`, and exits on `--version`. Argument registration order: reusable monitor helpers first, then domain-specific `pflag` calls, then `monitor.ParseBind()`.

## Sentry Integration

Every program sets up sentry at the top of `Main()`, before any other work. This captures unhandled panics via `SENTRY_LOCATOR` environment variable.

```go
if c := environment.Optional(sentry.LocatorEnvironment); c != "" {
    r := reporter.New("go<tool>", c, "", version)
    r.Start()
    defer func() { r.RecoverFlush(recover()) }()
}
```

- Sentry is the first thing in `Main()` so the defer runs last (captures panics from all later code)
- `environment.Optional` returns empty string when the variable is unset — no error, no Sentry
- `reporter.RecoverFlush(recover())` captures the panic value, reports it, and flushes before exit
- Uses the `version` parameter directly — no need to pass it through option structs

### `os.Exit` and sentry

Some tools call `os.Exit(1)` for expected failure conditions (no results, validation failures, upload errors). This intentionally bypasses the sentry defer — these are not crashes and should not be reported as errors. Sentry covers unexpected panics only.
