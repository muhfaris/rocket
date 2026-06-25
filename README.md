# rocket

**Rocket** generates production-ready Go projects from OpenAPI 3.0 specs — handlers, routes, domain models, services, repositories, and adapters, all wired in hexagonal architecture.

You provide an OpenAPI YAML file. Rocket produces a compilable Go project with:

- HTTP handlers (Fiber) mapped from your endpoints
- Service interfaces + implementations per domain
- Repository interfaces + DB adapter code (PostgreSQL, MySQL, SQLite, MongoDB)
- Cache adapter (Redis, in-memory)
- Route grouping, middleware, response envelopes
- Docker / Docker Compose files
- A `Makefile` with common targets

## Prerequisites

- **Go 1.22.5+** (see `go.mod`)
- OpenAPI 3.0 YAML spec (not Swagger 2.0)
- Optional: `goimports` (for import sorting — generator falls back to `gofmt`)

## Installation

```bash
go install github.com/muhfaris/rocket@latest
```

Or clone and build:

```bash
git clone https://github.com/muhfaris/rocket
cd rocket
go build -o rocket main.go
```

## Quickstart

Generate a complete project in one command using the included example:

```bash
go run main.go new -c rocket-books-api.yaml
cd books-api
go run main.go rest
```

This generates a Books API with:

- 6 endpoints across 2 route groups (`/api/v1` and `/api/v1` borrow group)
- Query parameters, path parameters, and request body handling
- Inline and structured response types
- A single service (`BookSvc`) with full handler/service/repository wiring

Open `http://localhost:7000/api/v1/books` to see it running.

The example spec is at [`spec/books-api.yaml`](spec/books-api.yaml) — use it as a reference for writing your own.

## Commands

```bash
rocket new [flags]
rocket add handler [flags]
rocket version
```

### `rocket new` — Create a new project

Generates a complete Go project from an OpenAPI spec.

```bash
rocket new \
  --package github.com/muhfaris/myproject \
  --project myproject \
  --openapi ./spec.yaml
```

| Flag | Type | Default | Description |
|---|---|---|---|
| `--package` | string | — | Go module path (e.g. `github.com/muhfaris/myproject`) |
| `--project` | string | — | Project directory name |
| `--openapi` | string | — | Path to OpenAPI YAML file |
| `--arch` | string | `hexagonal` | Architecture layout (only `hexagonal` is implemented) |
| `--cache` | string | — | Cache backend: `redis`, `inmemory` |
| `--database` | string | — | Database backend: `postgresql`, `mysql`, `sqlite`, `mongodb` |
| `--docker` | bool | `false` | Generate Dockerfile |
| `--config` | string | `./rocket.yaml` | Config file path |

### `rocket add handler` — Add handlers (WIP)

Intended to add a new handler + service/repository wiring to an existing generated project.

**Status: Work in progress.** The command accepts `--openapi` and `--operationid` flags but generation logic is not yet implemented. Tracked for a future release.

## Configuration: `rocket.yaml`

Instead of CLI flags, you can write a `rocket.yaml` file:

```yaml
openapi: ./spec.yaml
app:
  package: github.com/muhfaris/myproject
  project: myproject
  arch: hexagonal
  cache: redis
  database: postgresql
  docker: true
  ignore_data_response: true
```

Rocket searches for `rocket.yaml` in these locations (in order):

1. Path specified by `--config` / `-c`
2. Current directory (`./rocket.yaml`)
3. `./config/rocket.yaml`
4. `$HOME/.config/rocket.yaml`

CLI flags override values in the config file when both are provided.

### Config fields

| YAML key | Maps to flag | Description |
|---|---|---|
| `openapi` | `--openapi` | Path to OpenAPI spec |
| `app.package` | `--package` | Go module path |
| `app.project` | `--project` | Project directory name |
| `app.arch` | `--arch` | Architecture layout |
| `app.cache` | `--cache` | Cache backend |
| `app.database` | `--database` | Database backend |
| `app.docker` | `--docker` | Generate Dockerfile |
| `app.ignore_data_response` | — | When `true`, unwraps nested `data` objects in response schemas |


## OpenAPI Spec Requirements

Your OpenAPI spec must follow these conventions for the generator to produce correct code.

### Minimal spec

