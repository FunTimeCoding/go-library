# golint / goanalyze feature plan

---

Items moved to `doc/plan/goanalyze-autofix.md` (auto-fix focus) and
`doc/plan/lint-backlog.md` (parked ideas). This file tracks remaining
items that don't fit either plan.

## Linter features — deferred

| Item | Tool | Notes |
|------|------|-------|
| New banned segments: `href` (→ link/reference/locator), `def` (→ definition), `concat` (→ concatenate), `obj` (→ object), `stmt` (→ statement) | goanalyze / naming | DONE |
| Ban `&pkg.Struct{}` when `pkg` is owned by this project — enforce `pkg.New()` instead | goanalyze | DONE |
| Spec references per lint violation — link the relevant spec doc in the lint message | golint / goanalyze | Future: infra change to message format |
