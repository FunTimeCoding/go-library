# golint feature plan

See `doc/plan/goanalyze.md` for the goanalyze naming fix plan.

golint is text/string-based. goanalyze uses the Go analysis package (AST, type-checking,
scope). New rules that need scope or type information belong in goanalyze.

## Remaining work

### Auto-fixers

- `err` → `e` variable auto-fix: variable checker already detects; add `ChangedLine` fix
  path. Scope-aware rename within file.
- `defer f.Close()` → `defer errors.PanicClose(f)`: text match and rewrite. Catches real
  unchecked-close bugs.

### New rules

- `var` grouping: consecutive bare `var x = ...` → `var (...)` — text-based auto-fixer
- stdlib call replacement: `os.Create` etc. → go-library wrappers — call site easy,
  import rewrite harder
- Filename must match function/struct name — preventive; 95%+ compliance today
- Blob executable finder: scan tracked files for compiled binaries (by magic bytes or
  file mode). Images (png etc.) are excluded. Flag for deletion before damage is done.

### Parked — manual conformance batches

- Run golint stray_const, decide per-constant where it belongs
- Grep for stdlib calls that have go-library wrappers (before rules exist)