```yaml
openapi: 3.0.0
info:
  title: My API
  version: 1.0.0
servers:
  - url: http://localhost:8080
components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
    noauthAuth:
      type: http
      scheme: noauth
tags:
  - name: Books
paths:
  /books:
    get:
      operationId: ListBooks::BookSvc     # handler = ListBooks, service = BookSvc
      tags:
        - Books
      responses:
        "200":
          description: OK
          content:
            application/json:
              schema:
                type: object
                x-struct-response: ListBooksResponse
                properties:
                  items:
                    type: array
                    items:
                      type: object
                      properties:
                        id:
                          type: string
                        title:
                          type: string
                  total:
                    type: integer
```

### Required per operation

| Field | Requirement | Example |
|---|---|---|
| `operationId` | Unique name. Format: `HandlerName` or `HandlerName::ServiceName` | `ListBooks` or `ListBooks::BookSvc` |
| `tags` | At least one tag. Used as the domain filename (snake_case) | `- Books` -> `books.go` |

### The `::` convention

```
operationId: HandlerName::ServiceName
```

- **`HandlerName`** — becomes the Go function name (e.g. `ListBooks`, `CreateBook`).
- **`ServiceName`** — which service interface owns this handler. All endpoints with the same `ServiceName` share one interface.
- **Omitting `::ServiceName`** — defaults to `AppSvc`.

```yaml
# Produces: handler HealthCheck, service AppSvc
operationId: HealthCheck

# Produces: handler ListBooks, service BookSvc
operationId: ListBooks::BookSvc
```

## Custom OpenAPI Extensions Reference

Rocket defines several OpenAPI vendor extensions (`x-*`) that control code generation. These are placed inside your OpenAPI spec at the operation or schema level.

### Extension summary

| Extension | Scope | Purpose |
|---|---|---|
| `x-route-group` | Operation | Groups endpoints under a named Fiber route group with a path prefix |
| `x-parameters-name` | Operation | Names the Go struct for path/query parameters |
| `x-properties-name` | Schema (request body) | Names the Go struct for a JSON request body |
| `x-struct-response` | Schema (response) | Names the Go struct for a response schema |

---

### `x-route-group` — Route grouping

Groups endpoints into a named Fiber route group with a path prefix.

**Format:** `<groupName>::<pathPrefix>`

**Default:** `routeGroup` at `/`

```yaml
paths:
  /books:
    get:
      operationId: ListBooks
      x-route-group: bookGroup::/api/v1
```

Generated code:

```go
bookGroup := r.Group("/api/v1")
bookGroup.Get("/books", handlers_v1.ListBooks())
```

The route path `/books` is appended to the group prefix `/api/v1`, producing the full path `/api/v1/books`.

---

### `x-parameters-name` — Path/query parameter struct

Names the Go struct that holds path or query parameters for an operation.

```yaml
/books/{bookId}:
  get:
    operationId: GetBook
    x-parameters-name: GetBookParams
    parameters:
      - name: bookId
        in: path
        schema:
          type: string
```

Generated code:

```go
type GetBookParams struct {
    BookID string `params:"bookId"`
}
```

Without `x-parameters-name`, the struct is auto-named `<HandlerName>Params` (for path) or `<HandlerName>Query` (for query).

**Note:** When an operation has both path and query parameters, only one struct is generated. Use `x-parameters-name` to control its name.

---

### `x-properties-name` — Request body struct

Names the struct generated from a JSON request body schema.

```yaml
requestBody:
  content:
    application/json:
      schema:
        type: object
        x-properties-name: CreateBookRequest
        properties:
          title:
            type: string
          author:
            type: string
```

Generated code:

```go
type CreateBookRequest struct {
    Title  string `json:"title"`
    Author string `json:"author"`
}
```

If the request body has no properties (empty schema), the generator produces `map[string]any` instead.

---

### `x-struct-response` — Response struct

Names the struct generated from a response schema. Required when using inline response schemas (not `$ref`).

```yaml
responses:
  "200":
    content:
      application/json:
        schema:
          type: object
          x-struct-response: CreateBookResponse
          properties:
            id:
              type: string
```

Generated code:

```go
type CreateBookResponse struct {
    ID string `json:"id"`
}
```

**Required** for inline response schemas. Without it, the generator returns an error:

```
response should has x-struct-response as struct name
```

For `$ref` responses, the struct name is taken from the schema name in `components/schemas/`.

**Nested responses:** When an array item has its own `x-struct-response`, a separate struct is generated:

