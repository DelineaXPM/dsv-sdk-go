//go:build integration
package vault

import "testing"

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
		client, err := dsv.Client(config.Credentials.ClientID)

		if err != nil {
			t.Errorf("calling clients.Client(\"%s\"): %s", ID, err)
			return
		}

		if client.ClientID != config.Credentials.ClientID {
			t.Errorf("expecting %s but clients.Client was %s", ID, config.Credentials.ClientID)
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
