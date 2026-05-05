# Entrypoint Spec

Shared conventions for all `cmd/` programs - linker variables, `Main()`, reporter integration.

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

These are injected by `gobuild` at compile time - see `build.md`.

## Main Function

`Main()` is the first real function in every program. It always:

1. Creates the reporter (see below)
2. Registers flags and calls `monitor.ParseBind`
3. Delegates to the tool-specific logic (`Check()`, `Run()`, etc.)

```go
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    r := reporter.NewEnvironment("go<tool>", version)
    r.Start()
    defer func() { r.RecoverFlush(recover()) }()

    pflag.Int(argument.Port, web.ListenPort, web.PortUsage)
    monitor.ParseBind(version, gitHash, buildDate)
    o := option.New()
    o.Port = argument.RequiredInteger(argument.Port)
    Run(o, r)
}
```

`monitor.ParseBind` adds `--version` flag, calls `argument.ParseBind()`, and exits on `--version`. Argument registration order: domain-specific `pflag` calls first, then `monitor.ParseBind()`. After parsing, populate the option struct using `argument.RequiredInteger` (or `argument.RequiredString`, `environment.Required`, etc.).

## Reporter Integration

Every program creates a reporter at the top of `Main()`, before any other
work. The reporter captures unhandled panics and provides error reporting
to all downstream components.

```go
r := reporter.NewEnvironment("go<tool>", version)
r.Start()
defer func() { r.RecoverFlush(recover()) }()

// ... flag parsing, option construction
Run(o, r)
```

- Reporter is always created - empty locator produces noop behavior
  internally (no branching, no nil checks)
- Reporter is the first thing in `Main()` so the defer runs last
- `RecoverFlush` captures the panic, flushes to Sentry if configured,
  prints the panic value, and exits
- The reporter is passed as a separate `Run()` parameter, not on the
  option struct (option structs hold configuration, not constructed
  dependencies)
- `Run()` accepts `face.Reporter` (the interface), not the concrete
  `*reporter.Reporter`

See `three-pillars.md` for the full wiring pattern including logger and
recovery middleware.

### `os.Exit` and reporter

Some tools call `os.Exit(1)` for expected failure conditions (no results, validation failures, upload errors). This intentionally bypasses the reporter defer - these are not crashes and should not be reported as errors. The reporter covers unexpected panics only.
