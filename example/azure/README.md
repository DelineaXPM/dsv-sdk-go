# DSV Azure Authentication

- [DSV Overview](#dsv-overview)
- [DSV Authentication Provider](#Azure-Authentication-Provider)
- [Azure MSI Example](#Azure-User-Assigned-MSI-Example)
- [Azure MSI Code](#DSV-Azure-msi-code-example)
- [Azure Entra App Example](#Azure-Microsoft-Entra-App-example)
- [Azure Entra Code Example](#Azure-Microsoft-Entra-Code-Example)

## DSV Overview

You can use the DSV web UI or the DSV cli to configure authentication using Azure.
For this document we will use the DSV cli.
To begin using Azure authentication you need to configure a DSV authentication provider.\

To see all of your current authentication providers.
Run\
`dsv config auth-provider search -e yaml`\

Initially, the only authentication provider is Thycotic One, similar to this:

```Bash
created: "2019-11-11T20:29:20Z"
createdBy: users:thy-one:admin@company.com
id: xxxxxxxxxxxxxxxxxxxx
lastModified: "2020-05-18T03:58:15Z"
lastModifiedBy: users:thy-one:admin@company.com
name: thy-one
properties:
 baseUri: https://login.thycotic.com/
 clientId: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
 clientSecret: xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
type: thycoticone
version: "0"
```

## Azure Authentication Provider

To add an Azure account to act as an authentication provider run:\
`dsv config auth-provider create --name <name> --type azure --azure-tenant-id <Azure tenant ID>`
where:

- name is the friendly name used in DSV to reference this provider
- type is the authentication provider type; in this case, azure
- the property flag for Azure is `--azure-tenant-id`
  `*` You will need to get your tenant-id from your Azure account.

To view the resulting addition to the config file, you would use:\
`dsv config auth-provider <name> read -e yaml`\
where the example name we will use here is azure-prod

The readout would look similar to this:

```Bash
created: "2019-11-12T18:34:49Z"
createdBy: users:thy-one:admin@company.com
-id: xxxxxxxxxxxxxxxxxxxxx
lastModified: "2020-05-18T03:58:15Z"
lastModifiedBy: users:thy-one:admin@company.com
name: azure-prod
properties:
 tenantId: xxxxxxxxxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx
type: azure
version: "0"
```

## Azure User Assigned MSI Example

When using dsv-sdk-go from an Azure resouce i.e. Azure VM or Kubernetes service (AKS).\
First you will need to configure the User that corresponds to an [Azure User Assigned MSI](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview).
The username is a friendly name within DSV. It does not have to match the MSI username, but the provider must match the resource id of the MSI in Azure.\
`dsv user create --username test-api --provider azure-prod --external-id /subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourcegroups/build/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-api`\

When you have successfully created a DSV user with an Azure authentication provider you can use the dsv-sdk-go as in the following example code.

## DSV Azure msi code example

```Golang
package main

import (
        "fmt"
        "log"
        "os"

        "github.com/DelineaXPM/dsv-sdk-go/v2/auth"
        "github.com/DelineaXPM/dsv-sdk-go/v2/vault"
)

// Azure authentication example
func main() {
        dsv, err := vault.New(vault.Configuration{
                Credentials: vault.ClientCredential{
                    ClientID: os.Getenv("AZURE_CLIENT_ID"), // CLIENT_ID of the MSI identity you wish to use
                },
                Tenant: os.Getenv("DSV_TENANT"), // your tenant name
                TLD:    os.Getenv("DSV_TLD"), // defaults to com change if your domain is au, eu etc
                Provider: auth.AZURE, // required to enable Azure authentication
        })
        if err != nil {
                log.Fatalf("failed to configure vault: %v", err)
        }

        secret, err := dsv.Secret("<secret path or ID>")
        if err != nil {
                log.Fatalf("failed to fetch secret: %v", err)
        }
        fmt.Printf("\nsecret data: %v\n\n", secret.Data)
}
```

## Azure Microsoft Entra App example:

When using Azure authentication from an Non Azure resouce i.e. AWS, GCP or local development.\

First you will need to configure the User that corresponds to an [Azure Service Principal App](https://learn.microsoft.com/en-us/azure/developer/go/sdk/authentication/authentication-on-premises-apps?tabs=azure-cli%2Cbash).
After you have created and configured you service principal save the credentials (you will need them for configuration). You will also need to retrieve the "Object ID" for the service principle you just created

The username is a friendly name within DSV. It does not have to match the MSI username, but the provider must match the resource id of the MSI in Azure.\
`dsv user create --username test-api --provider azure-prod --external-id xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx`\
Where:
--external-id = "Object Id" (from the service principal you created)

When you have successfully created a DSV user with an Azure authentication provider you can use the dsv-sdk-go as in the following example code.

## Azure Microsoft Entra Code Example:

```golang
package main

import (
	"fmt"
	"log"

	"github.com/DelineaXPM/dsv-sdk-go/v2/vault"
)

func main() {
	// Azure authentication
	dsv, err := vault.New(vault.Configuration{
		Credentials: vault.ClientCredential{
			ClientID:     "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
			ClientSecret: "zzzz~zzzzzz.zzzzzzzzzzzzzzzzzzzzzzzz",
		},
		Tenant:      "<yourTenantName>",
		TenantID:    "xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx",
		TLD:         "com", // replace with au, eu as necessary
		Provider:    2,
	})

	if err != nil {
		log.Fatalf("failed to configure vault: %v", err)
	}

	secret, err := dsv.Secret("<secretPathORIdentifierHere")

	if err != nil {
		log.Fatalf("failed to fetch secret: %v", err)
	}

	fmt.Printf("\nsecret data: %v\n\n", secret.Data)
}
```
