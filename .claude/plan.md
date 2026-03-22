# golint / goanalyze feature plan

---

## Linter features — deferred

| Item | Tool | Notes |
|------|------|-------|
| Multiple bare top-level `var x = ...` → suggest `var (...)` | golint | New text-scanner checker |
| New banned segments: `href` (→ link/reference/locator), `def` (→ definition), `concat` (→ concatenate), `obj` (→ object), `stmt` (→ statement) | goanalyze / naming | DONE |
| Ban `&pkg.Struct{}` when `pkg` is owned by this project — enforce `pkg.New()` instead | goanalyze | DONE |
| Function parameters typed as struct or `*struct` must have single-letter names | goanalyze | Applies to all params incl. receivers; next-letter rule for collisions |
