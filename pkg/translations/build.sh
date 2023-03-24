#!/usr/bin/env bash

# Builds the translation json files and renames them into place.
# Add a new language by creating a folder for it: mkdir locales/xx

PKG="golang.org/x/text/cmd/gotext@latest"
LANG=$(ls locales|xargs|tr ' ' ,)

go run ${PKG} -srclang=en update \
    -out=catalog.go \
    -lang=${LANG} \
    github.com/Notifiarr/toolbarr/pkg/... && \
    for file in $(ls locales/*/out.gotext.json); do 
        mv "${file}" "$(dirname "${file}")/messages.gotext.json"
    done