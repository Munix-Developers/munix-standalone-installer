package validator

import (
	"github.com/stretchr/testify/assert"
	"net.matbm/munix/installer/parser"
	"testing"
)

const rightVersion = "1.0"
const wrongVersion = "potatoe"

func TestAcceptTheRightVersion(t *testing.T) {
	a := assert.New(t)

	config := parser.InstallConfig{Version: rightVersion}
	isValid, err := ValidateConfig(config)

	a.True(isValid)
	a.NoError(err)
}

func TestWrongVersionError(t *testing.T) {
	a := assert.New(t)

	config := parser.InstallConfig{Version: wrongVersion}
	isValid, err := ValidateConfig(config)

	a.False(isValid)
	a.Error(err)
}
