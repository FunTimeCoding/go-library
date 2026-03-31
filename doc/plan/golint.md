# golint / goanalyze feature plan

## golint — auto-fixers with immediate impact

| Item | Notes |
|------|-------|
| `err` → `e` variable auto-fix | Variable checker already detects; add `ChangedLine` fix path. Scope-aware rename within file. |
| `defer f.Close()` → `defer errors.PanicClose(f)` | Text match and rewrite. Catches real unchecked-close bugs. |

## golint — new rules and conformance

| Item | Notes |
|------|-------|
| `var` grouping: consecutive bare `var x = ...` → `var (...)` | Text-based auto-fixer |
| stdlib call replacement: `os.Create` etc. → go-library wrappers | Call site easy, import harder |
| Filename must match function/struct name | Preventive; 95%+ compliance today |

## goanalyze

See `doc/plan/goanalyze.md` for naming fix plan and roadmap.

## Parked — manual conformance batches

- Run golint stray_const, decide per-constant where it belongs
- Grep for stdlib calls that have go-library wrappers (before rules exist)
