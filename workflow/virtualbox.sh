#!/bin/bash

# Used for fast testing inside VirtualBox
# Run host-webserver.sh first.

# Run this command inside the machine:
# export M=HOST_IP && curl $M:7777/v | sh -

HOST_MUNIX_IP=$M
HOST_MUNIX_PORT=7777
HOST_MUNIX_SQUID_PORT=8080
export all_proxy="$HOST_MUNIX_IP:$HOST_MUNIX_SQUID_PORT"

curl "http://$HOST_MUNIX_IP:$HOST_MUNIX_PORT"/valid.json >valid.json &&
  curl "http://$HOST_MUNIX_IP:$HOST_MUNIX_PORT"/muinstaller >muinstaller &&
  chmod +x muinstaller &&
  ./muinstaller valid.json
