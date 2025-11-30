#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_hash_scrypt_base64: attrs" {
    run bash -c "tfschema data show -format=json corefunc_hash_scrypt_base64 | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '{"name":"input","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"salt","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
