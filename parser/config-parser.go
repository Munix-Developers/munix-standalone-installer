package parser

import (
	"encoding/json"
)

// Reads json bytes to a InstallConfig struct
func ReadConfig(bytes []byte, config *InstallConfig) error {
	return json.Unmarshal(bytes, config)
}
