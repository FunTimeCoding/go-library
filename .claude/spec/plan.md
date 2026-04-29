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
- **Delete the file** when all sections are gone and no deferred items
  remain. Code and git history are the primary record.
- **Don't delete plans with remaining items.** If a plan has deferred
  work, migration steps, or items waiting on deployment, it stays until
  those are resolved. Discuss with the user what to do with deferred
  items - whether they stay in the plan, move to a design doc as
  remaining work, or get extracted to their own seed.

## Completion

When implementation is done but before deleting the plan:

1. **Update design docs** in `doc/design/` to reflect what was built.
   The design doc is the final destination for system knowledge; plans
   are working documents that disappear.
2. **Review the seed** (if one exists) for what graduated vs what's
   still live. Mark completed parts, frame the next step.
3. **Give deferred items homes.** Tangential discoveries (bugs found,
   infrastructure issues, ideas for adjacent features) should be
   extracted to their own seeds, not left in the plan or silently
   dropped.
4. **Discuss disposition with the user.** Where things live affects
   visibility and priority. Whether remaining work goes in a design
   doc (long-term, less visible) vs stays in the plan (active, more
   visible) is a collaborative decision, not a mechanical filing step.

## Preserving

When a plan evolved into a system description during iteration - multiple
external integrations, ongoing work across sessions, package structure
worth orienting future sessions around - transform it instead of deleting:
strip task tracking and status markers, keep the system map, moving parts,
and remaining work. The result stays in `doc/plan/` as a living document.
Revisit and prune on subsequent sessions.
