#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "listing the data sources" {
    run tfschema data list corefunc

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == "corefunc_env_ensure" ]]
    [[ ${lines[1]} == "corefunc_int_leftpad" ]]
    [[ ${lines[2]} == "corefunc_str_camel" ]]
    [[ ${lines[3]} == "corefunc_str_constant" ]]
    [[ ${lines[4]} == "corefunc_str_iterative_replace" ]]
    [[ ${lines[5]} == "corefunc_str_kebab" ]]
    [[ ${lines[6]} == "corefunc_str_leftpad" ]]
    [[ ${lines[7]} == "corefunc_str_pascal" ]]
    [[ ${lines[8]} == "corefunc_str_snake" ]]
    [[ ${lines[9]} == "corefunc_str_truncate_label" ]]
}
