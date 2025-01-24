package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DelineaXPM/dsv-sdk-go/v2/auth"
	"github.com/DelineaXPM/dsv-sdk-go/v2/vault"
)

//nolint:typecheck //example code
func main() {
	// Azure authentication
	dsv, err := vault.New(vault.Configuration{
		Credentials: vault.ClientCredential{
			ClientID:     os.Getenv("AZURE_CLIENT_ID"),
			ClientSecret: os.Getenv("DSV_CLIENT_SECRET"),
		},
		URLTemplate: "https://%s.devbambe.%s/v1/%s%s",
		Tenant:      os.Getenv("DSV_TENANT"),
		TLD:         os.Getenv("DSV_TLD"),
		Provider:    auth.AZURE,
	})

	if err != nil {
		log.Fatalf("failed to configure vault: %v", err)
	}

	secret, err := dsv.Secret("5f7af80f-d786-4fae-8af9-c9744e8809d9")

	if err != nil {
		log.Fatalf("failed to fetch secret: %v", err)
	}

	fmt.Printf("\nsecret data: %v\n\n", secret.Data)
}
