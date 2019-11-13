package vault

import (
	"encoding/json"
	"log"
)

// secretsResource is the HTTP URL path component for the secrets resource
const secretsResource = "secrets"

// secretResource is composed with resourceMetadata to for SecretContents
type secretResource struct {
	Attributes []string
	Data       map[string]string
	Path       string
}

// Secret holds the contents of a secret from DSV
type Secret struct {
	resourceMetadata
	secretResource
}

// Secret gets the secret at path from the DSV of the given tenant
func (v Vault) Secret(path string) (*Secret, error) {
	secret := new(Secret)
	data, err := accessResource("GET", secretsResource, path, nil, v.config)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, secret); err != nil {
		log.Printf("[DEBUG] error parsing response from /%s/%s: %q", secretsResource, path, data)
		return nil, err
	}
	return secret, nil
}
