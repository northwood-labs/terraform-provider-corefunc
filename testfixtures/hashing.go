// Copyright 2024-2025, Northwood Labs, LLC <license@northwood-labs.com>
// Copyright 2023-2025, Ryan Parman <rparman@northwood-labs.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testfixtures // lint:no_dupe

var (
	// HashMD5TestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	HashMD5TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "d41d8cd98f00b204e9800998ecf8427e",
		},
		"a": {
			Input:    "a",
			Expected: "0cc175b9c0f1b6a831c399e269772661",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "5eb63bbbe01eeed093cb22bb8f5acdc3",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "6f5902ac237024bdd0c176cb93063dc4",
		},
	}

	// HashSHA1TestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	HashSHA1TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "da39a3ee5e6b4b0d3255bfef95601890afd80709",
		},
		"a": {
			Input:    "a",
			Expected: "86f7e437faa5a7fce15d1ddcb9eaeaea377667b8",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "2aae6c35c94fcfb415dbe95f408b9ce91ee846ed",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "22596363b3de40b06f981fb85d82312e8c0ed511",
		},
	}

	// HashSHA256TestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	HashSHA256TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855",
		},
		"a": {
			Input:    "a",
			Expected: "ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447",
		},
	}

	// HashSHA384TestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	HashSHA384TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input: "",
			Expected: "38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f6" +
				"5fbd51ad2f14898b95b",
		},
		"a": {
			Input: "a",
			Expected: "54a59b9f22b0b80880d8427e548b7c23abd873486e1f035dce9cd697e85175033caa88e6d57bc" +
				"35efae0b5afd3145f31",
		},
		"hello world": {
			Input: "hello world",
			Expected: "fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646" +
				"efd0819dd8c088de1bd",
		},
		"hello world\n": {
			Input: "hello world\n",
			Expected: "6b3b69ff0a404f28d75e98a066d3fc64fffd9940870cc68bece28545b9a75086b343d7a136683" +
				"8083e4b8f3ca6fd3c80",
		},
	}

	// HashSHA512TestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	HashSHA512TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input: "",
			Expected: "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f" +
				"2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e",
		},
		"a": {
			Input: "a",
			Expected: "1f40fc92da241694750979ee6cf582f2d5d7d28e18335de05abc54d0560e0f5302860c652bf08" +
				"d560252aa5e74210546f369fbbbce8c12cfc7957b2652fe9a75",
		},
		"hello world": {
			Input: "hello world",
			Expected: "309ecc489c12d6eb4cc40f50c902f2b4d0ed77ee511a7c7a9bcd3ca86d4cd86f989dd35bc5ff4" +
				"99670da34255b45b0cfd830e81f605dcf7dc5542e93ae9cd76f",
		},
		"hello world\n": {
			Input: "hello world\n",
			Expected: "db3974a97f2407b7cae1ae637c0030687a11913274d578492558e39c16c017de84eacdc8c62fe" +
				"34ee4e12b4b1428817f09b6a2760c3f8a664ceae94d2434a593",
		},
	}

	// Base64HashMD5TestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <https://github.com/golang/go/wiki/TableDrivenTests>
	Base64HashMD5TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "1B2M2Y8AsgTpgAmY7PhCfg==",
		},
		"a": {
			Input:    "a",
			Expected: "DMF1ucDxtqgxw5niaXcmYQ==",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "XrY7u+Ae7tCTyyK7j1rNww==",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "b1kCrCNwJL3QwXbLkwY9xA==",
		},
	}

	// Base64HashSHA1TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashSHA1TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "2jmj7l5rSw0yVb/vlWAYkK/YBwk=",
		},
		"a": {
			Input:    "a",
			Expected: "hvfkN/qlp/zhXR3cuerq6jd2Z7g=",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "Kq5sNclPz7QV2+lfQIuc6R7oRu0=",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "IlljY7PeQLBvmB+4XYIxLowO1RE=",
		},
	}

	// Base64HashSHA256TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashSHA256TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
		},
		"a": {
			Input:    "a",
			Expected: "ypeBEsobvcr6wjGzmiPcTaeG7/gUfE5yuYB3ha/uSLs=",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "qUiQTy8PR5uPgZdpSzAYSw0u0cHNKh7A+4XSmaGSpEc=",
		},
	}

	// Base64HashSHA384TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashSHA384TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "OLBgp1GsljhM2TJ+sbHjaiH9txEUvgdDTAzHv2P24donTt6/529l+9Ua0vFImLlb",
		},
		"a": {
			Input:    "a",
			Expected: "VKWbnyKwuAiA2EJ+VIt8I6vYc0huHwNdzpzWl+hRdQM8qojm1XvDXvrgta/TFF8x",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "/b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "aztp/wpATyjXXpigZtP8ZP/9mUCHDMaL7OKFRbmnUIazQ9ehNmg4CD5Ljzym/TyA",
		},
	}

	// Base64HashSHA512TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashSHA512TestTable = map[string]struct { // lint:no_dupe
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==",
		},
		"a": {
			Input:    "a",
			Expected: "H0D8ktokFpR1CXnubPWC8tXX0o4YM13gWrxU0FYOD1MChgxlK/CNVgJSql50IQVG82n7u86MEs/HlXsmUv6adQ==",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "2zl0qX8kB7fK4a5jfAAwaHoRkTJ01XhJJVjjnBbAF96E6s3Ixi/jTuThK0sUKIF/Cbaidgw/imZM6ulNJDSlkw==",
		},
	}

	// HashSHA3x256TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	HashSHA3x256TestTable = map[string]struct {
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "a7ffc6f8bf1ed76651c14756a061d662f580ff4de43b49fa82d80a4b80f8434a",
		},
		"a": {
			Input:    "a",
			Expected: "80084bf2fba02475726feb2cab2d8215eab14bc6bdd8bfb2c8151257032ecd8b",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "644bcc7e564373040999aac89e7622f3ca71fba1d972fd94a31c3bfbf24e3938",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "a8009a7a528d87778c356da3a55d964719e818666a04e4f960c9e2439e35f138",
		},
	}

	// Base64HashSHA3x256TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashSHA3x256TestTable = map[string]struct {
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "p//G+L8e12ZRwUdWoGHWYvWA/03kO0n6gtgKS4D4Q0o=",
		},
		"a": {
			Input:    "a",
			Expected: "gAhL8vugJHVyb+ssqy2CFeqxS8a92L+yyBUSVwMuzYs=",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "ZEvMflZDcwQJmarInnYi88px+6HZcv2Uoxw7+/JOOTg=",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "qACaelKNh3eMNW2jpV2WRxnoGGZqBOT5YMniQ5418Tg=",
		},
	}

	// HashSHA3x384TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	HashSHA3x384TestTable = map[string]struct {
		Input    string
		Expected string
	}{
		"blank": {
			Input: "",
			Expected: "0c63a75b845e4f7d01107d852e4c2485c51a50aaaa94fc61995e71bbee983a2ac3713831264ad" +
				"b47fb6bd1e058d5f004",
		},
		"a": {
			Input: "a",
			Expected: "1815f774f320491b48569efec794d249eeb59aae46d22bf77dafe25c5edc28d7ea44f93ee1234" +
				"aa88f61c91912a4ccd9",
		},
		"hello world": {
			Input: "hello world",
			Expected: "83bff28dde1b1bf5810071c6643c08e5b05bdb836effd70b403ea8ea0a634dc4997eb1053aa35" +
				"93f590f9c63630dd90b",
		},
		"hello world\n": {
			Input: "hello world\n",
			Expected: "28fc308d4d5c1ef9e60acedb13c3a1fcf7266560602c639000580ae3541dea5ce78a685de897e" +
				"96b65a0fc15515c3780",
		},
	}

	// Base64HashSHA3x384TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashSHA3x384TestTable = map[string]struct {
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "DGOnW4ReT30BEH2FLkwkhcUaUKqqlPxhmV5xu+6YOirDcTgxJkrbR/tr0eBY1fAE",
		},
		"a": {
			Input:    "a",
			Expected: "GBX3dPMgSRtIVp7+x5TSSe61mq5G0iv3fa/iXF7cKNfqRPk+4SNKqI9hyRkSpMzZ",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "g7/yjd4bG/WBAHHGZDwI5bBb24Nu/9cLQD6o6gpjTcSZfrEFOqNZP1kPnGNjDdkL",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "KPwwjU1cHvnmCs7bE8Oh/PcmZWBgLGOQAFgK41Qd6lznimhd6Jfpa2Wg/BVRXDeA",
		},
	}

	// HashSHA3x512TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	HashSHA3x512TestTable = map[string]struct {
		Input    string
		Expected string
	}{
		"blank": {
			Input: "",
			Expected: "a69f73cca23a9ac5c8b567dc185a756e97c982164fe25859e0d1dcc1475c80a615b2123af1f5f" +
				"94c11e3e9402c3ac558f500199d95b6d3e301758586281dcd26",
		},
		"a": {
			Input: "a",
			Expected: "697f2d856172cb8309d6b8b97dac4de344b549d4dee61edfb4962d8698b7fa803f4f93ff24393" +
				"586e28b5b957ac3d1d369420ce53332712f997bd336d09ab02a",
		},
		"hello world": {
			Input: "hello world",
			Expected: "840006653e9ac9e95117a15c915caab81662918e925de9e004f774ff82d7079a40d4d27b1b372" +
				"657c61d46d470304c88c788b3a4527ad074d1dccbee5dbaa99a",
		},
		"hello world\n": {
			Input: "hello world\n",
			Expected: "4a936cbc1db296bd08d1c0bbf5a66a1897f35ee6d93047e0edff893dfbcba02f1e1570e85d118" +
				"7ea26bea6d54199e0656f1b7c21b9cc2102b8ed2a12769f4531",
		},
	}

	// Base64HashSHA3x512TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashSHA3x512TestTable = map[string]struct {
		Input    string
		Expected string
	}{
		"blank": {
			Input:    "",
			Expected: "pp9zzKI6msXItWfcGFp1bpfJghZP4lhZ4NHcwUdcgKYVshI68fX5TBHj6UAsOsVY9QAZnZW20+MBdYWGKB3NJg==",
		},
		"a": {
			Input:    "a",
			Expected: "aX8thWFyy4MJ1ri5faxN40S1SdTe5h7ftJYthpi3+oA/T5P/JDk1huKLW5V6w9HTaUIM5TMycS+Ze9M20JqwKg==",
		},
		"hello world": {
			Input:    "hello world",
			Expected: "hAAGZT6ayelRF6FckVyquBZikY6SXengBPd0/4LXB5pA1NJ7GzcmV8YdRtRwMEyIx4izpFJ60HTR3MvuXbqpmg==",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Expected: "SpNsvB2ylr0I0cC79aZqGJfzXubZMEfg7f+JPfvLoC8eFXDoXRGH6ia+ptVBmeBlbxt8IbnMIQK47SoSdp9FMQ==",
		},
	}

	// HashArgon2idTestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	HashArgon2idTestTable = map[string]struct {
		Input    string
		Expected string
		Salt     []byte
	}{
		"blank": {
			Input:    "",
			Salt:     []byte("somesaltvalue"),
			Expected: "c16ec4448c31255286143861f549e58863f61813be094f711f5f6f0b26f5571d",
		},
		"a": {
			Input:    "a",
			Salt:     []byte("somesaltvalue"),
			Expected: "45d500c86de20eabae294edf4647f6b1138be8188f1a3a42adb074b28f7ec258",
		},
		"hello world": {
			Input:    "hello world",
			Salt:     []byte("somesaltvalue"),
			Expected: "660b8e259df3496e9c429bebd97eb5d58a88cd46851e30c1dd03b9448c066ea4",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Salt:     []byte("somesaltvalue"),
			Expected: "968c2f9ee2822916836372525249828c9b2f7ae22cfc207c94bbb6242a1d54b1",
		},
	}

	// Base64HashArgon2idTestTable is used by both the standard Go tests and
	// also the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashArgon2idTestTable = map[string]struct {
		Input    string
		Expected string
		Salt     []byte
	}{
		"blank": {
			Input:    "",
			Salt:     []byte("somesaltvalue"),
			Expected: "wW7ERIwxJVKGFDhh9UnliGP2GBO+CU9xH19vCyb1Vx0=",
		},
		"a": {
			Input:    "a",
			Salt:     []byte("somesaltvalue"),
			Expected: "RdUAyG3iDquuKU7fRkf2sROL6BiPGjpCrbB0so9+wlg=",
		},
		"hello world": {
			Input:    "hello world",
			Salt:     []byte("somesaltvalue"),
			Expected: "ZguOJZ3zSW6cQpvr2X611YqIzUaFHjDB3QO5RIwGbqQ=",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Salt:     []byte("somesaltvalue"),
			Expected: "lowvnuKCKRaDY3JSUkmCjJsveuIs/CB8lLu2JCodVLE=",
		},
	}

	// HashScryptTestTable is used by both the standard Go tests and also the
	// Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	HashScryptTestTable = map[string]struct {
		Input    string
		Expected string
		Salt     []byte
	}{
		"blank": {
			Input:    "",
			Salt:     []byte("somesaltvalue"),
			Expected: "ba11348e553f3b5a39ae7e76aa48139b56f66646ed2b46bc185db2e70c8f0435",
		},
		"a": {
			Input:    "a",
			Salt:     []byte("somesaltvalue"),
			Expected: "512eb34c1dcdc1e22cf97705760aa6d528cc95e36188430fb15e86905a026507",
		},
		"hello world": {
			Input:    "hello world",
			Salt:     []byte("somesaltvalue"),
			Expected: "c58e3444ea05044db9cdb1406ac143884b351101c3e23c95e010ec994c8d65df",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Salt:     []byte("somesaltvalue"),
			Expected: "eba26b2a22481f6069a69c0aeacb0039a4fb47e27cb90effd774c21324735de5",
		},
	}

	// Base64HashScryptTestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashScryptTestTable = map[string]struct {
		Input    string
		Expected string
		Salt     []byte
	}{
		"blank": {
			Input:    "",
			Salt:     []byte("somesaltvalue"),
			Expected: "uhE0jlU/O1o5rn52qkgTm1b2ZkbtK0a8GF2y5wyPBDU=",
		},
		"a": {
			Input:    "a",
			Salt:     []byte("somesaltvalue"),
			Expected: "US6zTB3NweIs+XcFdgqm1SjMleNhiEMPsV6GkFoCZQc=",
		},
		"hello world": {
			Input:    "hello world",
			Salt:     []byte("somesaltvalue"),
			Expected: "xY40ROoFBE25zbFAasFDiEs1EQHD4jyV4BDsmUyNZd8=",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Salt:     []byte("somesaltvalue"),
			Expected: "66JrKiJIH2BpppwK6ssAOaT7R+J8uQ7/13TCEyRzXeU=",
		},
	}

	// HashHMACSHA256TestTable is used by both the standard Go tests and also
	// the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	HashHMACSHA256TestTable = map[string]struct {
		Input    string
		Key      string
		Expected string
	}{
		"blank": {
			Input:    "",
			Key:      "secretkey",
			Expected: "0ec2fbe02ea7c3eb6dd73c12eb2cffc9061280dfc8365cdcfa5241c6e3d9c9a7",
		},
		"a": {
			Input:    "a",
			Key:      "secretkey",
			Expected: "f290a3fdc66295788929ae4df26563c9c23bb7c94f93bd97413f339c35162928",
		},
		"hello world": {
			Input:    "hello world",
			Key:      "secretkey",
			Expected: "ae6cd2605d622316564d1f76bfc0c04f89d9fafb14f45b3e18c2a3e28bdef29d",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Key:      "secretkey",
			Expected: "34a4c7e48c2838a6e86f312018600b00b86e6bd68f6898b7a50369c1288ba218",
		},
	}

	// Base64HashHMACSHA256TestTable is used by both the standard Go tests and
	// also the Terraform acceptance tests.
	// <github.com/golang/go/wiki/TableDrivenTests>
	Base64HashHMACSHA256TestTable = map[string]struct {
		Input    string
		Key      string
		Expected string
	}{
		"blank": {
			Input:    "",
			Key:      "secretkey",
			Expected: "DsL74C6nw+tt1zwS6yz/yQYSgN/INlzc+lJBxuPZyac=",
		},
		"a": {
			Input:    "a",
			Key:      "secretkey",
			Expected: "8pCj/cZilXiJKa5N8mVjycI7t8lPk72XQT8znDUWKSg=",
		},
		"hello world": {
			Input:    "hello world",
			Key:      "secretkey",
			Expected: "rmzSYF1iIxZWTR92v8DAT4nZ+vsU9Fs+GMKj4ove8p0=",
		},
		"hello world\n": {
			Input:    "hello world\n",
			Key:      "secretkey",
			Expected: "NKTH5IwoOKbobzEgGGALALhua9aPaJi3pQNpwSiLohg=",
		},
	}
)
