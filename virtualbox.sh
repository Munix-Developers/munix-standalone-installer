#!/bin/bash

# Assuming export MUNIX=192.168.1.107:7777 && curl $MUNIX/virtualbox.sh | sh -

curl $MUNIX/valid.json > valid.json && \
curl $MUNIX/muinstaller > muinstaller && \
chmod +x muinstaller && \
./muinstaller valid.json
