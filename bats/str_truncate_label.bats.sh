#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_str_truncate_label: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_truncate_label | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '{"name":"label","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"max_length","type":"number","required":false,"optional":true,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"prefix","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[3]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
