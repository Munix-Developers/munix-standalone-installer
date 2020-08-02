#!/usr/bin/env bash

# Usage   host-webserver.sh HOST-IP
# Use this script to open a HTTP server in 7777 port.
# It is used for the VirtualBox arch machine download the installer while running.
# Virtualbox machine MUST be in bridge network mode.

python -m http.server --bind "$1" 7777
