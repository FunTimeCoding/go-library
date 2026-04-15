# Deterministic variable naming

goanalyze analyzer that assigns canonical single-letter names to local variables
based on their type. Given a type and a scope, always picks the same letter.

## Scope

Local variables and struct receivers only. Function parameters are excluded
(callers read the signature — descriptive names preferred there). Struct fields
are excluded (exported API).

## Letter assignment precedence

Higher-precedence types claim their letter first. When a letter is taken,
the type falls through to its cycle sequence, then to `a-z`.

| Precedence | Type | Letters (in order) | Notes |
|------------|------|-------------------|-------|
| 1 | `error` | `e`, `f`, `g`, `h` | deepest convention |
| 2 | `*testing.T` | `t` | reserved in test functions |
| 3 | `*testing.B` | `b` | reserved in benchmarks |
| 4 | `io.Reader` | `r` | |
| 5 | `io.Writer` | `w` | |
| 6 | `*os.File` | `f` | after errors claim their letters |
| 7 | `*gzip.Writer` | `z` | existing convention |
| 8 | `*tar.Writer` | `t` | existing convention (non-test) |
| 9 | `string` | `s`, then `a-z` | |
| 10 | `int` / `int64` etc | `i`, then `a-z` | |
| 11 | `float64` / `float32` | `n` (number), then `a-z` | avoids `f` collision with error/file |
| 12 | `bool` | `b`, then `a-z` | |
| 13 | `byte` / `[]byte` | `b`, then `a-z` | rare to coexist with bool |
| 14 | `map[K]V` | `m`, then `a-z` | |
| 15 | `chan T` | `c`, then `a-z` | |
| 16 | slice of struct | `v` (values) | |
| 17 | slice of primitive | same letter as element type | |
| 18 | named struct | first letter of last word of type name, cycle through last word | |
| 19 | interface | first letter of interface name, cycle through name | |

## Conflict examples

```
func Example() {
    e := fmt.Errorf("first")     // error → e
    f := fmt.Errorf("second")    // error → f (e taken)
    file := os.Open("x")         // *os.File → f taken by error, try i,l,e → gets l? 
                                 // or cycle "file" → f,i,l,e → i? l?
    s := "hello"                 // string → s
    i := 42                      // int → i
    n := 3.14                    // float → n
    m := map[string]int{}        // map → m
}
```

The `*os.File` case when `f` is taken by a second error: cycle through "file" → `i`, `l`.
But `i` might be taken by an int. This is where the precedence ordering matters — errors
and the type-specific letters are assigned first in precedence order, then conflicts
cascade through the cycle/fallback.

## Algorithm

For a function scope:

1. Collect all local variables with their types
2. Sort by precedence (errors first, then testing, then reader/writer, etc.)
3. For each variable, try its letter sequence:
   a. Suggested letters for the type
   b. Letters of the long-form word (cycle)
   c. `a-z` exhaustive
4. First non-colliding letter wins; that letter is now taken for the scope
5. If all 26 exhausted, report error (function has too many variables)

## Interaction with abbreviation analyzer

The existing `naming` analyzer flags abbreviations (`cfg` → `configuration`). This
analyzer assigns letters by type. They compose:

- `naming` expands abbreviations in compound names (`cfgPath` → `configurationPath`)
- `variable_naming` assigns single letters to standalone local variables

A variable named `cfg` of type `*Configuration` would be flagged by `naming` for the
abbreviation AND by `variable_naming` for not using the canonical letter. The fix from
`variable_naming` takes precedence for standalone variables (it would become `c` or
whatever the struct rule assigns).

## Open questions

- Receivers already follow the struct rule across the codebase (`func (s *Store)`,
  `func (c *Client)`, `func (r *Reader)`). Validate only, don't reassign.
- Variables that are already single-letter and correct: skip. Only flag when the current
  name doesn't match what the oracle would assign.
- The `*os.File` / `f` / error `f` three-way: does File get its own precedence slot
  (current table says yes, at precedence 6), or does it follow the struct rule and get
  `f` from "File"? If it gets precedence 6, a function with one error and one file gets
  `e` and `f` naturally. A function with two errors and one file gets `e`, `f`, `g` for
  errors and the file needs to fall through — to what? Cycle "file": `i`, `l`, `e` → `l`.

## Parked

- Flag multi-character descriptive names (`result`, `letter`) for deterministic shortening
- Handle `ok`/`okay` booleans from type assertions and map lookups
- Fix path: scope-aware rename via `applyEdits` pipeline (detection-only for now)

## Not in scope

- Function parameters (descriptive names per conventions spec)
- Struct fields (exported API naming)
- Package-level variables
- Return value naming
