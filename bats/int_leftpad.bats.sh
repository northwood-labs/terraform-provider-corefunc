#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_int_leftpad: attrs" {
    run bash -c "tfschema data show -format=json corefunc_int_leftpad | jq -Mrc '.attributes[]'"

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == '{"name":"id","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"num","type":"number","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"pad_width","type":"number","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[3]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