```yaml
responses:
  "200":
    schema:
      type: object
      x-struct-response: ListBooksResponse
      properties:
        items:
          type: array
          items:
            type: object
            x-struct-response: ListBookItem
            properties:
              id:
                type: string
              title:
                type: string
```

This generates two structs: `ListBooksResponse` (with `Items []ListBookItem`) and `ListBookItem`.


### Response types: inline vs. `$ref`

**Inline** — define properties directly in the response. Requires `x-struct-response`.

```yaml
schema:
  type: object
  x-struct-response: GetBookResponse
  properties:
    id:
      type: string
    title:
      type: string
```

**Components reference** — reference a schema defined in `components/schemas/`. The struct name uses the schema name.

```yaml
schema:
  $ref: "#/components/schemas/Book"
```

With a `components/schemas/Book` definition:

```yaml
components:
  schemas:
    Book:
      type: object
      properties:
        id:
          type: string
        title:
          type: string
```

Generated in `internal/core/domain/books.go`:

```go
type Book struct {
    ID    string `json:"id"`
    Title string `json:"title"`
}
```

### Method behavior

| HTTP method | Generates |
|---|---|
| `GET` | Query/param parser, **no request body** |
| `POST` | Body parser + path param parser |
| `PATCH` | Same as POST (body parser + path param parser) |
| `PUT` | Same as POST |
| `DELETE` | Same as POST |

### Security schemes

Define schemes in `components/securitySchemes`, then reference them per-endpoint:

```yaml
components:
  securitySchemes:
    bearerAuth:
      type: http
      scheme: bearer
paths:
  /books:
    post:
      security:
        - bearerAuth: []
```

Schemes are passed through to generated Swaggo annotations but do not generate auth middleware. Two schemes are pre-configured in examples: `bearerAuth` and `noauthAuth`.

### Data type mapping

| OpenAPI type | Go type |
|---|---|
| `string` | `string` |
| `integer` | `int` |
| `number` | `float64` |
| `boolean` | `bool` |
| `null` / `nullable` | `any` |

### Full example

See [`spec/books-api.yaml`](spec/books-api.yaml) for a complete working spec demonstrating all features: route groups, path/query parameters, request bodies, inline responses, nested array responses, and multiple HTTP methods.

## Generated Project Structure

```
<project>/
├── main.go                        # Entry point, calls cmd.Execute()
├── cmd/
│   ├── root.go                    # CLI entry point (cobra)
│   └── rest.go                    # HTTP server bootstrap
├── config/
│   ├── config.go                  # Config loader
│   └── config.yaml                # Default config values
├── internal/
│   ├── adapter/
│   │   ├── inbound/
│   │   │   └── rest/
│   │   │       └── router/
│   │   │           ├── router.go          # Route registration
│   │   │           ├── group/
│   │   │           │   └── v1.go          # Route groups
│   │   │           └── v1/
│   │   │               ├── handler/
│   │   │               │   ├── handler.go          # Handler init
│   │   │               │   └── <handler_name>.go   # Per-endpoint handlers
│   │   │               ├── presenter/              # Response mapping (TODO stubs)
│   │   │               ├── middleware/
│   │   │               │   └── latency.go
│   │   │               └── response/
│   │   │                   └── response.go          # Response helpers
│   │   └── outbound/
│   │       ├── cache/redis/          # Redis adapter (if --cache redis)
│   │       ├── datastore/psql/       # PostgreSQL adapter (if --database postgresql)
│   │       ├── datastore/mysql/      # MySQL adapter (if --database mysql)
│   │       ├── datastore/sqlite/     # SQLite adapter (if --database sqlite)
│   │       └── datastore/mongo/      # MongoDB adapter (if --database mongodb)
│   └── core/
│       ├── domain/                   # Domain models (from OpenAPI schemas)
│       ├── port/
│       │   ├── inbound/
│       │   │   ├── adapter/          # Inbound port interfaces
│       │   │   ├── registry/         # Service registry (wiring)
│       │   │   └── service/          # Service interfaces
│       │   └── outbound/
│       │       ├── datastore/        # Repository interfaces
│       │       └── repository/       # Cache, DB-specific repos
│       └── service/                  # Service implementations (stubs)
├── shared/
│   ├── apierror/                    # API error types
│   └── context/                     # Shared context helpers
├── spec/
│   └── openapi.yaml                 # Copy of input spec
├── Dockerfile                       # If --docker
├── docker-compose.yml               # If --database or --cache
├── Makefile
├── rocket.yaml                      # Saved generator config
├── README.md
├── go.mod
└── go.sum
```

