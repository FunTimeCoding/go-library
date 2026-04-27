# Lint pipeline

## Pipeline overview

Three tools, run by `task lint` in this order:
1. **golint** — text/string-based fast checks (`golint --fix` applies safe rewrites).
2. **gofix** — type-aware code rewrites. Two modes: `gofix` (naming-segment substitutions
   like `ok → okay`) and `gofix --rename` (variable letter renames).
3. **goanalyze** — AST/type/scope-aware analyzers. Report-only.

When AST precision matters, rules belong in goanalyze instead of golint
(e.g. `forbidden_import` migrated to avoid false positives on string literals).

## goanalyze analyzers

Registered in `multichecker.Main`:
- `naming` — flags banned identifier segments, suggests replacements
- `forbidden_call` — flags banned function calls (`exec.Command`)
- `string_concatenation` — flags string `+` concatenation
- `struct_literal` — flags `&pkg.X{}` for owned packages
- `call_format` — flags multi-arg calls where arguments share a line
- `defer_close` — flags `defer x.Close()` where x implements `io.Closer`
- `file_identity` — one identity per file, filename matches identity in snake_case
- `string_constant` — flags string literals where a matching constant exists nearby
- `type_receiver` — flags packages with more than one type that has method receivers
- `forbidden_import` — flags banned package imports (`flag`, testify)
- `variable_naming` — deterministic type-based letter assignment (gated behind `--rename`)

## Source-of-truth map

**Shared between analyzer and fixer:**
- `pkg/lint/analyzer/variable_naming` — letter rules, scope/descendant analysis,
  `ComputeRenames` API. Imported by both `goanalyze` (via `Analyzer.Run` for
  diagnostics) and `gofix` (via `ComputeRenames` for edits). Single source.

**NOT shared (duplicated, sync manually):**
- `pkg/lint/analyzer/naming/naming.go` `suggestions` + `noSuggestion` maps vs
  `pkg/tool/gofix/suggestion.go` `suggestions` + `noSuggestion` maps. Two
  banned-segment lists. Adding a rule requires touching both files; missing one
  silently breaks auto-fix while analyzer still reports. Backlog: extract.

## Variable naming design

### Convergent determinism

Original variable names are noise; output is purely a function of
`(type, lexical position)`. Different AI sessions name the same variable
differently — the analyzer rewrites them all to the same canonical form.
Cross-file pattern recognition becomes type-driven, not name-driven.

Stability is structural: `assignLetters` uses `sort.SliceStable` with
deterministic tie-breakers, so same input always produces same output across
Go versions and runs.

### Eligibility

A variable enters the rename pool if:

| Variable kind | Eligibility |
|---|---|
| Local (`:=`, `var`, `range`) | Single-char (any type) OR error type (any length) |
| Parameter | Primitive type: single-char only. Non-primitive: any length |
| Receiver | Always (single-letter, type-derived) |

Exemptions (applied to non-primitive multi-char params):
- **Paired type**: 2+ params with same `types.Type.String()` in the same
  function — keep all (e.g., `before *Issue, after *Issue`).
- **Primitive slice**: `[]string`, `[]int` etc. — keep (descriptive name
  typically carries information beyond the type).

Tautological names (e.g. `users []User`, `items []ChecklistItem`) get
normalized to the type's preferred letter (`v` for slice of struct) — the
plural-of-element pattern adds nothing the type doesn't already say.

### Scope priority sort

Variables are processed in this order:
1. **scopePriority** ASC: receiver (0), parameter (1), local (2)
2. **precedence** ASC (within scope-priority tier)
3. **position** ASC (within precedence tier)

Receivers and parameters anchor first; locals fall through afterward via
`passTaken`. Result: in `func (h *Router) M(w http.ResponseWriter, q *http.Request) { var r CommentRequest; ... }`, receiver claims `r` (preferred for "Router"),
param `w` claims `w`, param `q` falls through to `e` (`r` taken), local `r`
falls through to `u` (`r` taken in `passTaken`, `var r` will be renamed too).

### Letter assignment

Each eligible variable gets a letter via `pickLetter(type, taken)`:
1. Try preferred letters from the type's letter list (table below) in order.
2. Fall through to a-z exhaustive when all preferred letters are taken.

The `taken` set is built per-variable as:

| Component | kindLocal | kindParameter / kindReceiver |
|---|---|---|
| `passTaken` | included | included |
| `scopedNames` (scope-chain walk) | included | filtered: skip names in `localNames` |
| `descendantNames` (forward block walk) | included | not used |

