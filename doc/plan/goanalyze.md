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
  `segmentSpans`, `containsSegment`, `parentScopeContains` (40 total)
- **Smoke tests on go-mint**: parent-scope shadow correctly avoided in `format.go`

## Planned — collision detection hardening

- **Import shadowing check**: prevent renames that shadow imported package names
  (e.g. renaming to `fmt`). Compiler catches it but produces confusing errors.
- **Cross-package export collision check**: when `--fix` renames an exported identifier,
  verify the new name doesn't collide with other exports in the same package.

## Planned — widen naming bans

Survey codebases for more abbreviations to add to the suggestions map. Nearly free:
one map entry per ban, automatic cross-package fix via `--fix`.

## Planned — auto-fix for other goanalyze analyzers

- **forbidden_call**: `exec.Command` → `run.Command` + import rewrite
- **string_concatenation**: `a + b` → `fmt.Sprintf(...)` (needs expression-to-source)
- **struct_literal**: `&pkg.X{}` → `pkg.NewX()` (needs constructor existence check)

## Future — deterministic local variable renaming

End-state: goanalyze re-renames ALL local variables in a function according to a
deterministic pattern derived at runtime from the suggestion lists. Requires the
collision machinery and two-list structure to be solid first. Not yet scoped.
