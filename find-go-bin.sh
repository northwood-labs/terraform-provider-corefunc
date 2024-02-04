#!/usr/bin/env bash

if [[ -n ${GOBIN} ]]; then
    echo "${GOBIN}"
elif [[ -n ${GOPATH} ]]; then
    echo "${GOPATH}/bin"
else
    echo "${HOME}/go/bin"
fi