`localNames` is the set of names of all `kindLocal` variables in the function.
Filtering them out of receiver/param `scopedNames` is what lets receivers/params
claim their preferred letter regardless of body-local conflicts; the body
locals process later and yield via `passTaken`.

### Shadow protection

`scopedNames` walks the variable's scope chain (own scope + ancestors via
`o.Parent()`). Self-exclusion only at the variable's own scope. Prevents
same-scope re-declaration build breaks (e.g., `:=` redeclaring a closure
parameter with a different type).

`descendantNames` walks the variable's enclosing `*ast.BlockStmt` forward from
the variable's position, recursing into nested blocks. Used only for
`kindLocal`. Prevents an outer local rename from creating a shadow with a
later-introduced name (e.g., a closure parameter inside the body).

For receivers/parameters: `descendantNames` is intentionally NOT consulted.
Letting body locals block receiver/param letter claims defeats the
type-anchored convention. Body locals are renamed around the receiver/param
via `passTaken`.

### Letter cycles vs single letters

Most types use cycles (letters of the type's name) instead of single letters.
This gives a deterministic fallback within type-meaning before resorting to
generic a-z:

| Type | Cycle | Origin |
|---|---|---|
| float | `f, l, o, a, t` | "float" |
| byte | `b, y, t, e` | "byte" |
| map | `m, a, p` | "map" |
| string | `s, t, r, i, n, g` | "string" |
| anonymous struct | `s, t, r, u, c` | "struct" |
| named struct | first letter of last word, then cycle | type name (last word) |
| interface | first letter of name, then cycle | type name |

Single-letter types (`error`, `int`, `bool`, `chan`, `*testing.T`, etc.) have
no cycle — they fall straight to a-z when preferred letter is taken.

### Pointer-to-struct slice classification

`[]*Struct` is classified as struct slice (preferred letter `v`), not
primitive slice. Implementation peels the pointer in the slice element check:

```
if s, ok := underlying.(*types.Slice); ok {
    elem := s.Elem()
    if ptr, ok := elem.(*types.Pointer); ok {
        elem = ptr.Elem()
    }
    if _, ok := elem.Underlying().(*types.Struct); ok {
        return precedenceStructSlice
    }
    ...
}
```

### Precedence table

| Precedence | Type | Letters (in order) |
|------------|------|--------------------|
| 1 | `error` | `e`, `f`, `g`, `h` |
| 2 | `*testing.T` | `t` |
| 3 | `*testing.B` | `b` |
| 4 | `io.Reader` | `r` |
| 5 | `io.Writer` | `w` |
| 6 | `context.Context` | `x` |
| 7 | `*os.File` | cycle "file": `f`, `i`, `l`, `e` |
| 8 | `*gzip.Writer` | `z` |
| 9 | `*tar.Writer` | `t` |
| 10 | `string` | cycle "string": `s`, `t`, `r`, `i`, `n`, `g` |
| 11 | `int` / `int64` etc | `i` |
| 12 | `float64` / `float32` | cycle "float": `f`, `l`, `o`, `a`, `t` |
| 13 | `bool` | `b` |
| 14 | `byte` / `[]byte` | cycle "byte": `b`, `y`, `t`, `e` |
| 15 | `map[K]V` | cycle "map": `m`, `a`, `p` |
| 16 | `chan T` | `c` |
| 17 | slice of struct (incl. `[]*Struct`) | `v` |
| 18 | slice of primitive | same as element type |
| 19 | named struct | first letter of last word, cycle through last word |
| 20 | anonymous struct | cycle "struct": `s`, `t`, `r`, `u`, `c` |
| 21 | interface | first letter of name, cycle through name |

All types fall back to `a-z` exhaustive when preferred letters are taken.

### Comma-ok idiom

The 2nd LHS of a comma-ok form must be `okay` (not `ok` — `ok` is banned as
an acronym; `okay` is the canonical form). Detected: type assertion `x.(T)`,
map index `m[k]`, channel receive `<-c`. NOT general function returns (those
use `err` per the error rules).

Per-scope numbering: each non-overlapping scope's first ok-form is `okay`.
When two ok-forms share the same scope, they become `okay, okay1, okay2, ...`.
Refactoring two ok-forms into sibling scopes resets the count.

## Remaining work — gofix

### Variable naming

~~- Flag multi-character descriptive names — done for params (any length non-primitive)
  and errors (any length). Locals are still single-char only.~~
~~- Handle `ok`/`okay` booleans — done via comma-ok idiom enforcement.~~
~~- Receiver validation — done.~~

- Multi-char local rename: locals are still gated to single-char only. Long
  descriptive locals (`result`, `letter`, `counter`) could be normalized to
  the type's preferred letter. Risky — descriptive locals often carry
  intentional meaning. Consider per-name pattern (e.g. `result := ...; return result`
  is a tell-tale that result is just `r`/`v`).

### Collision detection hardening

~~Keyword guard: `token.IsKeyword` check in `resolve_fix.go` — done.~~

~~Unloaded reference fix: `fixFileReferences` now parses with `go/parser` and skips
`*ast.SelectorExpr` targets — done.~~

~~Shadow prevention: scopedNames walks the variable's scope chain (parent walk),
descendantNames walks forward through the enclosing block. Build-break and
shadow scenarios both prevented. Done.~~

When `--fix` renames an exported identifier, verify the new name doesn't
collide with other exports in the same package.

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

## Remaining work — golint

~~- forbidden import check — migrated to goanalyze `forbidden_import` analyzer~~

- stdlib call replacement: `os.Create` etc. → go-library wrappers — call site easy,
  import rewrite harder

### Auto-fix messaging

When gofix renames variables, the output should make it obvious the change was
applied automatically — like golint's `(auto-fixed)` suffix. Currently models
reading the output think the renames are suggestions.

### Function signature formatting

Two rules:
- **Expand**: multi-arg single-line signatures → one argument per line when above
  a threshold (mirrors `call_format` but for declarations). Must also split
  shorthand types: `a, b string` → `a string,\nb string` (each param gets its
  own type annotation)
- **Collapse**: single-arg expanded signatures → single line when only one parameter

### Parked — manual conformance batches

- Run golint stray_const, decide per-constant where it belongs
- Grep for stdlib calls that have go-library wrappers (before rules exist)

## Backlog

- File rename to follow identity rename: when gofix renames a function or
  struct that is the sole identity in its file (per the `file_identity`
  convention), rename the file too (e.g., `check_ok.go → check_okay.go`).
  Eliminates the manual `mv` step after segment-substitution renames.
- Single source of truth for naming rules: `pkg/lint/analyzer/naming` and
  `pkg/tool/gofix` each maintain their own `suggestions` map, `noSuggestion`
  map, and multi-segment exemption logic. Adding a rule requires touching
  both; missing one silently breaks auto-fix while analyzer still reports
  (hit twice in this session — `ok` and `fn`). Extract shared data and
  predicates into a single package consumed by both.
- Suggest `strings.ToIntegerStrict` (or `ToIntegerStrict64`) instead of
  `strconv.Atoi` unless dedicated error handling is required. Eliminates
  the `xFail/yFail`-style error-cycling cascades in convert/sort closures.
- Anonymous struct eradication: scout the codebase for anonymous struct
  usages, decide if any are legitimate (likely none — usually a sign of
  missing type definition or library leakage). If clean, add an analyzer
  that bans anonymous struct declarations.
- Move all naming-related fix concerns into `gofix` only:
  - **variable_naming dry-run out of `goanalyze`**: variable_naming
    diagnostics swarm `task lint` output; logic already lives in
    `gofix --rename`. Remove the analyzer entry from `multichecker.Main`
    and drop the `--rename` CLI flag from goanalyze (point users to
    `gofix --rename --diff`).
  - **Strip `SuggestedFix` from the `naming` analyzer**: analyzer
    becomes pure-report (just diagnostics, no fix attachments). gofix
    is the sole place fixes are computed and applied.
  - Side effect: collapses `segment.ResolveFix` (parent-only,
    analysistest-friendly) and `segment.ResolveFixDeep` (full scope)
    into one. The split exists only as a workaround for analysistest's
    multi-pass file loading making deep-walk results differ between
    diagnostic and verification passes. No more analyzer-side fix
    means no more analysistest concern; gofix uses the single deep
    implementation.
  - Side effect: deletes all `*.go.golden` files under
    `pkg/lint/analyzer/naming/testdata/` — analyzer tests just check
    diagnostic messages, no post-fix verification needed.
  - Keeps lint focused on truly broken patterns; makes gofix the single
    home for all naming concerns.
- goanalyze single-file invocation gives misleading `undefined: X` errors.
  Detect single-file pattern and either expand to enclosing directory or
  refuse with a clear message ("provide a package directory pattern").
