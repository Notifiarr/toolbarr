#!/usr/bin/env bash
# This file builds a standard DMG installer for macOS.
# This only works on macOS.
###########################################

set -e -o pipefail

# If we are running in GH Actions, make a new keychain and import the certificate.
if [ -n "$APPLE_SIGNING_KEY" ]; then
  KEYCHAIN="ios-build.keychain"

  echo "==> Creating new keychain: $KEYCHAIN"
  security create-keychain -p secret $KEYCHAIN

  echo "==> Importing certificate into ${KEYCHAIN}"
  echo "${APPLE_SIGNING_KEY}" | base64 -d | \
    security import /dev/stdin -P '' -f pkcs12 -k $KEYCHAIN -T /usr/bin/codesign

  echo "==> Unlocking keychain ${KEYCHAIN}"
  security unlock-keychain -p secret $KEYCHAIN

  echo "==> Increase keychain unlock timeout to 1 hour."
  security set-keychain-settings -lut 3600 $KEYCHAIN
  
  security set-key-partition-list -S apple-tool:,apple: -s -k secret $KEYCHAIN

  echo "==> Add keychain to keychain-list"
  security list-keychains -s $KEYCHAIN
fi

echo "==> Signing App."
gon build/darwin/sign.json

# Creating non-notarized DMG.
mkdir -p build/bin
hdiutil create build/bin/Toolbarr.dmg -srcfolder build/bin/Toolbarr.app -ov

echo "==> Notarizing DMG."
gon build/darwin/notarize.json

echo "==> Finished."
