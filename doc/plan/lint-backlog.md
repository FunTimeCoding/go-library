## Parked: golint text-based auto-fixers

These are feasible within golint's existing line-scanner infrastructure
but are lower priority than the goanalyze auto-fix work.

- `defer f.Close()` → `defer errors.PanicClose(f)` — text match, rewrite
- `var` grouping: consecutive bare `var x = ...` → `var (...)`
- stdlib call replacement: `os.Create`, `os.MkdirAll`, `filepath.Rel`,
  `ReadString`, `Write` → go-library wrappers (call site easy, import harder)

## Parked: new lint rules (report-only)

- Filename must match function/struct name (preventive; 95%+ compliance today)
- Struct/`*struct` params must have single-letter names
- Spec references per lint violation (infra change to message format)

## Parked: manual conformance batches

These need human judgment per violation. The linter finds them, a model
or developer fixes them in grouped batches.

- Run goanalyze naming against full codebase, group by banned segment
- Run golint variable checker (`err` → `e`), fix scope-aware
- Run golint stray_const, decide per-constant where it belongs
- Grep for stdlib calls that have go-library wrappers (before rules exist)
- Run both tools against `pkg/lint/` itself (dogfooding)
