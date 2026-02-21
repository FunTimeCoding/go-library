# Plan

## Migrate tools to Main(version, gitHash, buildDate) pattern

### Batch 2 — Check tools, mechanical migration

All follow the same shape: register pflag args, `argument.ParseBind()`, build option struct, call `Check()`. Create `pkg/tool/` package, move logic, thin out `cmd/`.

- goalert
- gobrew
- gofile
- gogenie
- goghjob
- goghpr
- gogitlab
- gogitstatus
- gojira
- gokevt
- gosentry
- gosilence
- gov11y
- goversion
- gowiki

### Batch 3 — Complex extraction

Substantial inline logic that needs actual code extraction into `pkg/tool/`.

- goam (~10 lines, simple but different shape)
- gocat (~133 lines, file concatenation)
- god8y (~13 lines, no argument parsing)
- godebian (~134 lines, debian packaging)
- gomaintlog (~87 lines, mattermost + ollama)

### Skip

Not worth migrating — stubs, dead code, or no argument parsing.

- goansible (TODO placeholder)
- gomat (kubernetes client, no args)
- gomonitord (mostly disabled)
- gopg (mostly disabled)
- gorenovate (kubernetes client, no args)
- gosalt (TODO placeholder)
- gotrivy (kubernetes client, no args)
