package vault

import (
	"encoding/json"
	"log"
	"net/http"
)

// secretsResource is the HTTP URL path component for the secrets resource
const secretsResource = "secrets"

// secretResource is composed with resourceMetadata to for SecretContents
type secretResource struct {
	Attributes map[string]interface{}
	Data       map[string]interface{}
	Path       string
}

// Secret holds the contents of a secret from DSV
type Secret struct {
	resourceMetadata
	secretResource
}

// Secret gets the secret at path from the DSV of the given tenant
func (v Vault) Secret(path string) (*Secret, error) {
	data, err := v.accessResource(http.MethodGet, secretsResource, path, nil)
	if err != nil {
		return nil, err
	}

	secret := &Secret{}
	if err := json.Unmarshal(data, secret); err != nil {
		log.Printf("[DEBUG] error parsing response from /%s/%s: %q", secretsResource, path, data)
		return nil, err
	}
	return secret, nil
}
