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

package corefunc

import (
	"fmt"
	"os"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/v2/testfixtures"
)

func ExampleHashMD5() {
	output1 := HashMD5("")
	fmt.Println(output1)

	output2 := HashMD5("a")
	fmt.Println(output2)

	output3 := HashMD5("hello world")
	fmt.Println(output3)

	output4 := HashMD5("hello world\n")
	fmt.Println(output4)

	// Output:
	// d41d8cd98f00b204e9800998ecf8427e
	// 0cc175b9c0f1b6a831c399e269772661
	// 5eb63bbbe01eeed093cb22bb8f5acdc3
	// 6f5902ac237024bdd0c176cb93063dc4
}

func ExampleHashSHA1() {
	output1 := HashSHA1("")
	fmt.Println(output1)

	output2 := HashSHA1("a")
	fmt.Println(output2)

	output3 := HashSHA1("hello world")
	fmt.Println(output3)

	output4 := HashSHA1("hello world\n")
	fmt.Println(output4)

	// Output:
	// da39a3ee5e6b4b0d3255bfef95601890afd80709
	// 86f7e437faa5a7fce15d1ddcb9eaeaea377667b8
	// 2aae6c35c94fcfb415dbe95f408b9ce91ee846ed
	// 22596363b3de40b06f981fb85d82312e8c0ed511
}

func ExampleHashSHA256() {
	output1 := HashSHA256("")
	fmt.Println(output1)

	output2 := HashSHA256("a")
	fmt.Println(output2)

	output3 := HashSHA256("hello world")
	fmt.Println(output3)

	output4 := HashSHA256("hello world\n")
	fmt.Println(output4)

	// Output:
	// e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
	// ca978112ca1bbdcafac231b39a23dc4da786eff8147c4e72b9807785afee48bb
	// b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
	// a948904f2f0f479b8f8197694b30184b0d2ed1c1cd2a1ec0fb85d299a192a447
}

func ExampleHashSHA384() {
	output1 := HashSHA384("")
	fmt.Println(output1)

	output2 := HashSHA384("a")
	fmt.Println(output2)

	output3 := HashSHA384("hello world")
	fmt.Println(output3)

	output4 := HashSHA384("hello world\n")
	fmt.Println(output4)

	// Output:
	// 38b060a751ac96384cd9327eb1b1e36a21fdb71114be07434c0cc7bf63f6e1da274edebfe76f65fbd51ad2f14898b95b
	// 54a59b9f22b0b80880d8427e548b7c23abd873486e1f035dce9cd697e85175033caa88e6d57bc35efae0b5afd3145f31
	// fdbd8e75a67f29f701a4e040385e2e23986303ea10239211af907fcbb83578b3e417cb71ce646efd0819dd8c088de1bd
	// 6b3b69ff0a404f28d75e98a066d3fc64fffd9940870cc68bece28545b9a75086b343d7a1366838083e4b8f3ca6fd3c80
}

func ExampleBase64HashMD5() {
	output1 := Base64HashMD5("")
	fmt.Println(output1)

	output2 := Base64HashMD5("a")
	fmt.Println(output2)

	output3 := Base64HashMD5("hello world")
	fmt.Println(output3)

	output4 := Base64HashMD5("hello world\n")
	fmt.Println(output4)

	// Output:
	// 1B2M2Y8AsgTpgAmY7PhCfg==
	// DMF1ucDxtqgxw5niaXcmYQ==
	// XrY7u+Ae7tCTyyK7j1rNww==
	// b1kCrCNwJL3QwXbLkwY9xA==
}

func ExampleBase64HashSHA1() {
	output1 := Base64HashSHA1("")
	fmt.Println(output1)

	output2 := Base64HashSHA1("a")
	fmt.Println(output2)

	output3 := Base64HashSHA1("hello world")
	fmt.Println(output3)

	output4 := Base64HashSHA1("hello world\n")
	fmt.Println(output4)

	// Output:
	// 2jmj7l5rSw0yVb/vlWAYkK/YBwk=
	// hvfkN/qlp/zhXR3cuerq6jd2Z7g=
	// Kq5sNclPz7QV2+lfQIuc6R7oRu0=
	// IlljY7PeQLBvmB+4XYIxLowO1RE=
}

