#!/usr/bin/env bash
set -uo pipefail

##
# All shell scripts should end with a .sh file extension.
##

declare -i RC=0

has_sh_extension() {
    [[ "$(basename "${1}")" != "$(basename "${1}" .sh)" ]]
}

for filename in "${@}"; do
    if ! has_sh_extension "${filename}"; then
        echo "[ERROR] ${filename} lacks the required .sh extension"
        RC=1
    fi
done

# shellcheck disable=2248
exit ${RC}
