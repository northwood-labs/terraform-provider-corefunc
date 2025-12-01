#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_hash_bcrypt: attrs" {
    run bash -c "tfschema data show -format=json corefunc_hash_bcrypt | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '{"name":"cost","type":"number","required":false,"optional":true,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"input","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
