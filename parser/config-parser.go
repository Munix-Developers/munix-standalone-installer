package parser

import (
	"encoding/json"
)

func ReadConfig(bytes []byte, config *InstallConfig) error {
	return json.Unmarshal(bytes, config)
}
