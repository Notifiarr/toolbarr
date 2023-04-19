#!/usr/bin/env bash

# Builds the translation json files and renames them into place.
# Add a new language by creating a folder for it: mkdir locales/xx

RUN="golang.org/x/text/cmd/gotext@latest"
LANG=$(ls locales|xargs|tr ' ' ,)
PKGS="github.com/Notifiarr/toolbarr/pkg/..."

# go run ${RUN} -srclang=en update -out=catalog.go -lang=${LANG} ${PKGS} && \
#     mv locales/en/out.gotext.json locales/en/messages.gotext.json
