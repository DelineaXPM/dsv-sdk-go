package main

import (
	"fmt"
	"log"
	"os"

	"github.com/DelineaXPM/dsv-sdk-go/v2/vault"
)

func main() {
	dsv, err := vault.New(vault.Configuration{
		Credentials: vault.ClientCredential{
			ClientID:     os.Getenv("DSV_CLIENT_ID"),
			ClientSecret: os.Getenv("DSV_CLIENT_SECRET"),
		},
		Tenant: os.Getenv("DSV_TENANT"),
		TLD:    os.Getenv("DSV_TLD"),
	})

	if err != nil {
		log.Fatalf("failed to configure vault: %v", err)
	}

	secret, err := dsv.Secret("your secret path")

	if err != nil {
		log.Fatalf("failed to fetch secret: %v", err)
	}

	fmt.Printf("secret data: %v", secret.Data)
}