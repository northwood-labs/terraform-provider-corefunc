#!/bin/bash
# shellcheck disable=2312
gommit check message "$(cat "$1")"
