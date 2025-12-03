#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_time_parse: attrs" {
    run bash -c "tfschema data show -format=json corefunc_time_parse | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '{"name":"input","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"value","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
