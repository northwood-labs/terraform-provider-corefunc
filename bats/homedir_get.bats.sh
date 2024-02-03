#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_homedir_get: attrs" {
    run bash -c "tfschema data show -format=json corefunc_homedir_get | jq -Mrc '.attributes[]'"

    [[ "${status}" -eq 0 ]]
    [[ ${lines[0]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
