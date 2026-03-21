# Plan Files

Implementation plans live in `doc/plan/<feature>.md` in the relevant repo.
Plans are working documents, not permanent records - they shrink as work
completes and disappear entirely when done.

## When to Create

Create a plan file when a task spans multiple files or sessions and the
implementation choices are worth capturing before starting. Single-file
or obvious changes don't need one.

## How to Derive the Structure

Before writing a plan:
1. Read the relevant specs (`.claude/spec/`)
2. Explore existing code in the area - understand what already exists,
   what the file/package conventions are, and where new code will slot in
3. Identify all files to create or change, with their function signatures
   and responsibilities
4. Note any non-obvious decisions (e.g. re-reading a file vs threading
   state, detecting a format before parsing, panic vs error return)

The plan should reflect what was actually discovered, not generic steps.

## Format

Use `##` sections for distinct work areas. Under each, list:
- Files to create or change (with package-relative paths)
- Function signatures and what they do
- Any implementation constraints or decisions worth preserving

Keep it lean - just enough to execute without re-deriving, not a
full design document.

## Lifecycle

- **Delete completed sections** as the work is done, not after the whole
  plan is finished. The file shrinks as progress is made.
- **Delete the file** when all sections are gone.
- Plans are never committed as historical records - if you need to
  understand past decisions, read the code and git history.
