# goanalyze: fix correctness and naming improvements

## Completed

- **Struct field fix**: `checkNaming` now uses `v.IsField()` to exclude struct fields from
  single-letter suggestions. `Ref string` → `Reference string`, not `R string`.
- **Inner-scope shadowing**: `resolveFix` now checks child scopes recursively via
  `childScopeContains` before accepting a fix. Prevents renames that would be shadowed
  by inner declarations (e.g., `val → v` when a nested block declares `var v`).
- **New bans**: Added `src → s, source` and `dst → d, destination`.
- **Test consolidation**: Merged 8 separate `TestFix*` functions into one `TestFix` with
  `t.Run` subtests. Pipeline runs once instead of 8×.
- **Unit tests**: Added table-driven tests for `chooseFix`, `replaceSegment`, `segments`,
  `segmentSpans`, and `containsSegment` (22 cases). Total test count: 31.
- **Smoke test on go-mint**: 61 files, all renames correct including shadow fallback.
  `dstDirectory` correctly aborts (collision with existing `destinationDirectory` in scope).

## Planned — suggestions data structure rethink

The current `map[string][]string` mixes single-letter and multi-word suggestions in one
slice. Consider separating them to make special-case handling clearer:

- Single-letter suggestions have collision sensitivity (receiver letters, common locals)
- Multi-word suggestions are the "safe" fallback for fields, types, non-variables
- Some entries only have single-letter (`tx → t`, `ctx → x`) — these can't fix non-variables
  at all unless we add a multi-word alternative
- Receiver occupying any letter means ANY suggested first-letter could collide — the
  inner-scope check + fallback logic handles this, but the data structure could be cleaner

## Planned — deterministic local variable renaming

Eventually: goanalyze re-renames ALL local variables in a function according to a
deterministic pattern derived at runtime from the suggestions slice. This is the end-state
but requires the inner-scope and collision machinery to be solid first.