### Key files explained

| File | Purpose |
|---|---|
| `cmd/root.go` | Cobra root command, registers the `rest` subcommand |
| `cmd/rest.go` | Starts the Fiber HTTP server on the configured port |
| `internal/adapter/inbound/rest/router/router.go` | Registers all route groups with the Fiber app |
| `internal/adapter/inbound/rest/router/group/v1.go` | Defines route groups (e.g. `/api/v1`) |
| `internal/adapter/inbound/rest/router/v1/handler/<name>.go` | HTTP handler per endpoint — parses request, calls service |
| `internal/adapter/inbound/rest/router/v1/presenter/<name>.go` | Maps domain models to response DTOs (auto-generated TODO stubs) |
| `internal/core/port/inbound/service/<name>.go` | Service interface that handlers depend on |
| `internal/core/service/<name>.go` | Service implementation (auto-generated stubs, you fill in logic) |
| `internal/core/port/inbound/registry/registry.go` | Wires services and repositories together |
| `internal/core/domain/<tag>.go` | Domain models grouped by OpenAPI tag |
| `cmd/bootstrap/app_repository.go` | Dependency injection — repository construction |
| `cmd/bootstrap/app_service.go` | Dependency injection — service construction |

## Architecture

The generated project follows **hexagonal architecture** (ports & adapters):

- **Domain layer** (`internal/core/domain/`) — business entities, no external dependencies
- **Service layer** (`internal/core/service/`) — business logic (auto-generated stubs)
- **Port interfaces** (`internal/core/port/`) — contracts for inbound (driving) and outbound (driven) adapters
- **Inbound adapters** (`internal/adapter/inbound/`) — HTTP handlers, route registration
- **Outbound adapters** (`internal/adapter/outbound/`) — database, cache implementations

### Request flow

```
HTTP Request -> Router -> Handler -> Port Service -> Service Impl -> Port Repository -> DB Adapter
                                        ^
                                  Presenter (response mapping)
```

## Backend Support

### Cache

| Flag value | Backend | Package |
|---|---|---|
| `redis` | Redis | `github.com/redis/go-redis` |
| `inmemory` | In-memory (no-op adapter) | stdlib |

### Database

| Flag value | Backend | Driver |
|---|---|---|
| `postgresql` | PostgreSQL | `github.com/jackc/pgx` |
| `mysql` | MySQL | `github.com/go-sql-driver/mysql` |
| `sqlite` | SQLite | `modernc.org/sqlite` (no CGo) |
| `mongodb` | MongoDB | `go.mongodb.org/mongo-driver` |

Each backend generates:

- An adapter (connection + command wrappers)
- A repository interface in `internal/core/port/outbound/repository/`
- A query repository implementation in the adapter's `repository/` subdirectory

**Note on values:** Use `postgresql` (not `postgres`) and `inmemory` (not `memory`) — these map to the internal constant values that match adapter generation.

## Presenter `Out()` method

Each handler has a corresponding presenter file with an `In()` and `Out()` method:

- **`In(c *fiber.Ctx)`** — parses the HTTP request (params, query, body) into a domain model. Auto-generated.
- **`Out(c *fiber.Ctx, data domain.X)`** — transforms the domain model into the API response DTO. Generated as a **TODO stub**:

```go
func (req *ListBooks) Out(c *fiber.Ctx, data domain.ListBooks) any {
    // TODO: map domain.ListBooks to ListBooksResponse
    return data
}
```

The fallback `return data` lets the API respond immediately. Replace it with your actual mapping when you add business logic.

## Response format

All responses go through `response.Success()` which wraps them in a standard envelope:

```json
{
  "data": <payload>,
  "metadata": {
    "latency": "2.34ms",
    "request_id": "abc-123"
  }
}
```

Errors use `response.Fail()`:

```json
{
  "errors": [{"message": "not found", "code": 404}],
  "metadata": {
    "latency": "1.02ms",
    "request_id": "abc-123"
  }
}
```

## Post-Generation Workflow

After `rocket new` completes, here is what you typically do next:

### 1. Inspect the generated project

```bash
cd <project>
tree -L 4
```

The generator runs `go mod tidy`, `goimports`, and `gofmt` automatically. If `go mod tidy` fails (e.g. offline), run it manually:

```bash
go mod tidy
```

### 2. Start the server

```bash
go run main.go rest
```

The server starts on the port configured in `config/config.yaml` (default `:7000`).

