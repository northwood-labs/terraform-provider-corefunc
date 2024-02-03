#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "listing the data sources" {
    run tfschema data list corefunc

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == "corefunc_env_ensure" ]]
    [[ ${lines[1]} == "corefunc_homedir_expand" ]]
    [[ ${lines[2]} == "corefunc_homedir_get" ]]
    [[ ${lines[3]} == "corefunc_int_leftpad" ]]
    [[ ${lines[4]} == "corefunc_runtime_cpuarch" ]]
    [[ ${lines[5]} == "corefunc_runtime_goroot" ]]
    [[ ${lines[6]} == "corefunc_runtime_numcpus" ]]
    [[ ${lines[7]} == "corefunc_runtime_os" ]]
    [[ ${lines[8]} == "corefunc_str_camel" ]]
    [[ ${lines[9]} == "corefunc_str_constant" ]]
    [[ ${lines[10]} == "corefunc_str_iterative_replace" ]]
    [[ ${lines[11]} == "corefunc_str_kebab" ]]
    [[ ${lines[12]} == "corefunc_str_leftpad" ]]
    [[ ${lines[13]} == "corefunc_str_pascal" ]]
    [[ ${lines[14]} == "corefunc_str_snake" ]]
    [[ ${lines[15]} == "corefunc_str_truncate_label" ]]
    [[ ${lines[16]} == "corefunc_url_parse" ]]
}
