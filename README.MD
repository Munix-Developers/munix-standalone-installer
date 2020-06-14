# Munix Installer

> **XXX > !WARNING! < XXX** DON'T RUN THE `muinstaller` IN A ALREADY INSTALLED LINUX SYSTEM! THE EXECUTABLE MAY CAUSE **PERMANET DATA LOSS**. **PROCEED WITH CAUTION**.

The Munix installer is a command line tool to install the Munix Linux distribution.

## Arguments

* `config.json` path

## `config.json` structure

```json
{
  "version": "1.0",
  "keyboard": {
    "layout": "br-abnt2"
  },
  "localization": {
    "language": "en_US.UTF-8",
    "locales": [
      "en_US.UTF-8 UTF-8",
      "pt_BR.UTF-8 UTF-8"
    ],
    "timezone": "America/Sao_Paulo"
  },
  "computer": {
    "hostname": "crazywriter",
    "username": "mat",
    "password": "$6$?qdYwnMnJpf'wYG?$zjUsHpBTcDsbnYcMuCiUB0YuAwYUhL07A24x1k.Skwc5xnCOCzTnSS8xTRe4ABgCKbzuud1Pi7KNhLK0mJkPu/"
  },
  "storage": {
    "devices": [
      {
        "device": "/dev/sda",
        "partitions": [
          {
            "type": "ext4",
            "mount": "/boot",
            "start": 2048,
            "end": 1126399
          },
          {
            "type": "ext4",
            "mount": "/home",
            "start": 1126400,
            "end": 1952462941
          }
        ]
      },
      {
        "device": "/dev/sdb",
        "partitions": [
          {
            "type": "ext4",
            "mount": "/",
            "start": 2048,
            "end": 1952462941
          }
        ]
      }
    ]
  }
}
```
