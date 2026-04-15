# Deterministic variable naming

goanalyze `variable_naming` analyzer in `pkg/lint/analyzer/variable_naming/`.
Assigns canonical single-letter names to local variables based on their type.
Fix path via `goanalyze --rename ./...` (separate from `--fix`).

## Eligibility

Currently flags and fixes:
- Single-letter variables that are the wrong letter for their type
- Error variables of any name length (`err`, `err2`, etc.)

## Letter assignment precedence

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

All types fall back to `a-z` exhaustive when their preferred letters are taken.

## Remaining work

- Flag multi-character descriptive names (`result`, `letter`) for deterministic shortening
- Handle `ok`/`okay` booleans from type assertions and map lookups
- Receiver validation (already correct across the codebase, just needs a check)

## Not in scope

- Function parameters (descriptive names per conventions spec)
- Struct fields (exported API naming)
- Package-level variables
- Return value naming
