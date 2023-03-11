#!/usr/bin/env bash

set -e -o pipefail

# https://blog.synapp.nz/2017/06/16/code-signing-a-windows-application-on-linux-on-windows/
# https://stackoverflow.com/a/29073957/1142

function sign() {
  if [ -z "${EXE_SIGNING_KEY}" ]; then
    echo "Skipped ${FILE} .."
    exit
  fi

  echo "${EXE_SIGNING_KEY}" | base64 -d | \
  osslsigncode sign -pkcs12 /dev/stdin \
    -pass "${EXE_SIGNING_KEY_PASSWORD}" \
    -n "Toolbarr" \
    -i "https://www.golift.io" \
    -t "http://timestamp.comodoca.com/authenticode" \
    -in "${FILE}" -out "signed.${FILE}"

  mv "signed.${FILE}" "${FILE}"
  echo "Signed ${FILE} .."
}

[ -z "$1" ] || FILE=$1 sign

