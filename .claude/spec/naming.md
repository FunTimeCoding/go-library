# Naming Conventions

Enforced by the `goanalyze` naming analyzer (`pkg/lint/analyzer/naming/`). Banned segments cause
a lint error; this spec documents the resolution pattern for each.

## Banned segments and replacements

| Banned | Context | Use instead |
|--------|---------|-------------|
| `data` | local `[]byte` from `os.ReadFile`, `json.Marshal*` | `b` |
| `data` | local `os.FileInfo`/`fs.FileInfo` | `i` |
| `data` | struct field holding opaque `map[string]any` | `Payload` |
| `data` | struct field holding HTTP request body (`any`) | `Body` |
| `data` | struct field holding text/string content | `Content` |
| `data` | struct field holding Kubernetes Secret `.data` block | `Payload` |
| `data` | type name suffix (e.g. `FooData`) | strip suffix when no conflict; `FooPayload` when `Foo` is taken in the same package |
| `data` | local byte slice parameters (exported funcs) | domain-specific: `payload`, `archive`, etc. |
| `info` | local `os.FileInfo`/`fs.FileInfo` in walk/loop | `i` |
| `info` | type name suffix (e.g. `FooInfo`) | strip suffix or replace with descriptive noun (`Meta`) |
| `handler` | struct type implementing routes | `Router` (see openapi.md) |
| `handler` | function parameter | `serve` |
| `handler` | function name suffix | drop suffix: `logRequestHandler` → `logRequest` |
| `config` | type name | `Configuration` |
| `config` | function name | `LoadConfiguration`, etc. |
| `llm` | any identifier | `model` (e.g. `llmPrefix` → `modelPrefix`) |
| `tmp` | field prefix | drop prefix: `TmpFromAccountID` → `FromAccountID` |
| `doc` | local `*goquery.Document` | `d` |

## Type name conflicts

When stripping a banned suffix would produce a name already taken in the same package,
append `Payload` instead of stripping:

```go
// Comment type exists → can't strip to Comment
type CommentPayload struct { ... }

// No conflict → strip cleanly
type Exception struct { ... }  // was ExceptionData
```

## Metadata types

Avoid acronym-based type names. Prefer a plain English noun:

```go
type Meta struct { ... }  // was SDKInfo - holds SDK name/version/packages
```

`Meta` is appropriate when the struct holds metadata about a tool or runtime component
and no collision exists in the package.

## Domain terms

Some names contain a banned segment but refer to a real-world domain concept. Rename to the
most specific English term for that domain artifact:

| Original | Renamed | Rationale |
|----------|---------|-----------|
| `DataFile = "data.tar.gz"` | `ArchiveFile` | Alpine package archive segment |
| `PackageInfoFile = ".PKGINFO"` | `MetadataFile` | Alpine package metadata file |
| `CreateDataTar()` | `CreateArchive()` | creates the archive segment |
| `DataDirectory` (packager field) | `ArchiveDirectory` | directory whose contents become the archive |
| `dataDirectory` / `DataDirectory()` | `storageDirectory` / `StorageDirectory()` | app-level storage path |
| `DataDirectoryEnvironment` | `StorageDirectoryEnvironment` | env var pointing to storage dir |
| `DataHost` (CDN constant) | `AssetHost` | CDN endpoint for media files |

## HTTP handler conventions

See `openapi.md` for the full Router pattern. Summary:

- The struct implementing `ServerInterface` is named `Router`, never `Handler`
- Receiver on `*Router` methods: `r`
- `*http.Request` parameter: `q` (avoids collision with receiver `r`)
- Function parameter that accepts a handler func: `serve`, not `handler`

## Variable shorthands

These are established conventions across the codebase:

| Type | Name |
|------|------|
| `[]byte` from `ReadFile` / `Marshal*` | `b` |
| `os.FileInfo` / `fs.FileInfo` in walk/loop | `i` |
| `*goquery.Document` | `d` |
| `*gzip.Writer` | `z` |
| `*tar.Writer` | `t` |
| `error` (first in scope) | `e` |
| `error` (second in scope) | `f` |

See `conventions.md` for the full short variable name table.

## HTML custom attributes

`data-*` attributes are HTML's custom attribute mechanism. Use `Custom` as the replacement term:

```go
CustomPrefix       = "data-"   // was DataPrefix
KeepCustomAttribute bool        // was KeepDataAttribute
```
