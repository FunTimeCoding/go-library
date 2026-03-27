## Auto-fix for remaining analyzers

### forbidden_call — call site rewrite

- `exec.Command(args...)` → `run.Command(args...)`
- `exec.CommandContext(x, args...)` → `run.CommandContext(x, args...)`
- Single `TextEdit` replacing `exec.Command` selector with `run.Command`
- Import change: remove `"os/exec"`, add `"pkg/system/run"`
  (import edits are additional `TextEdit`s in the same `SuggestedFix`)

### string_concatenation — expression rewrite

- `a + b` → `fmt.Sprintf("%s%s", a, b)`
- `s += v` → `s = fmt.Sprintf("%s%s", s, v)`
- Need to render operand expressions as source text for the Sprintf args
- Multi-operand chains: `a + b + c` — only the outermost reports, so
  all operands are collected at once

### struct_literal — constructor call

- `&pkg.X{}` → `pkg.NewX()` — only when constructor exists and takes no args
- `new(pkg.X)` → `pkg.NewX()` — same constraint
- Must verify constructor existence (check if `NewX` is exported from package)
- When struct has field values (`&pkg.X{A: 1}`), skip — not auto-fixable
