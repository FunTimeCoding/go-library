# Check Tool Spec

Pattern for CLI tools that fetch, filter, and display entities from external systems.

## Package Structure

```
cmd/go<tool>/
└── main.go                         # Linker vars, delegates to Main()

pkg/tool/go<tool>/
├── main.go                         # Main(): register arguments, parse, build option, call Check
├── main_test.go                    # Stub test
└── ...

pkg/<domain>/check/<entity>/
├── check.go                        # Check(o): orchestrates collect → filter → format → print
├── collect.go                      # collect(): private, fetches entities from client
├── print_notation.go               # printNotation(): JSON output for monitoring
└── option/
    ├── <entity>.go                 # Option struct (named after entity, not "Option")
    └── new.go                      # Factory: New() *<Entity>

pkg/<domain>/constant/
└── constant.go                     # Format preset (option.Color.Copy() + domain tags)

pkg/monitor/
├── parse_bind.go                   # ParseBind(version, gitHash, buildDate)
├── copyable_argument.go            # CopyableArgument(): reusable --copyable flag
├── notation_argument.go            # NotationArgument(): reusable --notation flag
├── all_argument.go                 # AllArgument(): reusable --all flag
├── verbose_argument.go             # VerboseArgument(): reusable --verbose flag
└── item/constant/constant.go       # Item identifiers for monitor.NoRelevant()

pkg/argument/
└── constant.go                     # Argument name constants (Copyable, Notation, All, etc.)
```

## Entry Point

`cmd/go<tool>/main.go` declares linker variables and delegates:

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

## Main Function

`Main()` registers flags, parses with `monitor.ParseBind`, builds the option struct, and calls `Check`:

```go
func Main(
    version string,
    gitHash string,
    buildDate string,
) {
    monitor.CopyableArgument()
    monitor.NotationArgument()
    monitor.AllArgument()
    monitor.ParseBind(version, gitHash, buildDate)
    o := option.New()
    o.Copyable = viper.GetBool(argument.Copyable)
    o.Notation = viper.GetBool(argument.Notation)
    o.All = viper.GetBool(argument.All)
    <entity>.Check(o)
}
```

`monitor.ParseBind` adds `--version` flag, calls `argument.ParseBind()`, and exits on `--version`. Argument registration order: reusable monitor helpers first, then domain-specific `pflag` calls, then `monitor.ParseBind()`.

## Option Struct

Plain struct with bool/string fields. No methods beyond `New()`.

```go
package option

type Issue struct {
    Notation bool
    All      bool
    Copyable bool
}
```

Factory in `new.go`:

```go
func New() *Issue {
    return &Issue{}
}
```

## Check Function

`Check(o *option.Option)` is the canonical name. It orchestrates:

1. Collect entities
2. Handle notation mode (early return)
3. Build format from domain preset
4. Apply option flags to format
5. Iterate and print

```go
func Check(o *option.Issue) {
    elements := collect()

    if o.Notation {
        printNotation(elements, o)

        return
    }

    f := constant.Format

    if o.Copyable {
        f.Tag(tag.Copyable)
    }

    for _, e := range elements {
        fmt.Println(e.Format(f))
    }

    if len(elements) == 0 {
        monitor.NoRelevant(item.GoJira.Plural)
    }
}
```

## Collect Function

`collect()` is private, extracts data fetching from the check function. Returns a slice of entity pointers.

```go
func collect() []*issue.Issue {
    return query.Open(
        common.Jira(),
        environment.Required(constant.DefaultProjectNameEnvironment),
    )
}
```

## Notation Mode

Every check tool must support `--notation` for monitoring integrations. Notation produces JSON output via `printNotation()` and returns early before the format loop.

`printNotation()` takes a slice of entities and the option struct. It follows a fixed structure:

```go
func printNotation(
    v []*issue.Issue,
    o *option.Issue,
) {
    r := report.New()

    for _, e := range report.Trim(
        v,
        r,
        o.All,
        item.GoJira,
    ) {
        r.AddItem(
            item.GoJira,
            e.MonitorIdentifier,
            constant.Warning,
            e.Title,
            e.Link,
            e.Create,
        )
    }

    r.Print()
}
```

The steps:

1. `report.New()` — create a report
2. `report.Trim(v, r, o.All, item)` — cap the slice unless `--all` is set
3. `r.AddItem(item, identifier, severity, detail, link, create)` — add each entity as a report item
4. `r.Print()` — output JSON

Each entity must expose `MonitorIdentifier`, `Link`, and a timestamp field for `AddItem`. The severity (`constant.Critical`, `constant.Warning`) is domain-specific.

Items are registered in `pkg/monitor/item/constant/constant.go` via `collector.New(command, name, plural)`.

## Format Preset

Each domain defines a base format in its `constant/constant.go`:

```go
var Format = option.Color.Copy()
```

The check function copies this preset and applies option flags. Common flag mappings:

| Option        | Format effect            |
|---------------|--------------------------|
| `Copyable`    | `f.Tag(tag.Copyable)`    |
| `Extended`    | `f.Extended()`           |
| `Fingerprint` | `f.Tag(tag.Fingerprint)` |

## Reusable Argument Helpers

Common arguments live in `pkg/monitor/` as `<Name>Argument()` functions. Each registers a `pflag.Bool` with the constant from `pkg/argument/constant.go`.

| Helper               | Flag         | Description                                        |
|----------------------|--------------|----------------------------------------------------|
| `CopyableArgument()` | `--copyable` | Disable OSC8 links and add a copyable link instead |
| `NotationArgument()` | `--notation` | JSON output                                        |
| `AllArgument()`      | `--all`      | Include filtered items                             |
| `VerboseArgument()`  | `--verbose`  | Verbose output                                     |

Domain-specific flags (e.g., `--critical`, `--set`) use `pflag` directly in the entrypoint.

## Multi-Entity Domains

Some domains have multiple entity types that each warrant their own check tool. Each entity gets a separate entrypoint and check package, sharing the domain client and constants.

```
cmd/go<tool1>/
└── main.go                             # Entrypoint for entity 1
cmd/go<tool2>/
└── main.go                             # Entrypoint for entity 2

pkg/<domain>/
├── client.go                           # Shared client
├── new.go                              # Client factory
└── constant/
    └── constant.go                     # Shared Format preset

pkg/<domain>/check/<entity1>/
├── check.go
├── collect.go
├── print_notation.go
└── option/
    ├── <entity1>.go
    └── new.go

pkg/<domain>/check/<entity2>/
├── check.go
├── collect.go
├── print_notation.go
└── option/
    ├── <entity2>.go
    └── new.go
```

Each check package follows the same structure as a single-entity tool. The key differences:

- **Separate entrypoints**: one `cmd/go<tool>/` per entity type
- **Separate check packages**: `check/<entity1>/` and `check/<entity2>/` are siblings
- **Separate monitor items**: each entity has its own `collector.New()` entry
- **Shared domain**: client, constants, and format preset live in the parent domain package
- **Separate option structs**: named after the entity (`Job`, `Request`, `Outdated`), not the domain

Examples: `goghjob`/`goghpr` (GitHub jobs and pull requests), `goalert`/`gosilence` (Prometheus alerts and silences).

## No Relevant Output

When no entities match, print a standard message:

```go
if len(elements) == 0 {
    monitor.NoRelevant(item.GoJira.Plural)
}
```

Items are registered in `pkg/monitor/item/constant/constant.go`.
