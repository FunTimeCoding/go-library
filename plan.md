---
title: Sentry rollout and spec rework
---

## Done

- Lint rule: `missing_sentry` in `pkg/lint/`
- Proof of concept: `goalert`
- Bulk rollout: 35 standard check/CLI tools
- Decision: `os.Exit(1)` for expected failures is fine, intentionally bypasses sentry defer

## 1) Unify sentry placement in Main()

Move sentry from `Run()` to `Main()` in both daemons so the convention is universal:
- `goalertlogd` — move block from `run.go` to `main.go`, drop `o.Version`
- `gomaintlogd` — move block from `run.go` to `main.go`, drop `o.Version`
- Remove `Version` field from both option structs

## 2) Spec rework

Extract shared entrypoint conventions into `entrypoint.md`:
- Linker vars, `Main()` signature, `monitor.ParseBind`
- Sentry block (always top of `Main()`)
- `os.Exit` guidance: acceptable for expected failures, bypasses sentry defer

Update existing specs to reference `entrypoint.md` instead of duplicating:
- `service-tool.md` — remove entry point, sentry sections
- `check-tool.md` — remove entry point, Main() boilerplate
- `build.md` — remove "Entry Point Contract" section

## 3) Remaining edge cases

### `os.Exit(1)` bypasses defer — 3 files
- `gopackage` — add sentry to Main()
- `goupload` — add sentry to Main()
- `gorenovate` — add sentry to Main()

### `system.Exitf` — 2 files
- `goupdate` — check if system.Exitf calls os.Exit, add sentry to Main()
- `goversion` — check if system.Exitf calls os.Exit, add sentry to Main()

### Parameter name `v` instead of `version` — 1 file
- `godebian` — use `v` in `reporter.New()` call

### No `pkg/tool/` package — 4 cmd/ programs
- `gobump` — sentry block in `cmd/gobump/main.go` directly
- `gobundle` — sentry block in `cmd/gobundle/main.go` directly
- `goc` — sentry block in `cmd/goc/main.go` directly
- `gogw2` — sentry block in `cmd/gogw2/main.go` directly
