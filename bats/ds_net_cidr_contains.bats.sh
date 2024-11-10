#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_net_cidr_contains: attrs" {
    run bash -c "tfschema data show -format=json corefunc_net_cidr_contains | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '{"name":"contained_ip_or_cidr","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"containing_cidr","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"value","type":"bool","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
