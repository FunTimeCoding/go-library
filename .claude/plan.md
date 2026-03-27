# golint / goanalyze feature plan

## High priority — auto-fixers with immediate impact

| Item | Tool | Notes |
|------|------|-------|
| Widen naming bans — survey codebase for abbreviations, add to suggestions map | goanalyze | Nearly free: one map entry per ban, automatic cross-package fix via `--fix` |
| `err` → `e` variable auto-fix | golint | Variable checker already detects; add `ChangedLine` fix path. Scope-aware rename within file. High frequency. |
| `defer f.Close()` → `defer errors.PanicClose(f)` | golint | Text match and rewrite. Catches real unchecked-close bugs. |

## Medium priority — collision detection hardening

| Item | Tool | Notes |
|------|------|-------|
| Import shadowing check | goanalyze | Prevent renames that shadow imported package names (e.g. renaming to `fmt`). Compiler catches it but produces confusing errors. |
| Cross-package export collision check | goanalyze | When `--fix` renames an exported identifier, verify the new name doesn't collide with other exports in the same package. |

## Medium priority — auto-fix for remaining goanalyze analyzers

| Item | Tool | Notes |
|------|------|-------|
| forbidden_call SuggestedFixes | goanalyze | `exec.Command` → `run.Command` + import rewrite |
| string_concatenation SuggestedFixes | goanalyze | `a + b` → `fmt.Sprintf(...)`. Needs expression-to-source rendering. |
| struct_literal SuggestedFixes | goanalyze | `&pkg.X{}` → `pkg.NewX()`. Needs constructor existence check. |

## Lower priority — new rules and conformance

| Item | Tool | Notes |
|------|------|-------|
| `var` grouping: consecutive bare `var x = ...` → `var (...)` | golint | Text-based auto-fixer |
| stdlib call replacement: `os.Create` etc. → go-library wrappers | golint | Call site easy, import harder |
| Filename must match function/struct name | golint | Preventive; 95%+ compliance today |
| Struct/`*struct` params must have single-letter names | goanalyze | New analyzer |
| Spec references per lint violation | both | Infra change to message format |

## Parked — manual conformance batches

- Run golint stray_const, decide per-constant where it belongs
- Grep for stdlib calls that have go-library wrappers (before rules exist)
