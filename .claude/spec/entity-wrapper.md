# Entity Wrapper Spec

Pattern for wrapping entities from external systems with CLI rendering support.

## Package Structure

```
pkg/<service>/
├── client.go                 # Client struct with http.Client and base URL
├── new.go                    # Client factory
├── <entity>.go               # Client method to fetch single entity
├── <entities>.go             # Client method to fetch entity list
├── constant/
│   └── constant.go           # Service-level constants (e.g., Host)
├── <entity>/
│   ├── <entity>.go           # Pure struct definition
│   ├── constant.go           # Fallback display constants
│   ├── format.go             # Main Format(f *option.Format) method
│   ├── format_<field>.go     # Per-field formatters (private)
│   ├── new.go                # Factory from external type
│   ├── new_slice.go          # Slice factory
│   └── from_<source>.go      # Parsing factories for HTML sources (scraper only)
└── example/
    └── read.go               # CLI usage example
```

## Example: GitLab Jobs

### Entity Struct (`job/job.go`)

Pure data, no client references:

```go
package job

import (
    "github.com/funtimecoding/go-library/pkg/gitlab/project"
    "gitlab.com/gitlab-org/api/client-go/v2"
    "time"
)

type Job struct {
    MonitorIdentifier string
    Identifier        int64
    Name              string
    Status            string
    Stage             string
    Create            *time.Time
    Link              string

    Project *project.Project // nil unless enriched

    Trace string // empty unless enriched and job failed

    concern []string

    Raw *gitlab.Job
}
```

### Constants (`job/constant.go`)

Fallback values for missing data:

```go
package job

const (
    NoUser    = "no user"
    NoProject = "no project"

    FailConcern = "fail"
    Timeout     = "timeout"
)
```

### Main Formatter (`job/format.go`)

Composes field formatters into a status line:

```go
package job

import (
    "github.com/funtimecoding/go-library/pkg/console/status"
    "github.com/funtimecoding/go-library/pkg/console/status/option"
    "github.com/funtimecoding/go-library/pkg/console/status/tag"
)

func (j *Job) Format(f *option.Format) string {
    s := status.New(f)

    if f.HasTag(tag.Identifier) {
        s.Integer64(j.Identifier)
    }

    s.String(j.formatName(f))

    if f.HasTag(tag.Project) {
        s.String(j.formatProject())
    }

    s.String(
        j.formatUser(),
        j.formatDate(f),
        j.formatConcern(f),
    )
    s.DetailLink(j.Link, "GitLab", "")

    return s.RawList(j.Raw).Format()
}
```

### Field Formatter (`job/format_name.go`)

Each field has its own formatter with color support:

```go
package job

import (
    "github.com/funtimecoding/go-library/pkg/console"
    "github.com/funtimecoding/go-library/pkg/console/status/option"
)

func (j *Job) formatName(f *option.Format) string {
    if f.UseColor {
        return console.Cyan("%s", j.Name)
    }

    return j.Name
}
```

### Field Formatter with Fallback (`job/format_concern.go`)

Handle missing values gracefully:

```go
package job

import (
    "github.com/funtimecoding/go-library/pkg/console"
    "github.com/funtimecoding/go-library/pkg/console/status/option"
    "github.com/funtimecoding/go-library/pkg/strings/join"
)

func (j *Job) formatConcern(f *option.Format) string {
    if len(j.concern) == 0 {
        return ""
    }

    result := join.Comma(j.concern)

    if f.UseColor {
        result = console.Yellow("%s", result)
    }

    return result
}
```

### Factory (`job/new.go`)

Wraps the external library type into a pure entity:

```go
package job

import (
    "github.com/funtimecoding/go-library/pkg/monitor/item/constant"
    "gitlab.com/gitlab-org/api/client-go/v2"
)

func New(v *gitlab.Job) *Job {
    return &Job{
        MonitorIdentifier: constant.GoGitLab.Integer64Identifier(v.ID),
        Identifier:        v.ID,
        Name:              v.Name,
        Status:            v.Status,
        Stage:             v.Stage,
        Create:            v.CreatedAt,
        Link:              v.WebURL,
        Raw:               v,
    }
}
```

