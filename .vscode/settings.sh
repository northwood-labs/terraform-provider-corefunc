#!/usr/bin/env bash
set -euo pipefail

# @config-manager:start settings-merge
##
# Generate with:
#   bash settings.sh
##

vars=()
keys=()

# Resolve the directory where this script resides
SCRIPT_DIR="$(dirname "${BASH_SOURCE[0]}")"

files="$(find "${SCRIPT_DIR}" -type f -name "settings*.toml")"
sorted_files="$(echo "${files}" | sort)"

while IFS= read -r file; do
    name="${file#"${SCRIPT_DIR}/"}"                      # strip script directory prefix
    tname="${name#settings-}"                            # strip settings- prefix
    key="${tname%.toml}"                                 # strip .toml suffix
    key="$(echo "${key}" | tr -cd '[:alpha:][:digit:]')" # alphanumeric only
    vars+=("--var" "${key}=toml:file:${file}")
    keys+=("\$${key}")
done < <(echo "${sorted_files}")

merge="merge($(
    IFS=', '
    echo "${keys[*]}"
))"

# We want the spaces in the string represented by ${vars[*]}. This allows the
# output to be independent CLI flags.
# shellcheck disable=SC2145
dasel_command="dasel query --in toml --out json ${vars[*]} '${merge}'"

# Execute the command we constructed as a string. Pass through jq for
# formatting, and write output to sibling file.
eval "${dasel_command}" | jq -Mr --sort-keys '.' >"${SCRIPT_DIR}/settings.json"
# @config-manager:end settings-merge
