#!/bin/bash

# Used for fast testing inside VirtualBox
# Run host-webserver.sh first.

# Run this command inside the machine:
# export MUNIX=HOST_WEBSERVER_IP:7777 && curl $MUNIX/virtualbox.sh | sh -

curl $MUNIX/valid.json > valid.json && \
curl $MUNIX/muinstaller > muinstaller && \
chmod +x muinstaller && \
./muinstaller valid.json
