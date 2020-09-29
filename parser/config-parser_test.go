package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfigJsonLoad(t *testing.T) {
	a := assert.New(t)

	var config InstallConfig
	err := ReadConfig([]byte(validJson), &config)

	a.NoError(err)
	a.Equal("version", config.Version)
	a.Equal("layout", config.Keyboard.Layout)

	a.Equal("system-language", config.Localization.SystemLanguage)
	a.Equal("language", config.Localization.Locales[0].Language)
	a.Equal("encoding", config.Localization.Locales[0].Encoding)
	a.Equal("language", config.Localization.Locales[1].Language)
	a.Equal("encoding", config.Localization.Locales[1].Encoding)
	a.Equal("timezone", config.Localization.Timezone)

	a.Equal("hostname", config.Computer.Hostname)
	a.Equal("username", config.Computer.Username)
	a.Equal("password", config.Computer.Password)

	a.Equal("device", config.Storage.Devices[0].Device)
	a.Equal("type", config.Storage.Devices[0].Partitions[0].Type)
	a.Equal("mount", config.Storage.Devices[0].Partitions[0].Mount)
	a.Equal(false, config.Storage.Devices[0].Partitions[0].Boot)
	a.Equal(uint64(777), config.Storage.Devices[0].Partitions[0].StartMegaBytes)
	a.Equal(uint64(777), config.Storage.Devices[0].Partitions[0].SizeMegaBytes)
	a.Equal("type", config.Storage.Devices[0].Partitions[1].Type)
	a.Equal("mount", config.Storage.Devices[0].Partitions[1].Mount)
	a.Equal(true, config.Storage.Devices[0].Partitions[1].Boot)
	a.Equal(uint64(777), config.Storage.Devices[0].Partitions[1].StartMegaBytes)
	a.Equal(uint64(777), config.Storage.Devices[0].Partitions[1].SizeMegaBytes)

	a.Equal("device", config.Storage.Devices[1].Device)
	a.Equal("type", config.Storage.Devices[1].Partitions[0].Type)
	a.Equal("mount", config.Storage.Devices[1].Partitions[0].Mount)
	a.Equal(uint64(777), config.Storage.Devices[1].Partitions[0].StartMegaBytes)
	a.Equal(uint64(777), config.Storage.Devices[1].Partitions[0].SizeMegaBytes)

	a.Equal("mirror", config.Pacman.Mirror)
}

var validJson = `
{
  "version": "version",
  "keyboard": {
    "layout": "layout"
  },
  "localization": {
    "system_language": "system-language",
    "locales": [
		{
			"language": "language",
			"encoding": "encoding"
		},
		{
			"language": "language",
			"encoding": "encoding"
		}
	],
    "timezone": "timezone"
  },
  "computer": {
    "hostname": "hostname",
    "username": "username",
    "password": "password"
  },
  "storage": {
    "devices": [
      {
        "device": "device",
        "partitions": [
          {
            "type": "type",
			"boot": false,
            "mount": "mount",
            "start_mb": 777,
            "size_mb": 777
          },
          {
            "type": "type",
			"boot": true,
            "mount": "mount",
            "start_mb": 777,
            "size_mb": 777
          }
        ]
      },
      {
        "device": "device",
        "partitions": [
          {
            "type": "type",
            "mount": "mount",
            "start_mb": 777,
            "size_mb": 777
          }
        ]
      }
    ]
  },
  "pacman": {
	"mirror": "mirror"
  }
}
`
