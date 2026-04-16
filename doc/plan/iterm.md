# iTerm MCP

MCP server for collaborative terminal work — shared awareness of what's
happening, reading output together, sending commands step by step.

## System map

```
iTerm2 ← websocket → py/iterm_bridge/ (standalone Python process)
    ↕ HTTP 127.0.0.1:6124
pkg/iterm/            — Go client library
    ↕
pkg/tool/goitermmcp/  — MCP server (default :8080, --port override)
    ↕ streamable HTTP
Claude Code / CLI / web consumers
```

Bridge is a standalone process, not embedded in iTerm:
`cd py/iterm_bridge && uv run main.py`

Requires iTerm2 Python API enabled: Preferences > General > Magic > Enable Python API.

## Non-obvious decisions

- Bridge runs outside iTerm (connects via websocket), unlike Sublime where
  the bridge runs inside the editor as a plugin
- `uv run` with inline script metadata for zero-pollution Python dependency
  management — no venv activation, no system pip littering
- `send_text` does NOT append enter — two-step confirmation pattern for safety.
  User confirms the text, then confirms enter separately
- `set_tab_color` works via session profile properties (`LocalWriteOnlyProfile`),
  not directly on the tab object — iTerm API quirk
- `list_sessions` surfaces `path` (cwd), `jobName`, `commandLine` from iTerm's
  variable system — enables seeing SSH context and working directory without
  reading the screen
- `path` requires iTerm shell integration (installed by default on macOS)
- The `_find_session` helper uses a cached `connection.app` reference set at
  startup, while `list_sessions` and `create_tab` call `async_get_app` fresh
- HTTP handler dispatches async iTerm API calls via `run_coroutine_threadsafe`
  back to the main asyncio event loop
- Import fallback pattern (try relative, except absolute) needed because
  `uv run` executes as a script, not a module

## Safety model

Read operations (list, read) are safe for auto-accept.
Write operations (send) should go through tool confirmation — a misunderstanding
could send destructive commands to production servers.
Tool descriptions hint at risk level so Claude exercises caution.

## Deferred

- `close_session` — close a tab/pane
- `split_pane` — split horizontal/vertical (`async_split_pane`)
- `move_to_window` — move tab to its own window (`async_move_to_window`)
- `select_pane` — navigate to adjacent pane (`async_select_pane_in_direction`)
- `read_screen` raw ANSI mode (currently always stripped)
- tmux -CC control mode integration (expose tmux panes as native sessions)
- Event notifications: output changed, new session created
- `send_sequence` — one-shot a sequence of inputs (e.g. ctrl+c, up, enter)
  at a configurable interval (default 1s). Common pattern for daemon restarts.