func ExampleBase64HashSHA256() {
	output1 := Base64HashSHA256("")
	fmt.Println(output1)

	output2 := Base64HashSHA256("a")
	fmt.Println(output2)

	output3 := Base64HashSHA256("hello world")
	fmt.Println(output3)

	output4 := Base64HashSHA256("hello world\n")
	fmt.Println(output4)

	// Output:
	// 47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=
	// ypeBEsobvcr6wjGzmiPcTaeG7/gUfE5yuYB3ha/uSLs=
	// uU0nuZNNPgilLlLX2n2r+sSE7+N6U4DukIj3rOLvzek=
	// qUiQTy8PR5uPgZdpSzAYSw0u0cHNKh7A+4XSmaGSpEc=
}

func ExampleBase64HashSHA384() {
	output1 := Base64HashSHA384("")
	fmt.Println(output1)

	output2 := Base64HashSHA384("a")
	fmt.Println(output2)

	output3 := Base64HashSHA384("hello world")
	fmt.Println(output3)

	output4 := Base64HashSHA384("hello world\n")
	fmt.Println(output4)

	// Output:
	// OLBgp1GsljhM2TJ+sbHjaiH9txEUvgdDTAzHv2P24donTt6/529l+9Ua0vFImLlb
	// VKWbnyKwuAiA2EJ+VIt8I6vYc0huHwNdzpzWl+hRdQM8qojm1XvDXvrgta/TFF8x
	// /b2OdaZ/KfcBpOBAOF4uI5hjA+oQI5IRr5B/y7g1eLPkF8txzmRu/QgZ3YwIjeG9
	// aztp/wpATyjXXpigZtP8ZP/9mUCHDMaL7OKFRbmnUIazQ9ehNmg4CD5Ljzym/TyA
}

func ExampleBase64HashSHA512() {
	output1 := Base64HashSHA512("")
	fmt.Println(output1)

	output2 := Base64HashSHA512("a")
	fmt.Println(output2)

	output3 := Base64HashSHA512("hello world")
	fmt.Println(output3)

	output4 := Base64HashSHA512("hello world\n")
	fmt.Println(output4)

	// Output:
	// z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==
	// H0D8ktokFpR1CXnubPWC8tXX0o4YM13gWrxU0FYOD1MChgxlK/CNVgJSql50IQVG82n7u86MEs/HlXsmUv6adQ==
	// MJ7MSJwS1utMxA9QyQLytNDtd+5RGnx6m808qG1M2G+YndNbxf9JlnDaNCVbRbDP2DDoH2Bdz33FVC6TrpzXbw==
	// 2zl0qX8kB7fK4a5jfAAwaHoRkTJ01XhJJVjjnBbAF96E6s3Ixi/jTuThK0sUKIF/Cbaidgw/imZM6ulNJDSlkw==
}

func ExampleHashHMACSHA256() {
	output := HashHMACSHA256("message", "secretkey")
	fmt.Println(output)

	// Output:
	// 5c3e2f56de9411068f675ef32ffa12735210b9cbfee2ba521367a3955334a343
}

func ExampleBase64HashHMACSHA256() {
	output := Base64HashHMACSHA256("message", "secretkey")
	fmt.Println(output)

	// Output:
	// XD4vVt6UEQaPZ17zL/oSc1IQucv+4rpSE2ejlVM0o0M=
}

