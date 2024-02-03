#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_str_iterative_replace: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_iterative_replace | jq -Mrc '.attributes[]'"

    [[ "${status}" -eq 0 ]]
    [[ ${lines[0]} == '{"name":"replacements","type":"list(object({new=string,old=string}))","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"string","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
