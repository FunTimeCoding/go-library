# Claude Code

## Specs

Design and coding specs live in `.claude/spec/`:

- `conventions.md` - coding style, error handling, naming, structure
- `naming.md` - banned identifier segments, replacement patterns, type/field naming rules
- `entrypoint.md` - linker variables, Main(), sentry setup (shared by all cmd/ programs)
- `service-tool.md` - long-running service tool pattern (Run, lifecycle wiring, routes)
- `check-tool.md` - CLI check tool pattern (fetch, filter, format, print)
- `lifecycle.md` - lifecycle manager for servers and workers
- `entity-wrapper.md` - entity wrapper pattern
- `console-status.md` - fluent status line builder
- `testing.md` - integration testing patterns (mocks, lifecycle HTTP, store)
- `build.md` - gobuild cross-compilation and linker variable convention
- `taskfile.md` - task runner, git hooks, CI pipeline
- `locator.md` - fluent URL builder (`pkg/web/locator`)
- `openapi.md` - OpenAPI codegen pattern (server/, client/, route/ structure)
- `error-handling.md` - error handling strategies: default panic, MCP translation, flow control exception, store method rule

Read the relevant spec before working in that area.

## Project Structure

- `cmd/` - service entry points
- `pkg/` - library and service packages
- `.claude/spec` - specs
