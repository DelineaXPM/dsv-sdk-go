package vault

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	urlTemplate string = "https://%s.secretsvaultcloud.com/v1/%s%s"
)

// Configuration used to request an accessToken for the API
type Configuration struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Tenant       string
}

// Vault provides access to secrets stored in Thycotic DSV
type Vault struct {
	config Configuration
}

// New returns an initialized Secrets object
func New(config Configuration) *Vault {
	return &Vault{config: config}
}

// Each resource (Secret, Role, ...) embeds this in it's struct
type resourceMetadata struct {
	ID, Description           string
	Created, LastModified     time.Time
	CreatedBy, LastModifiedBy string
	Version                   string
}

// Each simple resource (Client, ...) embeds this in it's struct instead
type simpleResourceMetadata struct {
	ID        string `json:"id"`
	Created   time.Time
	CreatedBy string
}

// accessResource uses the accessToken to access the API resource.
// It assumes an appropriate combination of method, resource, path and input.
func accessResource(method, resource, path string, input interface{}, config Configuration) ([]byte, error) {
	switch resource {
	case "clients", "roles", "secrets":
	default:
		message := "unrecognized resource"

		log.Printf("[DEBUG] %s: %s", message, resource)
		return nil, fmt.Errorf(message)
	}

	if path != "" {
		path = "/" + strings.TrimLeft(path, "/")
	}

	url := fmt.Sprintf(urlTemplate, config.Tenant, resource, path)
	body := bytes.NewBuffer([]byte{})

	if input != nil {
		if data, err := json.Marshal(input); err == nil {
			body = bytes.NewBuffer(data)
		} else {
			log.Print("[DEBUG] marshaling the request body to JSON:", err)
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, body)

	if err != nil {
		log.Printf("[DEBUG] creating req: %s /%s/%s: %s", method, resource, path, err)
		return nil, err
	}

	accessToken, err := getAccessToken(config)

	if err != nil {
		log.Print("[DEBUG] error getting accessToken:", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	switch method {
	case "POST", "PUT":
		req.Header.Set("Content-Type", "application/json")
	}

	log.Printf("[DEBUG] calling %s", req.URL.String())

	data, _, err := handleResponse((&http.Client{}).Do(req))

	return data, err
}

type accessGrant struct {
	AccessToken, TokenType string
	ExpiresIn              int
}

var grant *accessGrant // TODO proper caching and expiration checking

// getAccessToken uses the client_id and client_secret, to call the token
// endpoint and get an accessGrant.
func getAccessToken(config Configuration) (string, error) {
	if grant != nil {
		return grant.AccessToken, nil
	}

	grantRequest, err := json.Marshal(struct {
		GrantType    string `json:"grant_type"`
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
	}{
		"client_credentials",
		config.ClientID,
		config.ClientSecret,
	})

	if err != nil {
		log.Print("[INFO] marshaling grantRequest")
		return "", err
	}

	url := fmt.Sprintf(urlTemplate, config.Tenant, "token", "")

	log.Printf("[DEBUG] calling %s with client_id %s", url, config.ClientID)

	data, _, err := handleResponse(http.Post(url, "application/json",
		bytes.NewReader(grantRequest)))

	if err != nil {
		log.Print("[DEBUG] grant response error:", err)
		return "", err
	}

	newGrant := new(accessGrant)

	if err = json.Unmarshal(data, &newGrant); err != nil {
		log.Print("[INFO] parsing grant response:", err)
		return "", err
	}

	grant = newGrant

	return grant.AccessToken, nil
}

// handleResponse processes the response according to the HTTP status
func handleResponse(res *http.Response, err error) ([]byte, *http.Response, error) {
	if err != nil { // fall-through if there was an underlying err
		return nil, res, err
	}

	data, err := ioutil.ReadAll(res.Body)

	if err != nil {
		return nil, res, err
	}

	// if the response was 2xx then return it, otherwise, consider it an error
	if res.StatusCode > 199 && res.StatusCode < 300 {
		return data, res, nil
	}

	// truncate the data to 64 bytes before returning it as part of the error
	if len(data) > 64 {
		data = append(data[:64], []byte("...")...)
	}

	return nil, res, fmt.Errorf("%s: %s", res.Status, string(data))
}