### Slice Factory (`job/new_slice.go`)

```go
package job

import "gitlab.com/gitlab-org/api/client-go/v2"

func NewSlice(v []*gitlab.Job) []*Job {
    var result []*Job

    for _, e := range v {
        result = append(result, New(e))
    }

    return result
}
```

### Client Method (`gitlab/project_jobs.go`)

Client handles the library call, delegates parsing to the entity package:

```go
package gitlab

import (
    "github.com/funtimecoding/go-library/pkg/gitlab/job"
    "github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func (c *Client) ProjectJobs(p *project.Project) []*job.Job {
    result, r, e := c.client.Jobs.ListProjectJobs(p.Identifier, nil)
    panicOnError(r, e)

    return c.enrichProjectJobs(job.NewSlice(result), p)
}
```

## The `basic/` Subpackage

**First, look for an importable client library.** Many APIs have well-maintained Go clients (GitHub, GitLab, Prometheus, etc.) - prefer those over building from scratch. `basic/` exists for two situations: (1) no good library exists and a hand-rolled HTTP client is the permanent solution, or (2) as temporary scaffolding while exploring an unfamiliar API before deciding whether to keep it or replace it with a library later.

When a client talks to a JSON API (vs. scraping HTML) and no suitable library exists, the raw HTTP layer lives in a `basic/` subpackage. The top-level client embeds it and returns typed entities.

```
pkg/<name>/
├── client.go          # Client struct: basic *basic.Client
├── new.go             # New(host, token string) *Client
├── new_environment.go # NewEnvironment() reads env vars, calls New()
├── <entities>.go      # Methods returning []*entity.Entity
└── basic/
    ├── client.go      # Client struct: host, port, auth fields
    ├── new.go         # New(...) *Client
    └── get.go         # Get(path string) string
```

`basic/get.go` from `pkg/jenkins/basic` - builds URL with locator, sets auth, sends:

```go
func (c *Client) Get(path string) string {
    r := web.NewGet(locator.New(c.host).Port(c.port).Path(path).String())
    r.SetBasicAuth(c.user, c.password)
    r.Header.Add(constant.ContentType, constant.Object)
    r.Header.Add(constant.Accept, constant.Object)
    response := web.Send(web.Client(), r)

    return web.ReadString(response)
}
```

For JSON APIs with typed responses, `Get` accepts an `out any` param and decodes into it:

```go
func (c *Client) Get(path string, params map[string]string, out any) {
    l := c.base.Copy().Path(path)
    for k, v := range params {
        l.Set(k, v)
    }
    r := web.NewGet(l.String())
    r.SetBasicAuth(c.user, c.token)
    r.Header.Set(webconstant.UserAgent, constant.UserAgent)
    response := web.Send(web.Client(), r)
    defer errors.PanicClose(response.Body)
    errors.PanicOnError(json.NewDecoder(response.Body).Decode(out))
}
```

The top-level client calls `basic.Get` and wraps results into entities:

```go
func (c *Client) Posts(tag string, limit int) []*post.Post {
    var out []post.PostData
    c.basic.Get("/posts", map[string]string{"tags": tag}, &out)
    return post.NewSlice(out)
}
```

`basic/` is omitted when the client is simple enough to do HTTP directly (e.g., scraper clients using goquery, single-endpoint tools).

## Key Principles

1. **Structs are pure data** - no client references, no HTTP logic
2. **Entities own their presentation** - `Format(f)` method on the entity
3. **Client does HTTP only** - fetches response, delegates parsing
4. **Wrap external types** - `New(v *library.Type)` maps fields, stores `Raw` for passthrough
5. **Color is optional** - controlled by `option.Format.UseColor`
6. **Missing data has fallbacks** - constants like `NoUser`, `NoProject`
7. **Prefer importable libraries** - `basic/` is a fallback, not a default
