//go:build integration
package vault

import "testing"

const secretName = "/test/secret"

// TestSecret tests Secret
func TestSecret(t *testing.T) {
	secret, err := dsv.Secret(secretName)

	if err != nil {
		t.Error("calling secrets.Secret:", err)
		return
	}

	if secret.Data == nil {
		t.Error("secret.Data is nil")
	}
}
