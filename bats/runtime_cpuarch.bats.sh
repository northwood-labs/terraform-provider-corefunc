#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_runtime_cpuarch: attrs" {
    run bash -c "tfschema data show -format=json corefunc_runtime_cpuarch | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
