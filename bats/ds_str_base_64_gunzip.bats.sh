#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_str_base64_gunzip: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_base64_gunzip | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '{"name":"gzipped_base64","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
