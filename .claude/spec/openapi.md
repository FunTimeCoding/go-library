# OpenAPI Codegen

Pattern for typed HTTP APIs using `oapi-codegen`. Use when a service exposes an HTTP API that needs a generated client (or when type safety on the server side is worth the codegen overhead).

## Package Structure

```
pkg/<service>/
├── server/
│   ├── openapi.yaml   # Hand-written spec
│   ├── config.yaml    # oapi-codegen server config
│   └── generated.go   # Generated: ServerInterface, HandlerFromMux, types
├── client/
│   ├── config.yaml    # oapi-codegen client config
│   └── generated.go   # Generated: Client, typed request/response methods
└── route/             # Sibling to server/ - implements ServerInterface
    ├── router.go      # Router struct (holds dependencies)
    ├── new.go         # New(dep *dep.Dep) *Router
    └── <operation>.go # One file per endpoint (post_deploy.go, get_status.go)
```

`route/` is a sibling to `server/`, not a child. It implements the generated interface but is not part of the generated package.

## Configs

`server/config.yaml`:
```yaml
package: server
generate:
  std-http-server: true
  models: true
  embedded-spec: true
output: generated.go
```

`client/config.yaml`:
```yaml
package: client
generate:
  client: true
  models: true
output: generated.go
```

## Generation

Run via taskfile:
```yaml
generate:
  cmds:
  - cd pkg/<service>/server && oapi-codegen --config config.yaml openapi.yaml
  - cd pkg/<service>/client && oapi-codegen --config config.yaml ../server/openapi.yaml
```

Add generated files to golint skip list:
```yaml
lint:
  cmds: ['golint --fix --skip tmp,pkg/<service>/server/generated.go,pkg/<service>/client/generated.go']
```

## Wiring into the Server

In `run.go`, the oapi-codegen server package is always aliased `generated`. Call `generated.HandlerFromMux` inside `lifecycle.WithServer`. It registers routes on the mux as a side effect - the return value can be ignored:

```go
import (
    generated "github.com/funtimecoding/go-library/pkg/tool/go<tool>d/server"
)

lifecycle.WithServer(
    web.Listen,
    func(m *http.ServeMux) {
        generated.HandlerFromMux(route.New(dep), m)
    },
)
```

The `generated` alias is the standard across all services - it makes clear the package is codegen output, and frees up the unaliased `server` name for any domain service package.

If the service has existing manual routes alongside generated ones, both can coexist on the same mux.

## Implementing the Router

This pattern is for OpenAPI-generated routes. For manually registered routes without a spec, see the `route/` section in `service-tool.md`.

`route/router.go` - plain struct, no embedding:

```go
type Router struct {
    deploy *deploy.Observer
}
```

`route/<operation>.go` - implements the generated `ServerInterface` method. Use the `generated` alias for the server package:

```go
func (r *Router) PostDeploy(w http.ResponseWriter, q *http.Request) {
    var body generated.PostDeployJSONRequestBody
    errors.PanicOnError(json.NewDecoder(q.Body).Decode(&body))
    result := r.deploy.TriggerTargets(body.Targets)
    w.Header().Set(constant.ContentType, constant.Object)
    errors.PanicOnError(json.NewEncoder(w).Encode(generated.DeployResponse{Tag: result}))
}
```

## Using the Generated Client

Wrap the host with `locator.New(host).String()` - hosts are stored without scheme, the generated client needs a full URL:

```go
c, e := client.NewClient(locator.New(environment.Required(constant.HostEnvironment)).String())
errors.PanicOnError(e)
resp, e := c.PostDeploy(ctx, client.PostDeployJSONRequestBody{...})
```

## What Not To Do

- Don't name the server package `api` - use `server`
- Don't nest `route/` inside `server/` - keep it a sibling
- Don't manually register routes that the spec already defines - let `HandlerFromMux` do it
- Don't edit `generated.go` - regenerate from the spec instead
