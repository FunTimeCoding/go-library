# goanalyze: naming fix correctness and improvements

## Completed

- **Struct field fix**: `v.IsField()` excludes struct fields from single-letter suggestions
- **Inner-scope (child) shadowing**: recursive `childScopeContains` prevents renames
  shadowed by inner declarations
- **Parent-scope collision**: `parentScopeContains` walks `scope.Parent()` upward,
  prevents renames that shadow outer variables
- **Two-list suggestions**: `suggestion{letters, words}` struct replaces flat slice;
  fallback tries explicit letters → word characters → abort
- **New bans**: `src → s, source` and `dst → d, destination`
- **Test consolidation**: single `TestFix` with 9 subtests, pipeline runs once
- **Unit tests**: table-driven for `chooseFix`, `replaceSegment`, `segments`,
  `segmentSpans`, `containsSegment`, `parentScopeContains`; 10 integration subtests,
  41 total
- **Smoke tests on go-mint**: parent-scope shadow correctly avoided in `format.go`
- **Unloaded file references**: text-based fallback pass for exported renames in files
  excluded by build constraints (e.g. `//go:build linux`). Walks all `.go` files, skips
  those already handled by the type-checked pipeline. Regression test with tagged fixture.

## Remaining work

### Collision detection hardening

Prevent renames that shadow imported package names (e.g. renaming to `fmt`) — compiler
catches it but produces confusing errors. Also: when `--fix` renames an exported identifier,
verify the new name doesn't collide with other exports in the same package.

### Struct attribute naming

Struct fields should not get single-letter suggestions beyond the current `IsField()`
exclusion. Needs a stricter rule for field-level suggestions.

### Widen naming bans

Survey codebases for more abbreviations. Nearly free: one map entry per ban, automatic
cross-package fix via `--fix`. New entries to add to `suggestions.go`:

- `cmd → c, command`
- `diff / diffs → d, difference`
- `auth → a, authentication`
- `enum → e, enumeration`
- `api → i, interface` (guard against keyword collision)
- `json`: add `n` as second letter alongside existing `j` — two-letter proposal: `{j, n}`

### Auto-fix for other analyzers

- **forbidden_call**: `exec.Command` → `run.Command` + import rewrite
- **string_concatenation**: `a + b` → `fmt.Sprintf(...)` (needs expression-to-source)
- **struct_literal**: `&pkg.X{}` → `pkg.NewX()` (needs constructor existence check)

### String-literal constant finder

New analyzer: scan non-test code for string literals that have a matching constant
defined in a nearby constant package (`pkg/constant`, `pkg/<thing>/constant`,
`pkg/tool/<thing>/constant`) but use the literal directly instead. First phase:
flag and report. Second phase: propose new constants for repeated literals, with
an exclude-list for patterns that don't warrant constants.

### Deterministic local variable renaming

End-state: goanalyze re-renames ALL local variables in a function according to a
deterministic pattern derived at runtime from the suggestion lists. Requires collision
machinery and two-list structure to be solid first.

Feeds into this:
- Count unique local variables per function (foundation for the algorithm)
- Collapse variable: if a local is declared once and used exactly once, not in a
  loop, flag it for inlining
- Error variable sequencing: when `e` is re-declared multiple times in a function,
  rename subsequent declarations to `f`, `g`, `h` — guarding against collision with
  existing `f` (file), `h` (header), etc. Algorithm needs a survey of actual textures
  across the two 100k+ LOC projects before the letter reservation can be finalized
