package main

import (
	"fmt"
	"log"
	"os"

	"github.com/thycotic/dsv-sdk-go/vault"
)

func main() {
	dsv, err := vault.New(vault.Configuration{
		Credentials: vault.ClientCredential{
			ClientID:     os.Getenv("DSV_CLIENT_ID"),
			ClientSecret: os.Getenv("DSV_CLIENT_SECRET"),
		},
		Tenant: os.Getenv("DSV_TENANT"),
	})
	secret, err := dsv.Secret("path:of:the:secret")

	if err != nil {
		log.Fatal("failure calling vault.Secret", err)
	}

	fmt.Print("the SSH public key is", secret.Data["public"])
}
