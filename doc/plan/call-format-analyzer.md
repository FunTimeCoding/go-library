# call_format analyzer

New goanalyze analyzer that enforces one-argument-per-line in function/method calls.

## Rule

A `CallExpr` with 2+ arguments is a violation when either:
- all arguments are on one line but the line exceeds 80 characters (including leading indentation)
- the call is already multi-line but two or more arguments share a line

Single-line calls at or under 80 characters are left alone.

## Fix

Rewrite so each argument sits on its own line, indented one level deeper than the call. Closing `)` on its own line at the caller's indent level. Run gofmt on the result.

## Scope

Function and method call arguments only. Not return statements, not composite literal elements, not function definitions.

## Files

All under `pkg/lint/analyzer/call_format/`.

- [x] `call_format.go` — `Analyzer` variable
- [x] `run.go` — inspector walk over `*ast.CallExpr`, detect violations
- [x] `check.go` — line-length and argument-sharing checks, `IsViolation` exported for fix path
- [x] `fix.go` — `BuildFixes` exported, computes indent from source bytes
- [x] `analyzer_test.go` — `analysistest.Run` with testdata
- [x] `testdata/src/example/` — compliant, long_single_line, shared_line, suppressed

## Integration

- [x] Register in `pkg/tool/goanalyze/main.go` alongside the other analyzers in `multichecker.Main`
- [x] Fix path: `run_call_format_fix.go` in goanalyze, uses `applyEdits`/`splice` pipeline
- [x] Manually tested: single-line >80, multi-line shared args, first-arg-on-paren-line — all fix correctly

## Out of scope

- Return statements
- Composite literals
- Function/method definitions (IDE handles these)
- Single-argument calls (long strings are a different problem)
