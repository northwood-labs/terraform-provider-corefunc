#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_str_iterative_replace: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_iterative_replace | jq -Mrc '.attributes[]'"

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == '@TODO' ]]
}
