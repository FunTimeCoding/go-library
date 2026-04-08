# Package Design

One struct with receivers per package. This is the core rule for package
scope — it drives when to extract sub-packages and how to organize types.

## The Rule

Each package has at most one struct that carries methods (receivers). That
struct is the package's identity. Its file is named after the struct in
snake_case. Constructor lives in `new.go`.

Data-only structs (no receivers) may coexist — one per file, named after
the struct. Group related data structs in a dedicated sub-package when
there are several (e.g. `response/` for response types, `option/` for
configuration). Data-only structs with external dependencies or distinct
domain identity may also warrant their own package, even without receivers,
to keep the parent package's import graph clean and the concept
self-contained.

## When to Extract

When a second struct with receivers appears in a package, extract it and
its helper functions into a sub-package. The sub-package is named after
the struct. The struct inside drops the package prefix:

```
pkg/tool/gosalt/analyzer/    → type Live struct
pkg/tool/goalertlogd/store/  → type Store struct
pkg/tool/goalertlogd/poller/ → type Poller struct
```

Not `analyzer.Analyzer` or `store.StoreStruct` — the package name
provides context.

## Interfaces

Domain-specific interfaces referencing the extracted struct's types are
defined at the consumer, not in the sub-package — unless the sub-package
owns the domain. The sub-package's mock lives in `mock_*/` next to the
real implementation.

## What Stays

Package-level functions without a receiver are fine alongside the single
struct. Constants live in `constant.go`. Each function gets its own file.

Tool entrypoint packages (`pkg/tool/go<name>/`) often have no struct with
receivers at all — just `Main()`, `run()`, and orchestration functions.
This is normal. The rule applies when structs with receivers do appear.
