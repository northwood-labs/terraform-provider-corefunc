#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_map_key_if: attrs" {
    run bash -c "tfschema data show -format=json corefunc_map_key_if | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '@TODO' ]]
}
