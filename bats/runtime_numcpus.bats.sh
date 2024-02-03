#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_runtime_numcpus: attrs" {
    run bash -c "tfschema data show -format=json corefunc_runtime_numcpus | jq -Mrc '.attributes[]'"

    [[ "${status}" -eq 0 ]]
    [[ ${lines[0]} == '{"name":"value","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
