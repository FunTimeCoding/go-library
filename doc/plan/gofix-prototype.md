## Goal

Prototype for cross-package identifier rename. Validates that we can:
1. Load all packages in the module with full type info and AST
2. Given an exported identifier, find every reference across all packages
3. Generate correct TextEdits for all references
4. Apply edits atomically (all files or none)

Once validated, this becomes the foundation for replacing `multichecker.Main()`
in goanalyze with a custom runner that does module-wide fixes.

## Architecture

Standalone tool at `cmd/gofix/main.go` + `pkg/tool/gofix/main.go`.
Follows entrypoint.md conventions (linker vars, sentry, Main()).

Takes two arguments:
- Package pattern (e.g., `./...`)
- No flags beyond standard ones initially

Runs the naming analyzer on all packages, then for EXPORTED identifiers
with violations, finds cross-package references and applies renames.

## Key pieces

### pkg/tool/gofix/main.go — Main()

- Sentry setup (per entrypoint.md)
- Call `load()` then `fix()`

### pkg/tool/gofix/load.go — load(patterns) []*packages.Package

- `go/packages.Load` with `packages.LoadSyntax` mode
  (gives us Files, TypesInfo, Syntax for all matched packages)
- Shared `token.FileSet` across all packages (default behavior)
- Error handling: print package errors, exit on failure

### pkg/tool/gofix/find_all_references.go — findAllReferences(packages, object) []reference

- For a `types.Object`, iterate ALL loaded packages
- Check each package's `TypesInfo.Defs` and `TypesInfo.Uses`
- Return list of `(file, ident)` pairs across all packages
- This is the cross-package version of naming's `findReferences`

### pkg/tool/gofix/fix.go — fix(packages)

- Run the naming analyzer's `segments()` + `suggestions` logic against
  all exported identifiers in all packages (reuse naming package's maps)
- For each violation: call `findAllReferences` to get module-wide references
- Compute replacement via `replaceSegment`
- Collect all edits grouped by file
- Apply edits to all files atomically (read all, compute diffs, write all)
- Print summary of what was renamed and in how many files

### pkg/tool/gofix/apply_edits.go — applyEdits(fileSet, edits)

- Group TextEdits by file
- Sort edits within each file by position (descending, so earlier edits
  don't shift later positions)
- Read each file, apply byte-range replacements, write back
- Optionally print unified diff instead of writing (--diff flag)

## What this prototype does NOT do

- Does not replace multichecker — goanalyze continues working as-is
- Does not run all analyzers — only naming violations on exported identifiers
- No conflict detection yet (collision with existing names)
- No import rewriting (that's forbidden_call territory)

## Unknowns to validate

- Does `packages.LoadSyntax` on `./...` actually give us TypesInfo for
  every package, including test packages?
- Memory footprint for loading ~80 packages with full syntax
- Are `token.Pos` values from different packages in the same FileSet
  comparable and usable for byte-offset calculation?
- Do packages loaded via `go/packages` include vendored or generated code
  we'd want to skip?
