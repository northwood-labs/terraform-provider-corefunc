#!/usr/bin/env bats
# https://bats-core.readthedocs.io/en/stable/writing-tests.html

@test "corefunc_url_parse: attrs" {
    run bash -c "tfschema data show -format=json corefunc_url_parse | jq -Mrc '.attributes[]'"

    [[ ${status} -eq 0 ]]
    [[ ${lines[0]} == '{"name":"decoded_port","type":"number","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[1]} == '{"name":"fragment","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[2]} == '{"name":"hash","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[3]} == '{"name":"host","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[4]} == '{"name":"hostname","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[5]} == '{"name":"normalized","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[6]} == '{"name":"normalized_nofrag","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[7]} == '{"name":"password","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[8]} == '{"name":"path","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[9]} == '{"name":"port","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[10]} == '{"name":"protocol","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[11]} == '{"name":"query","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[12]} == '{"name":"scheme","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[13]} == '{"name":"search","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
    [[ ${lines[14]} == '{"name":"url","type":"string","required":true,"optional":false,"computed":false,"sensitive":false}' ]]
    [[ ${lines[15]} == '{"name":"username","type":"string","required":false,"optional":false,"computed":true,"sensitive":false}' ]]
}
