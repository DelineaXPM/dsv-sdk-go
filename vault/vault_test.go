package vault

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

const (
	ConfigFile = "../test_config.json"
	roleName   = "test-role"
	secretName = "/test/secret"
)

var config = func() *Configuration {
	cj, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		return nil
	}

	config := new(Configuration)

	if err := json.Unmarshal(cj, &config); err == nil {
		return config
	}
	return nil
}()
var dsv = New(*config)

func TestClient(t *testing.T) {
	var ID string // set by TestNewClient used by Get and removed by Delete

	t.Run("TestNewClient", func(t *testing.T) {
		client := &Client{clientResource: clientResource{RoleName: roleName}}
		err := dsv.New(client)

		if err != nil {
			t.Errorf("calling clients.New(\"%s\"): %s", roleName, err)
			return
		}

		if client.ClientID == "" {
			t.Error("contents.ClientID was empty")
			return
		}
		ID = client.ClientID
	})
	t.Run("TestGetClient", func(t *testing.T) {
		client, err := dsv.Client(config.ClientID)

		if err != nil {
			t.Errorf("calling clients.Client(\"%s\"): %s", ID, err)
			return
		}

		if client.ClientID != config.ClientID {
			t.Errorf("expecting %s but clients.Client was %s", ID, config.ClientID)
			return
		}
	})
	t.Run("TestDeleteClient", func(t *testing.T) {
		client, err := dsv.Client(ID)

		if err != nil {
			t.Errorf("calling clients.Client(\"%s\"): %s", ID, err)
			return
		}

		if err := client.Delete(); err != nil {
			t.Errorf("calling client.Delete on Client %s: %s", ID, err)
			return
		}
	})
}

// TestRole tests Role
func TestRole(t *testing.T) {
	role, err := dsv.Role(roleName)

	if err != nil {
		t.Errorf("calling roles.Role(\"%s\"): %s", roleName, err)
		return
	}

	if role.Name != roleName {
		t.Errorf("expecting %s but roles.Role was %s", roleName, role.Name)
		return
	}
}

// TestNonexistentRole asks for the "nonexistent" role and fails if it gets it
func TestNonexistentRole(t *testing.T) {
	roleName := "nonexistent"
	_, err := dsv.Role(roleName)

	if err == nil {
		t.Errorf("role '%s' exists but but it should not", roleName)
		return
	}
}

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
