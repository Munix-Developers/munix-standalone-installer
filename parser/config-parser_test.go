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

	a.Equal("language", config.Localization.Language)
	a.Equal("locale", config.Localization.Locales[0])
	a.Equal("locale", config.Localization.Locales[1])
	a.Equal("timezone", config.Localization.Timezone)

	a.Equal("hostname", config.Computer.Hostname)
	a.Equal("username", config.Computer.Username)
	a.Equal("password", config.Computer.Password)

	a.Equal("device", config.Storage.Devices[0].Device)
	a.Equal("type", config.Storage.Devices[0].Partitions[0].Type)
	a.Equal("mount", config.Storage.Devices[0].Partitions[0].Mount)
	a.Equal(uint64(777), config.Storage.Devices[0].Partitions[0].Start)
	a.Equal(uint64(777), config.Storage.Devices[0].Partitions[0].End)
	a.Equal("type", config.Storage.Devices[0].Partitions[1].Type)
	a.Equal("mount", config.Storage.Devices[0].Partitions[1].Mount)
	a.Equal(uint64(777), config.Storage.Devices[0].Partitions[1].Start)
	a.Equal(uint64(777), config.Storage.Devices[0].Partitions[1].End)

	a.Equal("device", config.Storage.Devices[1].Device)
	a.Equal("type", config.Storage.Devices[1].Partitions[0].Type)
	a.Equal("mount", config.Storage.Devices[1].Partitions[0].Mount)
	a.Equal(uint64(777), config.Storage.Devices[1].Partitions[0].Start)
	a.Equal(uint64(777), config.Storage.Devices[1].Partitions[0].End)
}

var validJson = `
{
  "version": "version",
  "keyboard": {
    "layout": "layout"
  },
  "localization": {
    "language": "language",
    "locales": [
      "locale",
      "locale"
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
            "mount": "mount",
            "start": 777,
            "end": 777
          },
          {
            "type": "type",
            "mount": "mount",
            "start": 777,
            "end": 777
          }
        ]
      },
      {
        "device": "device",
        "partitions": [
          {
            "type": "type",
            "mount": "mount",
            "start": 777,
            "end": 777
          }
        ]
      }
    ]
  }
}
`
