package main

import (
	"fmt"
	"os"

	"github.com/DelineaXPM/dsv-sdk-go/v2/vault"
)

func main() {
	var exitCode = 1
	dsv, err := vault.New(vault.Configuration{
		Credentials: vault.ClientCredential{
			ClientID:     os.Getenv("DSV_CLIENT_ID"),
			ClientSecret: os.Getenv("DSV_CLIENT_SECRET"),
		},
		Tenant: os.Getenv("DSV_TENANT"),
		TLD:    os.Getenv("DSV_TLD"),
	})
	if err != nil {
		_ = fmt.Printf("failed to configure vault: %v", err)
		os.Exit(exitCode)
	}

	secret, err := dsv.Secret("your secret path")
	if err != nil {
		_ = fmt.Printf("failed to fetch secret: %v", err)
		os.Exit(exitCode)
	}
	fmt.Printf("secret data: %v", secret.Data)
}
