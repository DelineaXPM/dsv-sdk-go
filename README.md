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

The tests assume that `roleName` exists and has privilege to create, read,
and delete a client, as well as create, read and delete secrets with a `test:` path prefix.

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

## Contributors

<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- readme: collaborators,contributors -start -->
<table>
	<tbody>
		<tr>
            <td align="center">
                <a href="https://github.com/sheldonhull">
                    <img src="https://avatars.githubusercontent.com/u/3526320?v=4" width="100;" alt="sheldonhull"/>
                    <br />
                    <sub><b>Sheldonhull</b></sub>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/amigus">
                    <img src="https://avatars.githubusercontent.com/u/119477?v=4" width="100;" alt="amigus"/>
                    <br />
                    <sub><b>Adam C. Migus</b></sub>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/pacificcode">
                    <img src="https://avatars.githubusercontent.com/u/918320?v=4" width="100;" alt="pacificcode"/>
                    <br />
                    <sub><b>Bill Hamilton</b></sub>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/maxsokolovsky">
                    <img src="https://avatars.githubusercontent.com/u/17733533?v=4" width="100;" alt="maxsokolovsky"/>
                    <br />
                    <sub><b>Max Sokolovsky</b></sub>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/delineaKrehl">
                    <img src="https://avatars.githubusercontent.com/u/105234788?v=4" width="100;" alt="delineaKrehl"/>
                    <br />
                    <sub><b>Tim Krehl</b></sub>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/andrii-zakurenyi">
                    <img src="https://avatars.githubusercontent.com/u/85106843?v=4" width="100;" alt="andrii-zakurenyi"/>
                    <br />
                    <sub><b>Andrii Zakurenyi</b></sub>
                </a>
            </td>
		</tr>
		<tr>
            <td align="center">
                <a href="https://github.com/michaelsauter">
                    <img src="https://avatars.githubusercontent.com/u/215455?v=4" width="100;" alt="michaelsauter"/>
                    <br />
                    <sub><b>Michael Sauter</b></sub>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/EndlessTrax">
                    <img src="https://avatars.githubusercontent.com/u/17141891?v=4" width="100;" alt="EndlessTrax"/>
                    <br />
                    <sub><b>Ricky White</b></sub>
                </a>
            </td>
            <td align="center">
                <a href="https://github.com/timothyfield">
                    <img src="https://avatars.githubusercontent.com/u/12048504?v=4" width="100;" alt="timothyfield"/>
                    <br />
                    <sub><b>Tim Field</b></sub>
                </a>
            </td>
		</tr>
	<tbody>
</table>
<!-- readme: collaborators,contributors -end -->

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->
