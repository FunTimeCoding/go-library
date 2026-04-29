# Firefox MCP

MCP server for interacting with the user's running Firefox browser.
Read tabs, read page content, navigate, manage tabs and tab groups.
Single window, tab groups matter for organizing.

## System map

```
Firefox ← WebExtension (js/firefox_bridge/)
    ↕ WebSocket client → localhost:6125
pkg/firefox/            - Go (websocket server + client facade)
    ↕
pkg/tool/gofirefoxmcp/  - MCP server (default :8080, --port override)
                          also hosts websocket on :6125 (--bridge-port override)
    ↕ streamable HTTP
Claude Code / CLI / web consumers
```

Extension loaded as temporary add-on via about:debugging.
Two lifecycle server entries: one for the MCP HTTP transport, one for
the websocket that the extension connects to.

## Non-obvious decisions

- Extension-as-client pattern: the WebExtension connects to the Go process
  (not the other way around). Extension reconnects every 3 seconds when
  the Go process isn't running. Tolerates restarts on either side.
- `connecting` guard in background.js prevents reconnect storms - onerror
  delegates to onclose (which always fires after onerror), only onclose
  schedules reconnect
- `InsecureSkipVerify: true` on websocket.Accept - required because the
  extension's origin is its internal UUID, not a web origin
- Request/response correlation via integer ID: Client.send() creates a
  buffered channel per request, readLoop dispatches responses by ID
- 10-second timeout on send() if no response received - large DOMs
  (e.g. Thanos/Prometheus dashboards) may exceed this on read_tab
- "extension not connected" returns error immediately (no blocking wait
  for connection - the extension reconnects on its own schedule)
- Two separate lifecycle servers in run.go: websocket bridge on :6125,
  MCP HTTP on the --port flag. Both share the same Client instance.
- `read_tab` uses Readability.js by default, falls back to innerText on failure.
  Pass raw=true to skip readability (better for dashboards/web apps).
  Readability.js injected into tab via executeScript before parsing.
  Downloaded via taskfile, gitignored, not vendored.
- `navigate_back` uses `history.back()` via executeScript - no direct
  browser.tabs API for going back
- coder/websocket chosen over gorilla for context-native API and wsjson
  helpers - fits the request/response pattern cleanly
- `request.Integer` and `request.Boolean` types in mark/request/ handle
  both native JSON types and string-encoded values from MCP protocol.
  Used across all three MCP servers (Sublime, iTerm, Firefox).
- Tab group API requires Firefox 139+ (browser.tabs.group/ungroup,
  browser.tabGroups.update). Set tab color uses named colors
  (grey, blue, red, yellow, green, pink, purple, cyan, orange).

## Deferred

- `read_html` - raw HTML of a page
- `execute_script` - run JS in a tab (escape hatch, security concern)
- `search_history` - search browsing history
- `manage_bookmarks` - read/create bookmarks
- `capture_screenshot` - screenshot a tab
- Network monitoring (request/response capture)
- Increase read_tab timeout or make it configurable for large DOMs
