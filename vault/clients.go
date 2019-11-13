package vault

import (
	"encoding/json"
	"log"
)

// clientsResource is the HTTP URL path component for the clients resource
const clientsResource = "clients"

// Clients provides access to clients configured in Thycotic DSV
type Clients struct {
	clientCredential Configuration
	tenant           string
}

// clientResource is composed with resourceMetadata to for ClientContents
type clientResource struct {
	ClientID     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	RoleName     string `json:"role"`
}

// Client holds the contents of a client from DSV
type Client struct {
	simpleResourceMetadata
	clientResource
	config Configuration
}

// Client gets the client with id from the DSV of the given tenant
func (v Vault) Client(id string) (*Client, error) {
	client := new(Client)
	client.config = v.config
	data, err := accessResource("GET", clientsResource, id, nil, v.config)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, client); err != nil {
		log.Printf("[DEBUG] error parsing response from /%s/%s: %q", clientsResource, id, data)
		return nil, err
	}
	return client, nil
}

// Delete deletes the client from the DSV of the given tenant
func (c Client) Delete() error {
	if _, err := accessResource("DELETE", clientsResource, c.ClientID, nil, c.config); err != nil {
		return err
	}

	return nil
}

// New creates a new Client given a roleName
func (v Vault) New(client *Client) error {
	data, err := accessResource("POST", clientsResource, "/", client, v.config)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(data, client); err != nil {
		log.Printf("[DEBUG] error parsing response from /%s: %q", clientsResource, data)
		return err
	}

	return nil
}