### 3. Implement business logic

The generated service stubs are in `internal/core/service/`. Each method returns an error by default — replace with your logic:

```go
func (s *BookSvc) ListBooks(ctx context.Context, payload domain.ListBooks) (domain.ListBooks, error) {
    // Your business logic here
    return payload, nil
}
```

### 4. Implement repository queries

For database-backed projects, repository implementations are in `internal/adapter/outbound/datastore/<db>/repository/`. Each method is a TODO stub. Add your SQL/query logic:

```go
func (r *BookRepository) FindAll(ctx context.Context) ([]domain.Book, error) {
    // Your query logic here
    rows, err := r.db.Query(ctx, "SELECT * FROM books")
    // ...
}
```

### 5. Update presenter mappings

Presenter files in `internal/adapter/inbound/rest/router/v1/presenter/` have TODO stubs for the `Out()` method. Map domain models to response DTOs:

```go
func (req *ListBooks) Out(c *fiber.Ctx, data domain.ListBooks) any {
    var items []ListBookItem
    for _, item := range data.Items {
        items = append(items, ListBookItem{
            ID:     item.ID,
            Title:  item.Title,
            Author: item.Author,
        })
    }
    return ListBooksResponse{Items: items, Total: data.Total}
}
```

### 6. Regenerate after spec changes

To regenerate after modifying your OpenAPI spec, re-run `rocket new` with the same project name. The generator detects existing directories and skips them — but note that it **overwrites** handler, presenter, domain, and service files.

**Best practice:** Commit your generated code before modifying specs, so you can diff changes.

## Generated Artifacts

### Docker (`--docker`)

Generates a multi-stage `Dockerfile` for building a small production image with distroless or scratch base.

### Docker Compose

When `--database` or `--cache` is set, a `docker-compose.yml` is generated with the required services (PostgreSQL, MySQL, Redis, etc.).

### Makefile

| Target | Description |
|---|---|
| `run` | Start the service |
| `build` | Build the binary |
| `test` | Run tests |
| `lint` | Run linter |

### README

A project-specific README is generated with the project name, setup instructions, and API documentation placeholders.

## Troubleshooting

| Error | Likely cause | Fix |
|---|---|---|
| `response should has x-struct-response as struct name` | Inline response schema without `x-struct-response` | Add `x-struct-response: StructName` to the response schema, or use `$ref` to a component |
| `X redeclared in this block` | Two endpoints reference the same `$ref` schema | Use inline `x-struct-response` instead of shared `$ref`, or use one service per schema |
| `undefined: portservice.AppSvc` | Registry references a service interface that wasn't generated | Ensure all `operationId` values use consistent `::ServiceName` |
| `no required module provides package` | Freshly generated project missing external deps | Run `go mod tidy` manually |
| `operation must have tags at path` | Endpoint is missing the `tags` field | Add at least one tag to each operation |
| PostgreSQL adapter not generated | Used `--database postgres` instead of `postgresql` | Use `--database postgresql` (matches internal constant) |
| Cache adapter not generated | Used `--cache memory` instead of `inmemory` | Use `--cache inmemory` or `--cache redis` |
| `project X already exists` | Project directory already exists | Delete or rename the directory, or use a different project name |
| `import path must include a domain` | `--package` doesn't include a domain | Use format like `github.com/user/project` |
| `go mod tidy failed` | Offline or network issue | Run `go mod tidy` manually in the project directory |

## Known limitations

- **Shared `$ref` schemas across endpoints**: referencing the same component schema from multiple endpoints causes duplicate type declarations. Use inline `x-struct-response` as a workaround.
- **Multiple `::ServiceName` values**: each unique service name generates a separate interface file. The registry and service implementations reference them correctly.
- **`add handler` is a stub**: the `rocket add handler` command accepts flags (`--openapi`, `--operationid`) but generation logic is not yet implemented. The initial `rocket new` command generates all endpoints from the spec; incremental additions are planned for a future release.
- **Only `hexagonal` architecture is implemented**: the `--arch` flag accepts `hexagonal` (and mentions `cleancode` in help text), but only hexagonal templates exist.
- **The generated `go.mod` uses `go mod tidy`** to resolve dependencies. If you're offline, run `go mod download` or vendor the dependencies first.
- **Regeneration overwrites files**: running `rocket new` on an existing project directory overwrites handlers, presenters, domain models, and services. Custom changes to generated files will be lost. Commit before regenerating.
