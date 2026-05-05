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

`Main()` is the first real function in every program. There are two
shapes depending on whether the tool uses flat flags or subcommands.

### Choosing between flat flags and subcommands

**Flat flags** (pflag + viper + `monitor.ParseBind`): daemons and
single-purpose tools. Daemons should always use flat flags — they
take a port and maybe a config path, nothing more. Standalone tools
that do one thing also use flat flags.

**Subcommands** (cobra): tools with multiple distinct operations.
Typically CLI clients that talk to a daemon (gopostgres, gonetbox,
gohabitica), but also standalone tools that grew past a single verb.

### Flat-flag Main (daemons, single-purpose tools)

```go
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    r := reporter.New(constant.Name, version)
    r.Start()
    defer func() { r.RecoverFlush(recover()) }()

    pflag.Int(argument.Port, web.ListenPort, web.PortUsage)
    monitor.ParseBind(version, gitHash, buildDate)
    o := option.New()
    o.Port = argument.RequiredInteger(argument.Port)
    o.Version = version
    Run(o, r)
}
```

`monitor.ParseBind` adds `--version` flag, calls `argument.ParseBind()`, and exits on `--version`. Argument registration order: domain-specific `pflag` calls first, then `monitor.ParseBind()`. After parsing, populate the option struct using `argument.RequiredInteger` (or `argument.RequiredString`, `environment.Required`, etc.).

### Subcommand Main (multi-operation tools)

```go
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    r := reporter.New(constant.Name, version)
    r.Start()
    defer func() { r.RecoverFlush(recover()) }()
    c := client.NewEnvironment()
    root := &cobra.Command{
        Use:     constant.Name,
        Version: fmt.Sprintf("%s (%s %s)", version, gitHash, buildDate),
    }
    root.AddCommand(listItems(c))
    root.AddCommand(createItem(c))

    if f := root.Execute(); f != nil {
        errors.Printf("%v", f)
        os.Exit(1)
    }
}
```

No `monitor.ParseBind` — cobra handles `--version` via the `Version`
field. No option struct — each subcommand owns its own flags.

Each subcommand lives in its own file and returns a `*cobra.Command`:

```go
func createItem(c *client.Client) *cobra.Command {
    var name string
    result := &cobra.Command{
        Use:   "create-item [value]",
        Short: "Create a new item",
        Args:  cobra.ExactArgs(1),
        Run: func(
            _ *cobra.Command,
            arguments []string,
        ) {
            fmt.Println(c.Create(arguments[0], name))
        },
    }
    result.Flags().StringVar(
        &name,
        "name",
        "",
        "Item name",
    )

    return result
}
```

Subcommand flags use `result.Flags().StringVar` (bound to local
variables), not global `pflag` registration. Required flags use
`result.MarkFlagRequired`.

## Reporter Integration

Every program creates a reporter at the top of `Main()`, before any other
work. The reporter captures unhandled panics and provides error reporting
to all downstream components.

```go
r := reporter.New(constant.Name, version)
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
