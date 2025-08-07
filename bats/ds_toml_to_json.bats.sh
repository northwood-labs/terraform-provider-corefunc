#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_toml_to_json: attrs" {
    run bash -c "tfschema data show -format=json corefunc_toml_to_json | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '{"name":"json","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"toml","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
}
