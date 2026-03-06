#!/usr/bin/env bash

function handler () {
  set -eu
  echo "Received event, parsing..." > /dev/stderr
  FORMAT="$(echo "$1" | jq -r '.Parameters.Format')"
  CONFIG="$(echo "$1" | jq -r '.Content' | base64 -d)"

  echo "SOPS decrypting configuration of type ${FORMAT}..." > /dev/stderr
  DECRYPTED="$(echo "${CONFIG}" | sops --input-type "${FORMAT}" --output-type "${FORMAT}" --decrypt /dev/stdin | base64 -w 0)"
  echo "Decryption Complete" > /dev/stderr

  echo "{ \"Content\": \"${DECRYPTED}\" }"
}