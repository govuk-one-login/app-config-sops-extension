#!/usr/bin/env bash

source "function.sh"
CONTENT="$(cat "$1" | base64)"
EVENT_DATA="$(cat <<EOF
{
  "Parameters": {
    "Format": "json"
  },
  "Content": "${CONTENT}"
}
EOF
)"
handler "${EVENT_DATA}"