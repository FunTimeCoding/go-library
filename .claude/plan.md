# golint / goanalyze feature plan

---

## Linter features — deferred

| Item | Reason |
|------|--------|
| Direct `exec.Command` usage (suggest `pkg/system/run`) | Easy, deferred |
| String `+` concatenation ban (suggest `fmt.Sprintf` or `filepath.Join`) | Needs AST pass |
| Flag which items were auto-fixed in output | UX, format not decided |
| Stub test filename `main_test.go` enforcement in `pkg/tool/<program>` | Medium complexity |
| Multiple bare top-level `var x = ...` → suggest `var (...)` | New checker |
