package vault

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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
