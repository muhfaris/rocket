# rocket

is generator golang project based ddd.

## overview

```
> go run main.go --help
Usage:
  rocket [flags]

Flags:
  -e, --cache string      cache for new project (e.g redis) (default "none")
  -c, --config string     config for new project (e.g toml) (default "toml")
  -d, --database string   database for new project (e.g mysql,  postgresql) (default "postgresql")
  -h, --help              help for rocket
  -n, --new string        name for new project (default "rocker-sample")
  -q, --queue string      queue fro new project (e.g rabbitmq) (default "none")
```

## Openapi

Table of contents

1. [Grouping Route](#grouping-route)
2. [Parameters](#parameters)
3. [OperationId](#operationid)
4. [RequestBody](#requestbody)
5. [Tags](#tags)

### Grouping Route

You can grouping some apis into one group with `x-route-group` tag. The default
group name is `routeGroup` and the default route is `/`.

```yaml
paths:
  /register/partners/{partner_id}:
    get:
      operationId: GetDetailPartner
      x-route-group: partnerGroup::/api
      tags:
        - Register Partner
      summary: Get detail partner account
      ...
      ...
    patch:
      operationId: UpdatePartnerHandler
      x-route-group: partnerGroup::/api
      tags:
        - Register Partner
      summary: Update partner account
      ...
      ...
```

Generated code from openapi spec are:

```go
	routeGroup := r.Group("/")
	routeGroup.Post("/register/partners", handlersv1.GetPartners())

	partnerGroup := r.Group("/api")
	partnerGroup.Get("/register/partners/:partner_id", handlersv1.GetDetailPartner())
	partnerGroup.Patch("/register/partners/:partner_id", handlersv1.UpdatePartnerHandler())

```

### Parameters

For each parameters of route need to define `x-parameters-name` tag. This tag
will be used for generate code as struct name, you can't ignore this tag.

Example openapi spec:

```yaml
/register/partners/{partner_id}:
  get:
    operationId: GetDetailPartner
    x-parameters-name: DetailPartner
    parameters:
      - name: partner_id
        in: path
        schema:
          type: string
        required: true
        example: "{{partner_id}}"
```

The generated code will be:

```go
type DetailPartner struct {
	PartnerID string `params:"partner_id"`
}

func GetDetailPartner() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var partner_id DetailPartner
		if err := c.ParamsParser(&partner_id); err != nil {
			return err
		}

		return response.Success(c, "I'm Alive!")
	}
}
```

### OperationId

OperationId should have unique name and following format `HandlerAPI, You
should use **title letter capital**. This will be used for generate code as Handler
name and will be called in route.

### RequestBody

You can define field-field in request body into the properties tag. So, you
should add `x-properties-name` tag for naming the properties and this will be
used as struct name and use format **title letter capital**.

Example openapi spec:

```yaml
requestBody:
  content:
    application/json:
      schema:
        type: object
        x-properties-name: UpdatePartner
        properties:
          fullname:
            type: string
          email:
            type: string
```

The generated code will be:

```go
type UpdatePartner struct {
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

func UpdatePartnerHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var bodyRequest UpdatePartner
		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		return response.Success(c, "I'm Alive!")
	}
}
```

If request body doesn't have properties, it will generated as `map[string]any`
without struct declaration.

```go
func UpdatePartnerHandler() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {

		var bodyRequest map[string]any
		if err := c.BodyParser(&bodyRequest); err != nil {
			return err
		}

		return response.Success(c, "I'm Alive!")
	}
}
```

### Tags

Every route should have one or more tags. The tag will be used for filename of domain
model.
