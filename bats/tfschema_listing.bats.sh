#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "listing the data sources" {
    run tfschema data list corefunc

    [ "$status" -eq 0 ]
    [[ ${lines[0]} == "corefunc_env_ensure" ]]
    [[ ${lines[1]} == "corefunc_homedir_expand" ]]
    [[ ${lines[2]} == "corefunc_homedir_get" ]]
    [[ ${lines[3]} == "corefunc_int_leftpad" ]]
    [[ ${lines[4]} == "corefunc_str_camel" ]]
    [[ ${lines[5]} == "corefunc_str_constant" ]]
    [[ ${lines[6]} == "corefunc_str_iterative_replace" ]]
    [[ ${lines[7]} == "corefunc_str_kebab" ]]
    [[ ${lines[8]} == "corefunc_str_leftpad" ]]
    [[ ${lines[9]} == "corefunc_str_pascal" ]]
    [[ ${lines[10]} == "corefunc_str_snake" ]]
    [[ ${lines[11]} == "corefunc_str_truncate_label" ]]
}
