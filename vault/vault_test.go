package vault

import (
	"encoding/json"
	"io/ioutil"
)

const ConfigFile = "../test_config.json"

var config = func() *Configuration {
	cj, err := ioutil.ReadFile(ConfigFile)

	if err != nil {
		return nil
	}

	config := new(Configuration)

	if err := json.Unmarshal(cj, config); err == nil {
		return config
	}
	return nil
}()
var dsv, _ = New(*config)
