#!/bin/bash
set -euo pipefail

# shellcheck disable=SC2312
ROOT_DIR="$(dirname "$(realpath "${BASH_SOURCE[0]}")")"

# @TODO: Configure IAM trust policies.
# https://github.com/awslabs/mcp/blob/main/src/aws-api-mcp-server/DEPLOYMENT.md

aws-vault --debug exec --duration=15m --ecs-server --region=us-east-2 --lazy devenv-base-role -- \
    aws-mcp-credential-proxy -- \
    docker mcp gateway run \
    --servers=amazon-bedrock-agentcore \
    --servers=aws-api \
    --servers=aws-core-mcp-server \
    --servers=aws-diagram \
    --servers=aws-documentation \
    --servers=aws-terraform \
    --servers=curl \
    --servers=docker \
    --servers=gitmcp \
    --servers=kagisearch \
    --servers=remote-mcp \
    --servers=terraform \
    --servers=time \
    --servers=vuln-nist-mcp-server \
    --tools=call_aws \
    --tools=convert_time \
    --tools=curl \
    --tools=cve_change_history \
    --tools=cves_by_cpe \
    --tools=docker \
    --tools=fetch_agentcore_doc \
    --tools=fetch_generic_documentation \
    --tools=fetch_generic_url_content \
    --tools=generate_diagram \
    --tools=get_current_time \
    --tools=get_cve_by_id \
    --tools=get_diagram_examples \
    --tools=get_latest_module_version \
    --tools=get_latest_provider_version \
    --tools=get_latest_provider_version \
    --tools=get_module_details \
    --tools=get_policy_details \
    --tools=get_provider_capabilities \
    --tools=get_provider_details \
    --tools=get_provider_details \
    --tools=get_temporal_context \
    --tools=kagi_search_fetch \
    --tools=kagi_summarizer \
    --tools=kevs_between \
    --tools=list_icons \
    --tools=manage_agentcore_gateway \
    --tools=manage_agentcore_memory \
    --tools=manage_agentcore_runtime \
    --tools=match_common_libs_owner_repo_mapping \
    --tools=mcp-add \
    --tools=mcp-create-profile \
    --tools=mcp-find \
    --tools=prompt_understanding \
    --tools=read_documentation \
    --tools=recommend \
    --tools=search_agentcore_docs \
    --tools=search_cloudflare_documentation \
    --tools=search_cves \
    --tools=search_documentation \
    --tools=search_generic_code \
    --tools=search_generic_documentation \
    --tools=search_modules \
    --tools=search_policies \
    --tools=search_providers \
    --tools=SearchAwsProviderDocs \
    --tools=SearchUserProvidedModule \
    --tools=suggest_aws_commands \
    --additional-config="${ROOT_DIR}/config.yml" \
    ;
