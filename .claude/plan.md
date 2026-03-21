# golint / goanalyze feature plan

Feature requests for the next implementation pass, with design decisions recorded.

---

## Status

All planned fixes from the previous pass are complete:

- Fix 1 - `isTopLevelDecl` restricted to `func` and `type` only (done)
- Fix 2 - blank before `//` comment inside function emits immediately, no buffering (done)
- Top-level var/const blank removal - spurious blanks between consecutive top-level
  `var`/`const` declarations are now flagged (`extraneous_top_level_blank`) and removed (done)

---

## Out of scope (next pass)

| Item | Reason |
|------|--------|
| Direct `exec.Command` usage (flag, suggest `pkg/system/run`) | Easy, deferred to keep this pass focused |
| String `+` concatenation ban | Needs AST pass, not reliably detectable by text scanning |
| Flag which items were auto-fixed in output | Minor UX; output format not yet decided |
| Stub test filename `main_test.go` enforcement for `pkg/tool/<prog>` | Medium complexity; would change stub generation behavior |
| Multiple bare top-level `var x = ...` declarations → suggest `var (...)` grouping | New checker, not spacing |
