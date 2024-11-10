#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "listing the data sources" {
    run tfschema data list corefunc

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == "corefunc_env_ensure" ]]
    [[ ${lines[1]} == "corefunc_homedir_expand" ]]
    [[ ${lines[2]} == "corefunc_homedir_get" ]]
    [[ ${lines[3]} == "corefunc_int_leftpad" ]]
    [[ ${lines[4]} == "corefunc_net_cidr_contains" ]]
    [[ ${lines[5]} == "corefunc_runtime_cpuarch" ]]
    [[ ${lines[6]} == "corefunc_runtime_goroot" ]]
    [[ ${lines[7]} == "corefunc_runtime_numcpus" ]]
    [[ ${lines[8]} == "corefunc_runtime_os" ]]
    [[ ${lines[9]} == "corefunc_str_camel" ]]
    [[ ${lines[10]} == "corefunc_str_constant" ]]
    [[ ${lines[11]} == "corefunc_str_iterative_replace" ]]
    [[ ${lines[12]} == "corefunc_str_kebab" ]]
    [[ ${lines[13]} == "corefunc_str_leftpad" ]]
    [[ ${lines[14]} == "corefunc_str_pascal" ]]
    [[ ${lines[15]} == "corefunc_str_snake" ]]
    [[ ${lines[16]} == "corefunc_str_truncate_label" ]]
    [[ ${lines[17]} == "corefunc_url_decode" ]]
    [[ ${lines[18]} == "corefunc_url_parse" ]]
}
