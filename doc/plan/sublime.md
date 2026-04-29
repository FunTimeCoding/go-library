# Sublime Text MCP

Three layers: Python bridge plugin, Go client library, Go MCP server.

## System map

```
Sublime Text ← plugin_host (Python 3.8) ← py/sublime_bridge/
    ↕ HTTP 127.0.0.1:6123
pkg/sublime/          - Go client library, entity-wrapper pattern
    ↕
pkg/tool/gosublmcp/   - MCP server (default :8080, --port override)
    ↕ streamable HTTP
Claude Code / CLI / web consumers
```

Bridge plugin symlinked into Sublime:
`~/Library/Application Support/Sublime Text/Packages/SublimeBridge/ → py/sublime_bridge/`

## Non-obvious decisions

- Client methods return errors (not panic) - diverges from other domain clients
  because MCP is the primary consumer and future CLI/web callers benefit too
- Bridge edits use `view.find()` / `view.find_all()` with `sublime.LITERAL` for matching
- Edits require a `TextCommand` subclass (`SublimeBridgeReplaceCommand`) because
  Sublime only allows buffer mutations inside an `edit` context
- `.python-version` file with `3.8` required - without it Sublime loads the plugin
  under Python 3.3 which lacks `ThreadingHTTPServer`
- Bridge imports use try/except for relative vs absolute - Sublime's plugin loader
  resolves imports differently depending on context

## Deferred

- `get_selections` - cursor positions and selection ranges
- `set_selections` - move cursor / set selection ranges
- `find_in_view` - regex search within buffer, return matches with positions
- `run_command` - execute arbitrary Sublime command (escape hatch)
- `get_project` - project folders and settings
- Character-offset editing: `PUT /views/{id}/region`, `EditViewRegion()`, `edit_view_region` tool
- Event notifications: SSE/websocket from bridge, subscriber in `pkg/sublime/` (not MCP layer)
