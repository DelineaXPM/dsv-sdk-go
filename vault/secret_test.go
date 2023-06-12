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


func TestCreateSecret(t *testing.T) {
	secret, err := createSecret(secretName)

	if err != nil {
		t.Error("calling secrets.CreateSecret:", err)
		return
	}

	if secret.Data == nil {
		t.Error("secret.Data is nil")
	}
}

func TestDeleteSecret(t *testing.T) {
	createSecret("temporary")
	err := dsv.DeleteSecret("temporary")

	if err != nil {
		t.Error("calling secrets.DeleteSecret:", err)
		return
	}
}

func createSecret(name string) (*Secret, error) {
	return dsv.CreateSecret(name, SecretRequest{Data: map[string]interface{}{
		"foo": "bar",
	}})
}
