#!/usr/bin/env bash
terraform providers lock \
    -platform=darwin_amd64 \
    -platform=darwin_arm64 \
    -platform=freebsd_386 \
    -platform=freebsd_amd64 \
    -platform=freebsd_arm \
    -platform=freebsd_arm64 \
    -platform=linux_386 \
    -platform=linux_amd64 \
    -platform=linux_arm \
    -platform=linux_arm64 \
    -platform=netbsd_386 \
    -platform=netbsd_amd64 \
    -platform=openbsd_386 \
    -platform=openbsd_amd64 \
    -platform=windows_386 \
    -platform=windows_amd64 \
    -platform=windows_arm \
    -platform=windows_arm64 \
    ;
