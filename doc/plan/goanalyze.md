# Lint pipeline

golint is text/string-based. goanalyze uses the Go analysis package (AST, type-checking,
scope). New rules that need scope or type information belong in goanalyze.

## goanalyze analyzers

Registered in `multichecker.Main`:
- `naming` — flags banned identifier segments, suggests replacements
- `forbidden_call` — flags banned function calls (`exec.Command`)
- `string_concatenation` — flags string `+` concatenation
- `struct_literal` — flags `&pkg.X{}` for owned packages
- `call_format` — flags multi-arg calls where arguments share a line
- `defer_close` — flags `defer x.Close()` where x implements `io.Closer`
- `variable_naming` — deterministic type-based letter assignment for local variables

Fix paths: `--fix` (naming renames + call formatting), `--rename` (variable naming)

## Variable naming precedence

| Precedence | Type | Letters (in order) |
|------------|------|--------------------|
| 1 | `error` | `e`, `f`, `g`, `h` |
| 2 | `*testing.T` | `t` |
| 3 | `*testing.B` | `b` |
| 4 | `io.Reader` | `r` |
| 5 | `io.Writer` | `w` |
| 6 | `*os.File` | cycle "file": `f`, `i`, `l`, `e` |
| 7 | `*gzip.Writer` | `z` |
| 8 | `*tar.Writer` | `t` |
| 9 | `string` | `s` |
| 10 | `int` / `int64` etc | `i` |
| 11 | `float64` / `float32` | `n` (number, avoids `f` collision) |
| 12 | `bool` | `b` |
| 13 | `byte` / `[]byte` | `b` |
| 14 | `map[K]V` | `m` |
| 15 | `chan T` | `c` |
| 16 | slice of struct | `v` |
| 17 | slice of primitive | same as element type |
| 18 | named struct | first letter of last word, cycle through last word |
| 19 | interface | first letter of name, cycle through name |

All types fall back to `a-z` exhaustive when preferred letters are taken.

Currently eligible: single-letter wrong vars and error variables of any name length.
Not yet eligible: multi-char descriptive names, `ok`/`okay` booleans, receivers.

## Remaining work — goanalyze

### Variable naming

- Flag multi-character descriptive names (`result`, `letter`) for deterministic shortening
- Handle `ok`/`okay` booleans from type assertions and map lookups
- Receiver validation

### Collision detection hardening

Prevent renames that shadow imported package names. When `--fix` renames an exported
identifier, verify the new name doesn't collide with other exports in the same package.

### Struct attribute naming

Struct fields should not get single-letter suggestions beyond the current `IsField()`
exclusion. Needs a stricter rule for field-level suggestions.

### Widen naming bans

One map entry per ban, automatic cross-package fix via `--fix`. Candidates:

- `cmd → c, command`
- `diff / diffs → d, difference`
- `auth → a, authentication`
- `enum → e, enumeration`
- `api → i, interface` (guard against keyword collision)

### Auto-fix for other analyzers

- **forbidden_call**: `exec.Command` → `run.Command` + import rewrite
- **string_concatenation**: `a + b` → `fmt.Sprintf(...)` (needs expression-to-source)
- **struct_literal**: `&pkg.X{}` → `pkg.NewX()` (needs constructor existence check)

### String-literal constant finder

New analyzer: scan non-test code for string literals that match a constant
in a nearby constant package but use the literal directly. First phase: flag.
Second phase: propose new constants for repeated literals.

## Remaining work — golint

- stdlib call replacement: `os.Create` etc. → go-library wrappers — call site easy,
  import rewrite harder
- Filename must match function/struct name — preventive; 95%+ compliance today

### Parked — manual conformance batches

- Run golint stray_const, decide per-constant where it belongs
- Grep for stdlib calls that have go-library wrappers (before rules exist)
