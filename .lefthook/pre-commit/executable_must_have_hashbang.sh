#!/bin/bash

##
# 1. Get a list of changed files in the repo.
# 2. Get a list of executable files in the repo.
# 3. Use `comm` to get a list of files that are in both sets.
# 4. Once we have the list of (a) executable, (b) changed files...
#    a. Look at their contents.
#    b. Check the first 2 bytes for a hashbang.
##

# shellcheck disable=SC2016,SC2312
comm -12 \
    <(git diff --name-only | sort) \
    <(find . -type f -executable -print0 | xargs -0 -I% bash -c 'echo ${1#./}' _ "%" \; | sort) |
    xargs -I% bash -c '
        contents="$(cat "%")"
        if [[ ${contents:0:2} != "#!" ]]; then
            echo "[ERROR] File % is marked executable, but does not have a hashbang."
            exit 1
        fi
    '
