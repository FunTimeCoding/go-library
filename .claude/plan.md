# golint / goanalyze feature plan

---

## Batch 1 — Naming bans: clear abbreviations

All unambiguous — abbreviation → short var or full word. Ready to implement.

| Banned | Use instead |
|--------|-------------|
| `pos` | `p` (var) or `position` |
| `buf` | `b` (var) or `buffer` |
| `ptr` | `p` (var) or `pointer` |
| `addr` | `a` (var) or `address` |
| `ref` | `r` (var) or `reference` |
| `nav` | `n` (var) or `navigation` |

Note: `ref` and `req` both suggest `r` — that's fine, they won't appear in the same scope.

---

## Batch 2 — Naming bans: `prev`

`prev` is a cut-off abbreviation. Proposed replacement: `past` (not `previous` — too long, and still abbreviates to `prev`).

**Open question:** Does `past` read naturally in all contexts where `prev` appears? E.g. `pastState`, `pastValue`, `past []*alert.Alert` — check actual usages in codebase before committing.

---

## Batch 3 — Naming bans: markup/serialization group

`yaml`, `xml`, `html`, `json` all hold serialized data. Proposed alternatives:

- `yaml`, `xml`, `html` → `m` (var) or `markup`
- `json` → `j` (var) or `notation` (consistent with `pkg/notation` already in the codebase)

**Open question:** Should all four collapse to the same `markup`/`m`, or does `json` deserve its own track (`notation`/`j`) given the existing `notation` concept? And does the codebase actually use these as variable names, or mainly as import aliases (where the no-acronym rule already applies)?

---

## Batch 4 — Naming bans: `configuration` → `option`

Current state: `config`/`cfg` are banned → suggest `c`/`configuration`.

Proposed addition: `configuration` itself is too long for owned code — use `option` instead (the codebase already uses `option` structs for this pattern).

**Open questions:**
- The linter can't distinguish "owned code wrapping an external API that uses the word configuration" from "our own configuration type." How do we handle that? Possible approach: ban `configuration` only in *type names* and *parameter names*, and only when it refers to our own structs — but that's hard to automate.
- If `configuration` is banned, the `config`/`cfg` → `configuration` suggestion becomes a dead end. The chain would need to become `config`/`cfg` → `option` directly — is that too big a leap?
- Alternative: leave `configuration` as the intermediate suggestion (covers external API wrapping) and rely on code review for the final `option` rename in owned code.

---

## Batch 5 — Naming bans: `decl`

`decl` appears frequently in Go AST/analysis code (`ast.Decl`, `genDecl`, etc.). Banning it may cause significant friction in `pkg/lint/` itself.

**Needs investigation:** How often does `decl` appear in the codebase, and in what contexts? If it only appears in analyzer internals where `declaration` would be genuinely clearer, the ban is worth it. If it's pervasive and always obviously means "declaration," the rename buys little.

---

## Linter features — deferred

| Item | Reason |
|------|--------|
| Direct `exec.Command` usage (suggest `pkg/system/run`) | Easy, deferred |
| String `+` concatenation ban (suggest `fmt.Sprintf` or `filepath.Join`) | Needs AST pass |
| Flag which items were auto-fixed in output | UX, format not decided |
| Stub test filename `main_test.go` enforcement in `pkg/tool/<program>` | Medium complexity |
| Multiple bare top-level `var x = ...` → suggest `var (...)` | New checker |