func TestHashMD5(t *testing.T) {
	for name, tc := range testfixtures.HashMD5TestTable {
		t.Run(name, func(t *testing.T) {
			actual := HashMD5(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashSHA1(t *testing.T) {
	for name, tc := range testfixtures.HashSHA1TestTable {
		t.Run(name, func(t *testing.T) {
			actual := HashSHA1(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashSHA256(t *testing.T) {
	for name, tc := range testfixtures.HashSHA256TestTable {
		t.Run(name, func(t *testing.T) {
			actual := HashSHA256(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashSHA384(t *testing.T) {
	for name, tc := range testfixtures.HashSHA384TestTable {
		t.Run(name, func(t *testing.T) {
			actual := HashSHA384(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashSHA512(t *testing.T) {
	for name, tc := range testfixtures.HashSHA512TestTable {
		t.Run(name, func(t *testing.T) {
			actual := HashSHA512(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashSHA3x256(t *testing.T) {
	for name, tc := range testfixtures.HashSHA3x256TestTable {
		t.Run(name, func(t *testing.T) {
			actual := HashSHA3x256(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashSHA3x384(t *testing.T) {
	for name, tc := range testfixtures.HashSHA3x384TestTable {
		t.Run(name, func(t *testing.T) {
			actual := HashSHA3x384(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashSHA3x512(t *testing.T) {
	for name, tc := range testfixtures.HashSHA3x512TestTable {
		t.Run(name, func(t *testing.T) {
			actual := HashSHA3x512(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashMD5(t *testing.T) {
	for name, tc := range testfixtures.Base64HashMD5TestTable {
		t.Run(name, func(t *testing.T) {
			actual := Base64HashMD5(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashSHA1(t *testing.T) {
	for name, tc := range testfixtures.Base64HashSHA1TestTable {
		t.Run(name, func(t *testing.T) {
			actual := Base64HashSHA1(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashSHA256(t *testing.T) {
	for name, tc := range testfixtures.Base64HashSHA256TestTable {
		t.Run(name, func(t *testing.T) {
			actual := Base64HashSHA256(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashSHA384(t *testing.T) {
	for name, tc := range testfixtures.Base64HashSHA384TestTable {
		t.Run(name, func(t *testing.T) {
			actual := Base64HashSHA384(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashSHA512(t *testing.T) {
	for name, tc := range testfixtures.Base64HashSHA512TestTable {
		t.Run(name, func(t *testing.T) {
			actual := Base64HashSHA512(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashSHA3x256(t *testing.T) {
	for name, tc := range testfixtures.Base64HashSHA3x256TestTable {
		t.Run(name, func(t *testing.T) {
			actual := Base64HashSHA3x256(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashSHA3x384(t *testing.T) {
	for name, tc := range testfixtures.Base64HashSHA3x384TestTable {
		t.Run(name, func(t *testing.T) {
			actual := Base64HashSHA3x384(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashSHA3x512(t *testing.T) {
	for name, tc := range testfixtures.Base64HashSHA3x512TestTable {
		t.Run(name, func(t *testing.T) {
			actual := Base64HashSHA3x512(tc.Input)

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashArgon2id(t *testing.T) {
	if os.Getenv("CI") != "" {
		t.Skip("Skipping Argon2id tests in CI environment.")
	}

	for name, tc := range testfixtures.HashArgon2idTestTable {
		t.Run(name, func(t *testing.T) {
			actual, err := HashArgon2id(tc.Input, tc.Salt)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashArgon2id(t *testing.T) {
	for name, tc := range testfixtures.Base64HashArgon2idTestTable {
		t.Run(name, func(t *testing.T) {
			actual, err := Base64HashArgon2id(tc.Input, tc.Salt)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashScrypt(t *testing.T) {
	for name, tc := range testfixtures.HashScryptTestTable {
		t.Run(name, func(t *testing.T) {
			actual, err := HashScrypt(tc.Input, tc.Salt)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashScrypt(t *testing.T) {
	for name, tc := range testfixtures.Base64HashScryptTestTable {
		t.Run(name, func(t *testing.T) {
			actual, err := Base64HashScrypt(tc.Input, tc.Salt)
			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}

			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestHashHMACSHA256(t *testing.T) {
	for name, tc := range testfixtures.HashHMACSHA256TestTable {
		t.Run(name, func(t *testing.T) {
			actual := HashHMACSHA256(tc.Input, tc.Key)
			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func TestBase64HashHMACSHA256(t *testing.T) {
	for name, tc := range testfixtures.Base64HashHMACSHA256TestTable {
		t.Run(name, func(t *testing.T) {
			actual := Base64HashHMACSHA256(tc.Input, tc.Key)
			if actual != tc.Expected {
				t.Errorf("Expected %s, got %s", tc.Expected, actual)
			}
		})
	}
}

func BenchmarkHashMD5(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashMD5TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = HashMD5(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashSHA1(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA1TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = HashSHA1(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashSHA256(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA256TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = HashSHA256(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashSHA384(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA384TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = HashSHA384(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashSHA512(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA512TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = HashSHA512(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashSHA3x256(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA3x256TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = HashSHA3x256(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashSHA3x384(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA3x384TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = HashSHA3x384(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashSHA3x512(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA3x512TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = HashSHA3x512(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashHMACSHA256(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashHMACSHA256TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = HashHMACSHA256(tc.Input, tc.Key) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkBase64HashHMACSHA256(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.Base64HashHMACSHA256TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_ = Base64HashHMACSHA256(tc.Input, tc.Key) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashArgon2id(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashArgon2idTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_, _ = HashArgon2id(tc.Input, tc.Salt) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashScrypt(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashScryptTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for range b.N {
				_, _ = HashScrypt(tc.Input, tc.Salt) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkHashMD5Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashMD5TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = HashMD5(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashSHA1Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA1TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = HashSHA1(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashSHA256Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA256TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = HashSHA256(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashSHA384Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA384TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = HashSHA384(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashSHA512Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA512TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = HashSHA512(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashArgon2idParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashArgon2idTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = HashArgon2id(tc.Input, tc.Salt) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashScryptParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashScryptTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, _ = HashScrypt(tc.Input, tc.Salt) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashSHA3x256Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA3x256TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = HashSHA3x256(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashSHA3x384Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA3x384TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = HashSHA3x384(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashSHA3x512Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashSHA3x512TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = HashSHA3x512(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkHashHMACSHA256Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.HashHMACSHA256TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = HashHMACSHA256(tc.Input, tc.Key) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkBase64HashHMACSHA256Parallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.Base64HashHMACSHA256TestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = Base64HashHMACSHA256(tc.Input, tc.Key) // lint:allow_unhandled
				}
			})
		})
	}
}

func FuzzHashMD5(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashMD5TestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(_ *testing.T, in string) {
			_ = HashMD5(in) // lint:allow_unhandled
		},
	)
}

func FuzzHashSHA1(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashSHA1TestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(_ *testing.T, in string) {
			_ = HashSHA1(in) // lint:allow_unhandled
		},
	)
}

func FuzzHashSHA256(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashSHA256TestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(_ *testing.T, in string) {
			_ = HashSHA256(in) // lint:allow_unhandled
		},
	)
}

func FuzzHashSHA384(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashSHA384TestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(_ *testing.T, in string) {
			_ = HashSHA384(in) // lint:allow_unhandled
		},
	)
}

func FuzzHashSHA512(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashSHA512TestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(_ *testing.T, in string) {
			_ = HashSHA512(in) // lint:allow_unhandled
		},
	)
}

func FuzzHashSHA3x256(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashSHA3x256TestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(_ *testing.T, in string) {
			_ = HashSHA3x256(in) // lint:allow_unhandled
		},
	)
}

func FuzzHashSHA3x384(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashSHA3x384TestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(_ *testing.T, in string) {
			_ = HashSHA3x384(in) // lint:allow_unhandled
		},
	)
}

func FuzzHashSHA3x512(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashSHA3x512TestTable {
		f.Add(tc.Input)
	}

	f.Fuzz(
		func(_ *testing.T, in string) {
			_ = HashSHA3x512(in) // lint:allow_unhandled
		},
	)
}

func FuzzHashArgon2id(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashArgon2idTestTable {
		f.Add(tc.Input, tc.Salt)
	}

	f.Fuzz(
		func(_ *testing.T, in string, salt []byte) {
			_, _ = HashArgon2id(in, salt) // lint:allow_unhandled
		},
	)
}

func FuzzHashScrypt(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashScryptTestTable {
		f.Add(tc.Input, tc.Salt)
	}

	f.Fuzz(
		func(_ *testing.T, in string, salt []byte) {
			_, _ = HashScrypt(in, salt) // lint:allow_unhandled
		},
	)
}

func FuzzHashHMACSHA256(f *testing.F) {
	if os.Getenv("CI") != "" {
		f.Skip("Skipping tests in CI environment.")
	}

	for _, tc := range testfixtures.HashHMACSHA256TestTable {
		f.Add(tc.Input, tc.Key)
	}

	f.Fuzz(
		func(_ *testing.T, in, key string) {
			_ = HashHMACSHA256(in, key) // lint:allow_unhandled
		},
	)
}
