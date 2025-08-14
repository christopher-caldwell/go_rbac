Hello, I am not sure if this library is still maintained, but I am looking to get the enforcer setup client side.

## Versions

- casbin.js: "^0.5.1"
- github.com/casbin/casbin/v2 v2.116.0

The API is in go, and the server side model & policy work as expected. I am sending the result of `CasbinJsGetPermissionForUser`, but when using the client enforcer, it says false for the same comparison I am doing in the backend.

When I change it to what the example says:

```json
{
  "read": ["data1", "data2"],
  "write": ["data1"]
}
```

it works as expected.

<details>
  <summary>Model</summary>

```conf
[request_definition]
r = sub, act, obj

[policy_definition]
p = sub, act, obj

[role_definition]
g = _, _
g2 = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && g(p.act, r.act) && r.obj == p.obj
```

</details>
<details>
  <summary>Policy</summary>

```csv
p, 123, user, data1
p, 456, admin, data1

g, user, GetRbacResource
g, admin, GetRbacResource
g, admin,CreateRbacResource

```

</details>
<details>
  <summary>Result of CasbinJsGetPermissionForUser</summary>

```json
{
  "g": [
    [
      "g",
      "user",
      "GetRbacResource"
    ],
    [
      "g",
      "admin",
      "GetRbacResource"
    ],
    [
      "g",
      "admin",
      "CreateRbacResource"
    ]
  ],
  "m": "[request_definition]\nr = sub, act, obj\n[policy_definition]\np = sub, act, obj\n[role_definition]\ng = _, _\ng2 = _, _\n[policy_effect]\ne = some(where (p.eft == allow))\n[matchers]\nm = r.sub == p.sub && g(p.act, r.act) && r.obj == p.obj\n",
  "p": [
    [
      "p",
      "123",
      "user",
      "data1"
    ],
    [
      "p",
      "456",
      "admin",
      "data1"
    ]
  ]
}
```

</details>


## Usage

```ts

const authorizer = new casbinjs.Authorizer('manual')

// get permissions from API - responds with above block
authorizer.setPermission(responseFromApi)

// tried this too, but had no effect
// authorizer.setUser('123')

authorizer.can('CreateRbacResource', 'data1').then((can) => {
  // can is false
})
```

When I swapped out the response from the API with the docs snippet, it worked as expected.

```ts
// drastically different response to what I get
authorizer.setPermission({
  read: ['data1', 'data2'],
  write: ['data1'],
})

authorizer.can('write', 'data1').then((can) => {
  // can is true
})
```