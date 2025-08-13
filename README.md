# Go RBAC Demo Using Casbin

## Rust

I played with the idea of doing this in Rust and the output is disgusting.
```sh
npx @openapitools/openapi-generator-cli generate -i server/openapi/spec.yml -g rust-axum -o rust_server
```

## Highlights

- Go server code generated from OpenAPI spec
- Security defined in the OpenAPI template
- RBAC handled in a middleware before your app code ever runs

## Casbin

Theis demo uses the Casbin library for enforcing RBAC

### Roles

> :warning: <br/> This is only my experience so far. I do not speak with authority on how to structure or use this policy. The built in DB adapters are a no go. Would need to build you own adapter, which doesn't seem too hard.

There are 2 roles and 2 users:

- User (ID 123)
- Admin (ID 456)

The 123 user has been assigned to the User role, and 456 has been assigned to the Admin role. This is super basic, and largely just to illustrate the flow.

The assignment of roles is in the [rbac_policy.csv](./server/session/rbac_policy.csv). The syntax is pretty wookie.

### Policy

In order to associate a user with a policy, enter a line into the CSV.

| Type of declaration | Identifier | Group | Resource |
| ------------------- | ---------- | ----- | -------- |
| p                   | 123        | user  | data1    |

### Group

Once a user has been registered, they can be assigned to a group with the following line in the CSV

| Type of declaration | Group | Action          |
| ------------------- | ----- | --------------- |
| g                   | user  | GetRbacResource |

The action in this case is defined in the [OpenAPI spec](./server/openapi/paths/rbac.yml#L8). It is passes automatically to the middleware.

### Middleware

The [OpenAPI spec](./server/openapi/paths/rbac.yml#L8) defines that both `BearerAuth` and `Rbac` security schemes are required for this route's application code is to run.

Those schemes are defined in the [root spec](./server/openapi/schema.yml#L14). The names are arbitrary. 

A minor con is that OpenAPI only supports a few security types. I declared RBAC as a Bearer auth, which is not accurate. This is the only way I was able to get it working with the auto gen. It doesn't impact anything, because you define the logic of the security handler.

#### Bearer

I have not actually implemented any JWT logic for this demo. I simply check the result of the token. The 2 allowed values are `123` (User) and `456` (Admin).

The code for that is in the [bearer.go](./server/session/bearer.go) file.

This function runs whenever you declare a path to be required for BearerAuth, such as the [protected route](./server/openapi/paths/protected.yml#L10)

#### RBAC

Much like Bearer, adding the Rbac security scheme to a path makes the Rbac function run and must pass for your application code to run.

The code for the RBAC middleware is in the [rbac.go](./server/session/rbac.go) file.

The [RBAC route](./server/openapi/paths/rbac.yml#L11) declares both `BearerAuth` and `Rbac` security. This means that both functions **MUST** pass in order for application code to run.

### Managing the policy assignments

Some work will need to be done in order to figure out the best way to manage the policies. They are read at server start. There are some adapters to connect to a DB, but that's outside this demo scope.

## Testing

Run the server with `make run-dev`.

In the [docs/](./docs/) dir, there is a [Bruno](https://www.usebruno.com/) API suite to test the various scenarios.

Alternatively you can curl these routes with the applicable headers.

| Scenario                                    | Passes             |
| ------------------------------------------- | ------------------ |
| GET to `/unprotected`                              | :white_check_mark: |
| GET to `/protected` without a token                | :x:                |
| GET to `/protected` with a bad token               | :x:                |
| GET to `/protected` with a good token (123 or 456) | :white_check_mark: |
| GET to `/rbac` without a token              | :x:                |
| GET to `/rbac` with a token (123)           | :white_check_mark: |
| GET to `/rbac` with a token (456)           | :white_check_mark: |
| POST to `/rbac` with a token (123)          | :x:                |
| POST to `/rbac` with a token (456)          | :white_check_mark: |
