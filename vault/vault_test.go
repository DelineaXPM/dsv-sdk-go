//go:build integration

package vault

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

var config = func() *Configuration {
	if cj, err := ioutil.ReadFile("../test_config.json"); err == nil {
		c := new(Configuration)

		_ = json.Unmarshal(cj, &c) //nolint:musttag //dynamic struct
		return c
	}
	return &Configuration{
		Tenant: os.Getenv("DSV_TENANT"),
		Credentials: ClientCredential{
			ClientID:     os.Getenv("DSV_CLIENT_ID"),
			ClientSecret: os.Getenv("DSV_CLIENT_SECRET"),
		},
	}
}()
var dsv, _ = New(*config)
