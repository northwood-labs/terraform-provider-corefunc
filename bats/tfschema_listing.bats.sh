#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "listing the data sources" {
    run tfschema data list corefunc

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == "corefunc_env_ensure" ]]
    [[ ${lines[1]} == "corefunc_hash_md5" ]]
    [[ ${lines[2]} == "corefunc_hash_md5_base64" ]]
    [[ ${lines[3]} == "corefunc_homedir_expand" ]]
    [[ ${lines[4]} == "corefunc_homedir_get" ]]
    [[ ${lines[5]} == "corefunc_int_leftpad" ]]
    [[ ${lines[6]} == "corefunc_json_to_toml" ]]
    [[ ${lines[7]} == "corefunc_net_cidr_contains" ]]
    [[ ${lines[8]} == "corefunc_runtime_cpuarch" ]]
    [[ ${lines[9]} == "corefunc_runtime_numcpus" ]]
    [[ ${lines[10]} == "corefunc_runtime_os" ]]
    [[ ${lines[11]} == "corefunc_str_base64_gunzip" ]]
    [[ ${lines[12]} == "corefunc_str_camel" ]]
    [[ ${lines[13]} == "corefunc_str_constant" ]]
    [[ ${lines[14]} == "corefunc_str_iterative_replace" ]]
    [[ ${lines[15]} == "corefunc_str_kebab" ]]
    [[ ${lines[16]} == "corefunc_str_leftpad" ]]
    [[ ${lines[17]} == "corefunc_str_pascal" ]]
    [[ ${lines[18]} == "corefunc_str_snake" ]]
    [[ ${lines[19]} == "corefunc_toml_to_json" ]]
    [[ ${lines[20]} == "corefunc_url_decode" ]]
    [[ ${lines[21]} == "corefunc_url_parse" ]]
}
