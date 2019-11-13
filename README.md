# The Thycotic DevOps Secrets Vault SDK for Go

A Golang API and examples for [Thycotic](https://thycotic.com/)
[DevOps Secrets Vault](https://thycotic.com/products/devops-secrets-vault-password-management/).

## Configure

The API requires a `Configuration` object containing a `ClientID`, `ClientSecret`
and `Tenant`.

For example, the tests populates `Configuration` from JSON:

```golang
config := new(Configuration)

if cj, err := ioutil.ReadFile("../test_config.json"); err == nil {
    json.Unmarshal(cj, &config)
}

tss := New(*config)
```

Example JSON configuration:

```json
{
    "client_id": "93d866d4-635f-4d4e-9ce3-0ef7f879f319",
    "client_secret": "xxxxxxxxxxxxxxxxxxxxxxxxx-xxxxxxxxxxx-xxxxx",
    "tenant": "mytenant"
}
```

## Test

`vault/vault_test.go` declares:

```golang
const (
	ConfigFile = "../test_config.json"
	roleName   = "test-role"
	secretName = "/test/secret"
)
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
