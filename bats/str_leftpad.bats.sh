#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_str_leftpad: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_leftpad | jq -Mrc '.attributes[]'"

    [[ "${status}" -eq 0 ]]
    [[ ${lines[0]} == '{"name":"pad_char","type":"string","required":false,"optional":true,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"pad_width","type":"number","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"string","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[3]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
