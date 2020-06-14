package validator

import (
	"github.com/stretchr/testify/assert"
	"net.matbm/munix/muinstaller/parser"
	"testing"
)

const rightVersion = "1.1"
const wrongVersion = "potatoe"

func TestAcceptValidConfig(t *testing.T) {
	a := assert.New(t)

	config := validConfig()
	err := ValidateConfig(config)

	a.NoError(err)
}

func TestWrongVersionError(t *testing.T) {
	a := assert.New(t)

	config := validConfig()
	config.Version = wrongVersion
	err := ValidateConfig(config)

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
							Type:           "ext4",
							Mount:          "/boot",
							StartMegaBytes: 0,
							SizeMegaBytes:  2.56e+8,
						},
						{
							Type:           "ext4",
							Mount:          "/",
							StartMegaBytes: 2.56e+8 + 1,
							SizeMegaBytes:  5e+10,
						},
						{
							Type:           "ext4",
							Mount:          "/home",
							StartMegaBytes: 5e+10 + 1,
							SizeMegaBytes:  4.5e+11,
						},
					},
				},
			},
		},
	}
}
