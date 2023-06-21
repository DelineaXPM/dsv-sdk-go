//go:build integration
package vault

import (
	"math/rand"
	"testing"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func TestSecret(t *testing.T) {
	path := makeRandomSecretPath()
	_, cleanup := createSecret(t, path)
	defer cleanup()

	secret, err := dsv.Secret(path)

	if err != nil {
		t.Fatalf("Secret for path=%s: %s", path, err)
		return
	}

	if secret.Data == nil {
		t.Error("secret.Data is nil")
	}
}

func TestCreateSecret(t *testing.T) {
	path := makeRandomSecretPath()
	secret, err := dsv.CreateSecret(
		path, &SecretCreateRequest{Data: map[string]interface{}{"foo": "bar"}},
	)

	if err != nil {
		t.Fatalf("CreateSecret for path=%s: %s", path, err)
		return
	} else {
		defer func() { deleteSecret(t, path) }()
	}

	if secret.Data == nil {
		t.Error("secret.Data is nil")
	}
}

func TestDeleteSecret(t *testing.T) {
	path := makeRandomSecretPath()
	_, _ = createSecret(t, path) // no cleanup required as test tries to delete anyway

	err := dsv.DeleteSecret(path)
	if err != nil {
		t.Errorf("DeleteSecret for path=%s: %s", path, err)
	}
}

// createSecret creates a secret with given path and returns the created secret as well
// as a cleanup function which should be deferred to remove the secret from the vault.
func createSecret(t *testing.T, path string) (s *Secret, cleanup func()) {
	t.Helper()
	s, err := dsv.CreateSecret(path, &SecretCreateRequest{Data: map[string]interface{}{
		"foo": "bar",
	}})
	if err != nil {
		t.Fatal(err)
	}
	cleanup = func() { deleteSecret(t, path) }
	return
}

// deleteSecret deletes the secret given by path from the vault.
func deleteSecret(t *testing.T, path string) {
	t.Helper()
	if err := dsv.DeleteSecret(path); err != nil {
		t.Fatal(err)
	}
}

// makeRandomSecretPath creates a pseudo-random secret path.
func makeRandomSecretPath() string {
	b := make([]byte, 10)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return "/test/" + string(b)
}
