# golint / goanalyze feature plan

---

## Linter features — deferred

| Item | Tool | Notes |
|------|------|-------|
| Multiple bare top-level `var x = ...` → suggest `var (...)` | golint | New text-scanner checker |
| New banned segments: `href` (→ link/reference/locator), `def` (→ definition), `concat` (→ concatenate), `obj` (→ object), `stmt` (→ statement) | goanalyze / naming | DONE |
| Ban `&pkg.Struct{}` when `pkg` is owned by this project — enforce `pkg.New()` instead | goanalyze | DONE |
| Function parameters typed as struct or `*struct` must have single-letter names | goanalyze | Applies to all params incl. receivers; next-letter rule for collisions |
| `defer f.Close()` → suggest `defer errors.PanicClose(f)` | golint | errcheck surface: bare defer close |
| `os.Create(path)` unchecked → suggest `system.Create(path)` | golint | errcheck surface: unchecked create |
| `os.MkdirAll(...)` unchecked → suggest `system.MakeDirectory(...)` | golint | errcheck surface: unchecked mkdir |
| `filepath.Rel(...)` unchecked → suggest `system.RelativePath(...)` | golint | errcheck surface: unchecked rel |
| `r.ReadString(...)` unchecked → suggest `system.ReadLine(r)` | golint | errcheck surface: unchecked ReadString |
| `f.Write(b)` unchecked → suggest `_, e = f.Write(b); errors.PanicOnError(e)` | golint | errcheck surface: unchecked Write |
| Spec references per lint violation — link the relevant spec doc in the lint message | golint / goanalyze | Not yet possible; future capability once tool supports it |
