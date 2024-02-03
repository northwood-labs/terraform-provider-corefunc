#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_env_ensure: attrs" {
    run bash -c "tfschema data show -format=json corefunc_env_ensure | jq -Mrc '.attributes[]'"

    [[ "${status}" -eq 0 ]]
    [[ ${lines[0]} == '{"name":"name","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"pattern","type":"string","required":false,"optional":true,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
