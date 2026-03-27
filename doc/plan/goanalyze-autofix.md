## Discovery: multichecker already supports -fix

`multichecker.Main()` delegates to `internal/checker.Run()` which:
1. Registers `-fix` and `-diff` flags via `checker.RegisterFlags()`
2. After running analyzers, collects `SuggestedFixes` from diagnostics
3. Applies `TextEdit`s to source files (or prints unified diff with `-diff`)

No custom runner needed. Analyzers just need to switch from
`p.Reportf(pos, msg)` to `p.Report(analysis.Diagnostic{..., SuggestedFixes: [...]})`.

## Analyzer auto-fix feasibility

### naming — identifier rename

Switch `p.Reportf` to `p.Report` with `SuggestedFixes` containing `TextEdit`s.

Challenge: renaming a declaration also requires renaming every reference.
Within a single package (one analysis pass), this is feasible:
- `p.TypesInfo.Defs` maps declaration identifiers to `types.Object`
- `p.TypesInfo.Uses` maps reference identifiers to `types.Object`
- For a given banned identifier, find all idents whose Object matches
- Emit one `TextEdit` per occurrence

Scope:
- Local variables, parameters, receivers: all references in same package ✓
- Exported identifiers: references in other packages won't be fixed ✗
  (acceptable — report without fix, or fix with model-assisted batch)

Segment replacement logic:
- `segments()` already splits `fooUrl` into `["foo", "url"]`
- Need `replaceSegment(name, old, new)` → `fooUrl` + `url→locator` = `fooLocator`
- Must handle camelCase, snake_case, and preserve casing conventions

### forbidden_call — call site rewrite

- `exec.Command(args...)` → `run.Command(args...)`
- `exec.CommandContext(x, args...)` → `run.CommandContext(x, args...)`
- Single `TextEdit` replacing `exec.Command` selector with `run.Command`
- Import change: remove `"os/exec"`, add `"pkg/system/run"`
  (import edits are additional `TextEdit`s in the same `SuggestedFix`)

### string_concatenation — expression rewrite

- `a + b` → `fmt.Sprintf("%s%s", a, b)`
- `s += v` → `s = fmt.Sprintf("%s%s", s, v)`
- Need to render operand expressions as source text for the Sprintf args
- Multi-operand chains: `a + b + c` — only the outermost reports, so
  all operands are collected at once

### struct_literal — constructor call

- `&pkg.X{}` → `pkg.NewX()` — only when constructor exists and takes no args
- `new(pkg.X)` → `pkg.NewX()` — same constraint
- Must verify constructor existence (check if `NewX` is exported from package)
- When struct has field values (`&pkg.X{A: 1}`), skip — not auto-fixable

## New bans to add (with auto-fix from the start)

| Banned | Suggestions | Notes |
|--------|-------------|-------|
| `var` (the abbreviation, not the keyword) | `v`, `variable` | e.g. `isVar` → `isVariable` or `v` |
| `const` (the abbreviation, not the keyword) | `c`, `constant` | e.g. `isConst` → `isConstant` or `c` |

The naming analyzer's `segments()` function splits on camelCase and `_`
boundaries, so `isVar` → segments `["is", "var"]`. The `var` segment matches
the ban; the Go keyword `var` is a statement, never an identifier, so no
false positives from the keyword itself.

## Implementation order

### Phase 1 — infrastructure + naming (highest leverage)

Files to create:
- `pkg/lint/analyzer/naming/replace_segment.go` — `replaceSegment(name, old, new string) string`
  Handles camelCase (`fooUrl` → `fooLocator`), snake_case (`foo_url` → `foo_locator`),
  preserves leading uppercase for exported names (`FooUrl` → `FooLocator`)
- `pkg/lint/analyzer/naming/replace_segment_test.go` — unit tests for the above
- `pkg/lint/analyzer/naming/find_references.go` — `findReferences(p *analysis.Pass, o types.Object) []*ast.Ident`
  Iterates p.TypesInfo.Defs and p.TypesInfo.Uses, returns all idents whose Object == o
- `pkg/lint/analyzer/naming/find_references_test.go`

Files to change:
- `pkg/lint/analyzer/naming/check.go` — switch `p.Reportf()` → `p.Report(analysis.Diagnostic{...})`
  with SuggestedFixes containing TextEdits for the declaration + all references.
  When suggestion is ambiguous (multiple alternatives), use the first multi-char
  alternative for the fix (single-letter requires human judgment on scope).
- `pkg/lint/analyzer/naming/naming.go` — add `"var": {"v", "variable"}` and
  `"const": {"c", "constant"}` to the suggestions map
- `pkg/lint/analyzer/naming/analyzer_test.go` — switch `analysistest.Run` →
  `analysistest.RunWithSuggestedFixes`

Test fixtures:
- Existing `.go` testdata files continue to work (// want comments unchanged)
- Add `.golden` files alongside each fixture showing expected fixed output
  (framework applies all SuggestedFixes and compares against .golden)
- Add new fixtures for `var`/`const` bans

### Phase 2 — integrate gofix into goanalyze

Replace `multichecker.Main()` with a custom runner that:
1. Loads all packages via `go/packages.Load` with `LoadSyntax` + `Tests: true`
2. Runs all four analyzers per-package (same diagnostics as today)
3. With `--fix`: does cross-package renames for exported naming violations
4. Dedup by `token.Pos` to handle test package duality
5. Skip interface methods (`isInterfaceMethod`) and generated files
6. Filter edits to working directory only (no build cache)

Delete `cmd/gofix/` and `pkg/tool/gofix/` after merge.
Move needed logic into `pkg/lint/analyzer/naming/` or a new `pkg/lint/fix/` package.

### Phase 3 — integration test for cross-package rename

Test that creates a temp Go module with multiple packages, runs the
fix logic, and verifies output files. Validates:
- Unexported identifiers renamed within package
- Exported identifiers renamed across packages
- Interface methods skipped
- Generated files skipped
- References in test files included

### Phase 4 — forbidden_call + string_concatenation + struct_literal

- forbidden_call: add SuggestedFixes, import rewriting
- string_concatenation: expression-to-source rendering
- struct_literal: constructor existence check
