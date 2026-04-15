# golint feature plan

See `doc/plan/goanalyze.md` for the goanalyze naming fix plan.

golint is text/string-based. goanalyze uses the Go analysis package (AST, type-checking,
scope). New rules that need scope or type information belong in goanalyze.

## Remaining work

### Auto-fixers

- `err` → `e` variable auto-fix: variable checker already detects; add `ChangedLine` fix
  path. Scope-aware rename within file.

### New rules

- stdlib call replacement: `os.Create` etc. → go-library wrappers — call site easy,
  import rewrite harder
- Filename must match function/struct name — preventive; 95%+ compliance today
- ~~Blob executable finder~~ — done (b5304c8f)
- ~~`defer f.Close()` → `defer errors.PanicClose(f)`~~ — moved to goanalyze as
  `defer_close` analyzer (needs type info to check `io.Closer` conformance)
- ~~`var` grouping~~ — done (76190a3c)

### Parked — manual conformance batches

- Run golint stray_const, decide per-constant where it belongs
- Grep for stdlib calls that have go-library wrappers (before rules exist)
