# goanalyze: naming fix correctness and improvements

## Analyzers

Current analyzers registered in `multichecker.Main`:
- `naming` — flags banned identifier segments, suggests replacements
- `forbidden_call` — flags banned function calls (`exec.Command`)
- `string_concatenation` — flags string `+` concatenation
- `struct_literal` — flags `&pkg.X{}` for owned packages
- `call_format` — flags multi-arg calls where arguments share a line
- `defer_close` — flags `defer x.Close()` where x implements `io.Closer`
- `variable_naming` — deterministic type-based letter assignment for local variables

Fix paths: `--fix` (naming renames + call formatting), `--rename` (variable naming)

## Remaining work

### Collision detection hardening

Prevent renames that shadow imported package names (e.g. renaming to `fmt`) — compiler
catches it but produces confusing errors. Also: when `--fix` renames an exported identifier,
verify the new name doesn't collide with other exports in the same package.

### Struct attribute naming

Struct fields should not get single-letter suggestions beyond the current `IsField()`
exclusion. Needs a stricter rule for field-level suggestions.

### Widen naming bans

Survey codebases for more abbreviations. One map entry per ban, automatic
cross-package fix via `--fix`. Candidates:

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
