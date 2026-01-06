#!/bin/bash
set -euo pipefail

##
# View installed servers.
#   docker mcp server ls
#
# View installed tools.
#   docker mcp tools ls
#
# Helpful tools for learning more.
#   docker mcp tools inspect <tool-name>
#   docker mcp tools call <tool-name>
##

# shellcheck disable=SC2312
ROOT_DIR="$(dirname "$(realpath "${BASH_SOURCE[0]}")")"

# Install into Kiro
docker mcp client connect kiro

# Set Kagi search API config
op read --no-newline "op://Private/ibzwf3qxqukhwvi4mdrnxozozq/credential" | docker mcp secret set kagisearch.api_key

# Symlink
ln -s "${ROOT_DIR}/configure-aws-mcp.sh" /usr/local/bin/aws-mcp
