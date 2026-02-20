# Coding Conventions

## Error Handling

- **Prefer `PanicOnError` over returning errors** in test/internal code
- **Prefer `PanicClose` over checking close errors**
- Use `LogClose` only in loops/uncertain contexts (e.g., validation)
- Production APIs still return `error`, but helpers panic
- **Error variable naming:**
  - `e` for first error in scope
  - `f` for second error in same scope (avoid shadowing)
  - `g` for third error in same scope
  - When running out of e/f/g in deeply nested scopes: `createFail`, `saveFail` (prefer short words like "fail" over "error")
  - Eliminate variables entirely when possible: `errors.PanicOnError(someCall())`
  - Descriptive names like `createFail` or `printFail` are still "smells" - indicates code needs refactoring

## Testing Philosophy

- **Unix approach: silence on success, noise on failure**
- No progress logs (`t.Log("doing X...")`)
- No success logs (`t.Logf("✓ X succeeded")`)
- Only `t.Fatalf()` for must-succeed setup, `t.Errorf()` for validation failures
- Extract helper functions, make them panic instead of returning errors

## Code Style

- **Never-nesting ideal** - avoid deep indentation
- **No defensive programming in tests** - "should work or blow up"
- Remove explanatory comments if code is self-documenting
- Helper functions have minimal/no error returns in test code
- **Short variable names:**
  - `e`/`f`/`g` for errors in nested scopes
  - `h` for headers, `i` for FileInfo in loops
  - `z` for gzip writers, `t` for tar writers, `f` for files
  - `path` over `outputPath`, `name` over `fileName`
  - Single letters for obvious types
  - Avoid `-ing` and plural forms in names
  - Prefer shorter words: "fail" over "error", "path" over "filePath"
- **Function/method chaining** - pass return values directly: `p.WritePKGINFO(p.CreateDataTar())`
- **Blank identifier for byte counts** - ignore bytes read/written when not needed: `content, _ := io.ReadAll(tr)`, `_, f = outFile.Write(data)`
  - Never ignore error returns, only byte counts
- **Defer placement** - immediately after resource creation and error check:
  ```go
  file, e := os.Open(path)
  errors.PanicOnError(e)
  defer errors.PanicClose(file)
  ```
- **Empty line placement:**
  - Blank lines go BEFORE and AFTER entire control structure blocks (around the braces)
  - **Return statement blank lines:**
    - At function level: blank line before return (unless return is the only statement in entire function)
    - In control structure blocks (if/for/switch): blank line before return ONLY if multiple statements in that block
  - Only control structures and return statements break consecutive simple statement chains
  - No blank lines at function/method start or end (first/last statements sit flush with braces)
  - Consecutive simple statements (assignments, function calls, defer) should be compact (no blank lines)
  - Multiline statements (spanning multiple lines) count as one statement for blank line purposes
  - Aim for functions/methods with ~15 calls or fewer to maintain readability
  - Line length: break at 80 characters, prioritizing fewer total lines (break single-arg calls before multi-arg calls)
- **Multiline function calls** - when args don't fit one line, each arg gets its own line with closing paren separate
- **Prefer go-library wrappers** - use `system.Open()`, `system.Create()`, `system.FileStat()`, `system.Copy()`, `system.TarWriteHeader()`, `errors.PanicFlush()` instead of stdlib + error checking

## Structure

- **One-file-per-function** - extract helpers to dedicated files named after the function (e.g., `http_fail.go`, `add_file_to_tar.go`)

- Tests use `assert.*` helpers to reduce nesting
- Fail fast with stacktraces (Sentry integration)
- Helpers return only success values, panic on errors
- Prefer extracting logic to testable helpers over inline test code

## Deviations

**vs Claude defaults:** Claude writes defensive, verbose tests with every error checked and logged

**vs Idiomatic Go:** Idiomatic Go returns `(T, error)` everywhere; we prefer panic in non-public APIs
