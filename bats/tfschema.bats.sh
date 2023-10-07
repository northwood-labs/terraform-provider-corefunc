#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "listing the data sources" {
    run tfschema data list corefunc

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == "corefunc_env_ensure" ]]
    [[ ${lines[1]} == "corefunc_str_camel" ]]
    [[ ${lines[2]} == "corefunc_str_constant" ]]
    [[ ${lines[3]} == "corefunc_str_kebab" ]]
    [[ ${lines[4]} == "corefunc_str_pascal" ]]
    [[ ${lines[5]} == "corefunc_str_snake" ]]
    [[ ${lines[6]} == "corefunc_str_truncate_label" ]]

}

@test "corefunc_env_ensure: attrs" {
    run bash -c "tfschema data show -format=json corefunc_env_ensure | jq -Mrc '.attributes[]'"

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == '{"name":"id","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"name","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"pattern","type":"string","required":false,"optional":true,"computed":false,"sensitive":false}' ]]
    [[ ${lines[3]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}

@test "corefunc_str_camel: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_camel | jq -Mrc '.attributes[]'"

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == '{"name":"id","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"string","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}

@test "corefunc_str_constant: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_constant | jq -Mrc '.attributes[]'"

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == '{"name":"id","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"string","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}

@test "corefunc_str_kebab: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_kebab | jq -Mrc '.attributes[]'"

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == '{"name":"id","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"string","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}

@test "corefunc_str_pascal: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_pascal | jq -Mrc '.attributes[]'"

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == '{"name":"acronym_caps","type":"bool","required":false,"optional":true,"computed":false,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"id","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"string","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[3]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}

@test "corefunc_str_snake: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_snake | jq -Mrc '.attributes[]'"

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == '{"name":"id","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"string","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}

@test "corefunc_str_truncate_label: attrs" {
    run bash -c "tfschema data show -format=json corefunc_str_truncate_label | jq -Mrc '.attributes[]'"

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == '{"name":"id","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"label","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"max_length","type":"number","required":false,"optional":true,"computed":false,"sensitive":false}' ]]
    [[ ${lines[3]} == '{"name":"prefix","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[4]} == '{"name":"value","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
