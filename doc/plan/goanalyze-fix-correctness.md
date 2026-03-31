# goanalyze: fix correctness and naming improvements

## Completed

- **Struct field fix**: `checkNaming` uses `v.IsField()` to exclude struct fields from
  single-letter suggestions. `Ref string` → `Reference string`, not `R string`.
- **Inner-scope (child) shadowing**: `resolveFix` checks child scopes recursively via
  `childScopeContains` before accepting a fix. Prevents renames that would be shadowed
  by inner declarations (e.g., `val → v` when a nested block declares `var v`).
- **New bans**: Added `src → s, source` and `dst → d, destination`.
- **Test consolidation**: Merged 8 separate `TestFix*` functions into one `TestFix` with
  `t.Run` subtests. Pipeline runs once instead of 8×.
- **Unit tests**: Table-driven tests for `chooseFix`, `replaceSegment`, `segments`,
  `segmentSpans`, and `containsSegment` (22 cases). Total test count: 31.
- **Smoke tests on go-mint**: Two rounds — confirmed shadow fallback works, confirmed
  `dstDirectory` correctly aborts (collision with existing `destinationDirectory`).
- **Parent-scope collision check**: `scopeContains` now also walks `scope.Parent()` chain
  upward (stopping before universe scope) via `parentScopeContains`. Prevents renames
  that shadow outer variables (e.g., `src → s` when outer `s` exists in parent scope).
- **Two-list suggestions restructure**: Replaced `map[string][]string` with
  `map[string]suggestion` where `suggestion` has separate `letters` and `words` fields.
  `checkNaming` appends letters only when `allowSingleLetter`, then words always.
  Fallback in `resolveFix`: try explicit letters → walk first word's characters → abort.
  Renamed `firstMultiCharacter` to `firstWord`.

## Done — parent-scope collision check

**Problem**: `scopeContains` only checks the same scope and child scopes. It does NOT
check parent scopes. This produces broken code:

```go
s := status.New(...)           // outer scope: s is a status builder
for _, src := range p.Sources { // src lives in for-scope
    s.DetailLink(src, ...)      // uses outer s
}
// goanalyze renames src → s, producing:
for _, s := range p.Sources {   // shadows outer s!
    s.DetailLink(s, ...)         // BROKEN: calls DetailLink on a string
}
```

**Fix**: Before accepting a rename, walk `scope.Parent()` upward (stopping at the
package scope) and check `parent.Lookup(candidate)` at each level. If the candidate
name exists in any ancestor scope, it would shadow that binding — treat as collision.

Implementation:
- Add `parentScopeContains(scope, name)` — walks `scope.Parent()` chain, checks
  `parent.Lookup(name)` at each level, stops at package scope (or nil)
- Update `scopeContains` to also call `parentScopeContains`
- Add integration test: fixture with outer `s` variable + inner `src` that would rename
  to `s` → should fall back or abort
- Add unit test for `parentScopeContains`
- Rebuild, re-smoke on go-mint

## Done — two-list suggestions restructure

**Problem**: The current `map[string][]string` mixes single-letter and multi-word
suggestions in one flat slice. The fallback logic walks letters of the first multi-char
word, producing confusing results like `dst → s` (from "destination"[2]).

**New structure**: Separate letters and words explicitly:

```go
type suggestion struct {
    letters []string   // single-letter options: ["d"]
    words   []string   // full-word options: ["destination"]
}
```

**Revised collision fallback**:
1. Try `letters[0]` (for single-segment names) or `words[0]` (for multi-segment)
2. If collision: try remaining letters in order
3. If all explicit letters collide: walk characters of `words[0]` as last resort
4. If everything collides: abort

The "walk word letters" fallback stays as a last resort — it's deterministic, which is
the goal. Explicit letters just get priority.

**Benefits**:
- Entries like `tx → t` or `ctx → x` (letter-only) are structurally clear
- Entries like `args → arguments` or `dirs → directories` (word-only) are clear too
- `checkNaming` can filter on `letters` vs `words` instead of checking `len(alternative)`
- Easier to add multiple letter alternatives per segment without them being confused
  with word suggestions

Migration: mechanical — each current slice entry becomes either a letter or a word based
on its length. No behavioral change except the explicit ordering.

## Future — deterministic local variable renaming

End-state: goanalyze re-renames ALL local variables in a function according to a
deterministic pattern derived at runtime from the suggestion lists. Requires the
collision machinery (parent + child scope checks) and the two-list structure to be
solid first. Not yet scoped.
