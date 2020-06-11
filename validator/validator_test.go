package validator

import (
	"github.com/stretchr/testify/assert"
	"net.matbm/munix/installer/parser"
	"testing"
)

const rightVersion = "1.0"
const wrongVersion = "potatoe"

func TestAcceptValidConfig(t *testing.T) {
	a := assert.New(t)

	config := validConfig()
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

func TestLocaleExists(t *testing.T) {
	t.Skip("Not implemented yet.")
}

func validConfig() parser.InstallConfig {
	return parser.InstallConfig{
		Version: rightVersion,
		Keyboard: parser.KeyboardConfig{
			Layout: "br-abnt2",
		},
		Localization: parser.LocalizationConfig{
			SystemLanguage: "en_US.UTF-8",
			Locales: []parser.LocaleConfig{
				{
					Language: "pt_BR.UTF-8",
					Encoding: "UTF-8",
				},
				{
					Language: "pt_BR",
					Encoding: "ISO-8859-1",
				},
				{
					Language: "en_US.UTF-8",
					Encoding: "UTF-8",
				},
			},
			Timezone: "America/Sao_Paulo",
		},
		Computer: parser.ComputerConfig{
			Hostname: "crazywriter",
			Username: "mat",
			Password: "231412341235123",
		},
		Storage: parser.StorageConfig{
			Devices: []parser.DeviceConfig{
				{
					Device: "/dev/sda2",
					Partitions: []parser.PartitionConfig{
						{
							Type:        "ext4",
							Mount:       "/boot",
							OffsetBytes: 0,
							SizeBytes:   2.56e+8,
						},
						{
							Type:        "ext4",
							Mount:       "/",
							OffsetBytes: 2.56e+8 + 1,
							SizeBytes:   5e+10,
						},
						{
							Type:        "ext4",
							Mount:       "/home",
							OffsetBytes: 5e+10 + 1,
							SizeBytes:   4.5e+11,
						},
					},
				},
			},
		},
	}
}
