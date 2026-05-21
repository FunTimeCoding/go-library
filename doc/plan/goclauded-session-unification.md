# goclauded — architecture and design

## Session Model

Alias and description are columns on the `session` table. One-time
migration from the legacy `alias` table (copies data, drops table).
`store/alias` package removed.

Unified write path: `store.EditSession` with
`service/argument.EditSession` (`*string` pointer fields — nil
means don't touch, empty string means clear).

## Store Layer

Pure persistence. SQLite via GORM. No knowledge of claude/JSONL.
`UpdateFields(identifier, map)` for generic column updates.
`HasChanges` is a pure read — no side effects. `ConsumeRoster`,
`ConsumeTimeout`, `ConsumeReannounce` are separate consume-once
methods.

## Service Layer

The service owns session lifecycle orchestration, all external
dependencies, and data assembly across sources.

**Dependencies (all private, accessor methods):**
- `store` — `Store()` accessor
- `claude` — `face.ClaudeSource` interface (9 methods), `Claude()`
  accessor. `mock_claude/` for tests.
- `memory` — gomemoryd client for impression forwarding
- `indexer` — goqueryd indexer for summary push

No nil guards — mocks guarantee non-nil in tests.

**Lifecycle methods (no error):**
Announce, Update, Complete, Release, Moment, ClearBindings, Send,
EditSession, DeleteSession (JSONL + sweep + store), ListSessions,
AllSessions, EnrichSession, Register.

**Lifecycle methods (returns error):**
Summarize, EditEvent — indexer push can fail. Must variants for
callers that PanicOnError (REST/web handlers).

**Data assembly methods:**
- `EnrichedSessions(limit, offset)` → `[]*enriched_session.Session`
  Joins JSONL sessions with store data (alias, description, turn
  count, active flag). Used by REST GetSessions, MCP list_sessions,
  web sessions page, conversations sidebar, dashboard roster.
- `SessionDetail(query)` → `*session_detail.Detail`
  Resolves identifier, joins JSONL + store + completions + summary.
  Used by REST GetSessionDetail, web session detail page.
- `Check(sessionIdentifier)` → `*check_result.Result`
  Assembles hook check response: roster, messages, completions,
  memory activity, timeouts, reannounce.

**Domain types (subpackages of service/):**
- `enriched_session.Session` — JSONL + store merged view with
  active flag
- `session_detail.Detail` — full session with completions + summary
- `check_result.Result` — hook check response
- `memory_activity.Activity` — memory change entry

## ClaudeSource Interface

`face/claude_source.go`:
Sessions, Resolve, Messages, FirstUserMessage, Peek, Delete,
SessionsByTool, ToolCalls, ToolContext.

`mock_claude/` — one file per method, zero-value returns. Used by
all test testers. Pattern follows gofirefoxd, gohabiticad,
goitermd, gosublimed face/ conventions.

## MCP Tools

- `edit_session` — alias, description, topic, files (all optional,
  at least one required). Replaces `rename` + `describe`.
- `edit_event` — edit event body with cascade to summary/completion.
  Replaces `edit`.
- `list_sessions` — browse all sessions with limit/offset. Uses
  `EnrichedSessions` — shows slug/alias, lines, turns, active flag,
  description.
- `roster` — active sessions with descriptions.
- Existing: announce, update, complete, release, summarize, moment,
  send, listen, status, history, history_count.

`captureFail` on MCP server for error-to-tool-result translation
on Summarize and EditEvent.

## REST Surface

- `POST /api/sessions/edit` with `EditSessionRequest` (alias,
  description, topic, files). Replaces `/api/alias` endpoints.
- `GET /api/sessions` — uses `service.EnrichedSessions`. Returns
  descriptions, supports limit/offset.
- `GET /api/sessions/{identifier}/detail` — uses
  `service.SessionDetail`.
- `GET /api/check` — uses `service.Check`.
- `POST /api/register` — uses `service.Register`.

Server struct holds only `*service.Service` + logger + paths.
No direct claude or store references.

## CLI (goclaude)

- `edit --name --description` — replaces `rename` + `describe`.
- `list` with `--detail`, `--limit`, `--offset`.
- `command_context.Context` with `PersistentPreRun` initialization.
- `--host`/`--port` persistent flags, `CLAUDE_HOST`/`CLAUDE_PORT`
  env var fallback to `localhost:8080`.
- `GOCLAUDED_*` env vars renamed to `CLAUDE_*`.

## Web UI

**Dashboard** (`/`): Session cards showing pool name, alias, topic,
description, files, turns, lines, last seen. Cards link to session
detail. Lines from JSONL via `EnrichedSessions`. SSE live updates
via roster + activity sections.

**Sessions** (`/sessions`): Table with name, alias, description,
lines, turns. All data from `EnrichedSessions`. Active/all filter
via `Active` flag. Pagination (20/page). Rows link to detail.

**Session detail** (`/sessions/{identifier}`): Full metadata from
`service.SessionDetail` — alias, description, slug, identifier,
created, turns, last seen. Completions table. Summary text.
Edit/conversation/delete links.

**Session edit** (`/sessions/{identifier}/edit`): Alias +
description form, redirects to detail.

**Conversations** (`/conversations`): Sidebar entries from
`EnrichedSessions` — alias, description, timestamp. Edit form
loads in panel (alias + description). Sidebar refreshes after edit
via HX-Trigger. Infinite scroll with sentinel pagination.

**History** (`/history`): Timeline with inline editing of event
bodies (complete, summarize, moment). Cascade to summary/completion
records. Pagination.

**Styling:** No underlines on content links. Card border highlight
on hover. Path parameters use `{identifier}` consistently.

## Test Infrastructure

Composable hierarchy via embedding:
- `store_tester` — store + controllable clock + `Advance()`
- `service_tester` — embeds store_tester, adds service + indexer
- `base` — embeds service_tester, adds full HTTP stack (REST + MCP)

Single shared instances: one store, one clock, one indexer used
by service and model_context. `mock_claude` for the claude source.
No duplicate creation.

**Coverage:** 34 store tests, 25 service tests, MCP coordination
and model_context tests. Lifecycle, edit, check, register,
resolution, timeouts, messages, indexer push cascade, delete.

## Evolution — Context Engineering

### 1. Proactive context from past sessions

The hook injects roster + messages + memory activity, but not
relevant history. When a session announces a topic, goclauded
could search goqueryd for past session summaries and completions
matching that topic and inject pointers into the hook context.
The infrastructure exists: summaries push to goqueryd, the hook
fires every turn, the topic is known from announce. The missing
piece is a relevance query during the check flow.

### 2. Session continuity links

A "continue from" field on the session linking to a predecessor.
When a session continues prior work, the link is explicit — the
hook injects the predecessor's summary and completions. The
wrapup skill could ask which session to continue. Not a new
table, just a field on the session row.

### 3. Completion-driven handoffs

When a session completes a task, the completion could
automatically enrich the next session that announces work on a
related topic. Completions become addressable context — not just
event log entries, but handoff artifacts that flow into successor
sessions.

### 4. The hook as a context compiler

The check handler currently assembles flat, independent sections.
A context compiler would weigh and prioritize: highlight sessions
touching the same files, surface freshly-created "always" memories
immediately, rank roster entries by relevance to the current
session's topic. The hook becomes intelligent about what matters
to this specific session right now.

### 5. Session archetypes

An archetype field (investigation, buildout, review, maintenance)
that adjusts what the hook injects. Investigation sessions get
more history and search results. Buildout sessions get file-change
awareness and lint/test status. Set by the model or the user at
announce time.

**Priority:** Start with #1 — proactive context. Everything else
builds on it, and the infrastructure is already in place.
