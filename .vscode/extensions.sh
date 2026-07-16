#!/usr/bin/env bash
set -euo pipefail

# @config-manager:start extensions-merge
##
# Generate with:
#   bash extensions.sh
##

vars=()
keys=()

# Resolve the directory where this script resides
SCRIPT_DIR="$(dirname "${BASH_SOURCE[0]}")"

files="$(find "${SCRIPT_DIR}" -type f -name "extensions*.toml")"
sorted_files="$(echo "${files}" | sort)"

while IFS= read -r file; do
    name="${file#"${SCRIPT_DIR}/"}"                      # strip script directory prefix
    tname="${name#extensions-}"                          # strip extensions- prefix
    key="${tname%.toml}"                                 # strip .toml suffix
    key="$(echo "${key}" | tr -cd '[:alpha:][:digit:]')" # alphanumeric only
    vars+=("--var" "${key}=toml:file:${file}")
    keys+=("\$${key}.recommendations...")
done < <(echo "${sorted_files}")

merge="[$(
    IFS=', '
    echo "${keys[*]}"
)]"

# We want the spaces in the string represented by ${vars[*]}. This allows the
# output to be independent CLI flags.
# shellcheck disable=SC2145
dasel_command="dasel query --in toml --out json ${vars[*]} '${merge}'"

# Execute the command we constructed as a string. Pass through jq for
# formatting, and write output to sibling file.
eval "${dasel_command}" | jq -Mr '. | sort_by(. | ascii_downcase) | { recommendations: . }' \
    >"${SCRIPT_DIR}/extensions.json"
# @config-manager:end extensions-merge
