# golint / goanalyze feature plan

Feature requests for the next implementation pass, with design decisions recorded.

---

## In scope

### 1. Naming analyzer — extended bans + message format

**Location:** `pkg/lint/analyzer/naming/naming.go` (and testdata)

Replace the single `map[string]string` blacklist with two structures:

```go
// suggestions: banned segment → one or more acceptable alternatives
var suggestions = map[string][]string{
    "url":    {"l", "locator"},
    "mcp":    {"c", "model_context"},
    "dir":    {"d", "directory"},
    "dirs":   {"directories"},
    "tx":     {"t"},
    "ctx":    {"x"},
    "param":  {"parameter"},
    "msg":    {"m", "message"},
    "req":    {"r", "request"},   // multiple OK: one-letter preferred, full-word fallback
    "doc":    {"d", "document"},
    "config": {"c", "configuration"},
    "cfg":    {"c", "configuration"},
    "llm":    {"m", "model"},
    "tmp":    {"t"},
}

// noSuggestion: banned noise words with no prescribed replacement
var noSuggestion = map[string]bool{
    "handler": true,
    "data":    true,
    "info":    true,
}
```

Message formats:
- One suggestion: `use "r" instead of "ctx" in name`
- Multiple: `use "r" or "request" instead of "req" in name`
- No suggestion: `avoid "handler" in name, use a more specific term`

Rationale for no-suggestion words: shortening to e.g. `h` for `loginHandler` → `loginH` is silly.
For the suggestion words: one-letter is preferred, full word is acceptable fallback when
the one-letter is already in use.

Files to update: `naming.go`, `segments_test.go`, `analyzer_test.go`, testdata.

---

### 2. Ban `flag` stdlib package, require `pflag`

**Location:** `pkg/lint/forbidden_import.go` (new), `pkg/lint/constant/constant.go`

New `ForbiddenImport` checker. Scans for:
- Single-line `import "flag"`
- Inside multi-import blocks, a line that is `\t"flag"`

Reports with suggestion to use `github.com/spf13/pflag`. No auto-fix (import path and
API both need manual updating).

New constants:
```go
ForbiddenImportKey  = "forbidden_import"
ForbiddenImportText = `Use "github.com/spf13/pflag" instead of "flag"`
```

Register alongside `Import` in `check.go`.

---

### 3. One file per function

**Location:** `pkg/lint/function_count.go` (new), `pkg/lint/constant/constant.go`

New `FunctionCount` checker, mirrors `TypeCount`. Counts top-level `func` declarations
per file and flags if more than one is found.

Rules:
- Counts ALL `func` declarations: regular functions, methods, and `init()`
- `_test.go` files are exempt (multiple test/bench/example functions expected)
- Flags on the first function definition found when count exceeds 1

New constants:
```go
MultipleFunctionsKey  = "multiple_functions"
MultipleFunctionsText = "Multiple function definitions in one file"
```

Register in `check.go`.

---

### 4. Stray `const` outside of constant files

**Location:** `pkg/lint/stray_const.go` (new), `pkg/lint/constant/constant.go`

New `StrayConst` checker. Flags top-level `const` declarations unless the file is in
a constant context.

A file is considered a constant context if **either**:
- Its base name is `constant.go` or `constant_test.go`, OR
- Its `package` declaration is `package constant`

This allows both the common single-file pattern (`constant.go`) and the multi-file
pattern (a `constant/` package with many grouped files like `http.go`, `path.go`, etc.).

Only flags top-level const — `const` inside function bodies is ignored (tracked via
brace depth).

New constants:
```go
StrayConstKey  = "stray_const"
StrayConstText = "Top-level const outside of a constant file or constant package"
```

Register in `check.go`.

---

### 5. Blank lines between statements inside function bodies

**Location:** `pkg/lint/spacing.go`, `pkg/lint/constant/constant.go`

Extend the existing `Spacing` checker with a pending-blank pattern.

**Rule:** Inside a function body, a blank line is only valid before a control structure
(`if`/`for`/`switch`/`select`), before an exit statement (`return`/`break`/`continue`),
or after a control block's closing `}`. All other blank lines are removed.

Rationale: deterministic formatting over "logical groups". If logical grouping feels
necessary, the function should be broken down further.

**Implementation — pending-blank pattern:**

When a blank line is encountered inside a function body (`len(blockStack) >= 1`):
- Do NOT emit it immediately
- Set `pendingBlank = true` and continue

When the next non-blank line arrives with a pending blank:
- **Emit** the pending blank if any of:
  - `needBlankAfterClosingBrace` is true (blank is earned by prior control block)
  - Line is a control start (`if`/`for`/`switch`/`select`)
  - Line is an exit (`return`/`break`/`continue`)
- **Discard** (concern) if none of the above — i.e., it's a normal statement or
  a closing brace
- After deciding, clear `needBlankAfterClosingBrace` so the existing "add missing
  blank after control" check doesn't double-fire

Package-level blank lines (`len(blockStack) == 0`) pass through unchanged — blanks
between top-level declarations remain fine.

New constant (distinct from `ExtraneousBlankLine` which covers double-blanks):
```go
BlankInsideFunctionKey  = "blank_inside_function"
BlankInsideFunctionText = "Blank line between statements inside function body"
```

Test cases to add to `spacing_test.go`:
- Blank between two assignments → flagged, removed
- Blank at start of function body (before first statement) → flagged, removed
- Blank at end of function body (before closing `}`) → flagged, removed
- Blank before return that is NOT preceded by a control block → flagged, removed
- Blank between assignments inside nested block (e.g. inside `if`) → flagged
- Blank after control block before normal statement → NOT flagged (earned)
- Blank before `if` inside a function → NOT flagged (required blank before control)
- Closure/func literal: same rules apply inside closure body

---

## Out of scope (next pass)

| Item | Reason |
|------|--------|
| Direct `exec.Command` usage (flag, suggest `pkg/system/run`) | Easy, deferred to keep this pass focused |
| String `+` concatenation ban | Needs AST pass, not reliably detectable by text scanning |
| Flag which items were auto-fixed in output | Minor UX; output format not yet decided |
| Stub test filename `main_test.go` enforcement for `pkg/tool/<prog>` | Medium complexity; would change stub generation behavior |
