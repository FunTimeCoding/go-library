# Claude Code

## Specs

Design and coding specs live in `.claude/spec/`:

- `conventions.md` — coding style, error handling, naming, structure
- `service-tool.md` — long-running service tool pattern (Main/Run, lifecycle wiring, routes)
- `check-tool.md` — CLI check tool pattern (fetch, filter, format, print)
- `lifecycle.md` — lifecycle manager for servers and workers
- `entity-wrapper.md` — entity wrapper pattern
- `console-status.md` — fluent status line builder
- `testing.md` — integration testing patterns (mocks, lifecycle HTTP, store)
- `build.md` — gobuild cross-compilation and linker variable convention
- `taskfile.md` — task runner, git hooks, CI pipeline

Read the relevant spec before working in that area.

## Project Structure

- `cmd/` — service entry points
- `pkg/` — library and service packages
- `.claude/spec` — specs
