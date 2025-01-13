# Authentication: Azure
You can use the DSV web UI or the DSV cli to configure authentication using Azure.
For this doc we will use the DSV cli.

Run <br />
`dsv config auth-provider search -e yaml` 
<br />
to see all of your current authentication providers.

Initially, the only authentication provider is Thycotic One, similar to this:
```
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

To add an Azure account to act as an authentication provider:<br />
`dsv config auth-provider create --name <name> --type azure --azure-tenant-id <Azure tenant ID>`
where:

* name is the friendly name used in DSV to reference this provider
* type is the authentication provider type; in this case, azure
* the property flag for Azure is `--azure-tenant-id`

To view the resulting addition to the config file, you would use:<br />
`dsv config auth-provider <name> read -e yaml` <br />
where the example name we will use here is azure-prod

The readout would look similar to this:
```
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

First you will need to configure the User that corresponds to an [Azure User Assigned MSI](https://docs.microsoft.com/en-us/azure/active-directory/managed-identities-azure-resources/overview).
The username is a friendly name within DSV. It does not have to match the MSI username, but the provider must match the resource id of the MSI in Azure.<br />
`dsv user create --username test-api --provider azure-prod --external-id /subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourcegroups/build/providers/Microsoft.ManagedIdentity/userAssignedIdentities/test-api`<br />

## DSV Azure code example

```
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
			ClientID:     os.Getenv("AZURE_CLIENT_ID"), // CLIENT_ID of the MSI identity you wish to use
		},
		Tenant: os.Getenv("DSV_TENANT"), // your tenant name 
		TLD:    os.Getenv("DSV_TLD"), // defaults to com change if your domain is au, eu etc  
		Provider: auth.AZURE, // required to enable Azure authentication
	})
	if err != nil {
		log.Fatalf("failed to configure vault: %v", err)
	}

	secret, err := dsv.Secret("<secret path or ID")
	if err != nil {
		log.Fatalf("failed to fetch secret: %v", err)
	}
	fmt.Printf("\nsecret data: %v\n\n", secret.Data)
}
```