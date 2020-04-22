package validator

import (
	"errors"
	"net.matbm/munix/installer/parser"
)

const version = "1.0"

func ValidateConfig(config parser.InstallConfig) (bool, error) {
	if config.Version != "1.0" {
		return false, errors.New("version: " + "only " + version + " accepted")
	}
	return true, nil
}
