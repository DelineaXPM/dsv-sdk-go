# The Delinea DevOps Secrets Vault SDK for Go

![Tests](https://github.com/DelineaXPM/dsv-sdk-go/workflows/Tests/badge.svg)

A Golang API and examples for [Delinea](https://delinea.com/)
[DevOps Secrets Vault](https://delinea.com/products/devops-secrets-management-vault).

## Configure

The API requires a `Configuration` object containing a `ClientID`, `ClientSecret`
and `Tenant`:

```golang
type ClientCredential struct {
    ClientID, ClientSecret string
}

type Configuration struct {
    Credentials              ClientCredential
    Tenant, TLD, URLTemplate string
}
```

The unit tests populate `Configuration` from `test_config.json`:

```golang
config := new(Configuration)

if cj, err := ioutil.ReadFile("../test_config.json"); err == nil {
    json.Unmarshal(cj, &config)
}

tss := New(*config)
```

Create `test_config.json`:

```json
{
  "credentials": {
    "clientId": "",
    "clientSecret": ""
  },
  "tenant": "mytenant"
}
```

## Test

`vault/role_test.go` declares:

```golang
const roleName = "test-role"
```

`vault/secret_test.go` declares:

```golang
const secretName = "/test/secret"
```

The tests assume that `roleName` can exists and has privilege to create, read,
and delete a client, and read `secretName`.

## Use

Define a `Configuration` then use it to create an instance of `Vault`:

```golang
dsv := vault.New(vault.Configuration{
    ClientID:     os.Getenv("DSV_CLIENT_ID"),
    ClientSecret: os.Getenv("DSV_CLIENT_SECRET"),
    Tenant:       os.Getenv("DSV_TENANT"),
})
secret, err := dsv.Secret("path:of:the:secret")

if err != nil {
    log.Fatal("failure calling vault.Secret", err)
}

fmt.Print("the SSH public key is", secret.Data["public"])
```
