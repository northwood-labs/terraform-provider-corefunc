// Copyright 2023-2025, Northwood Labs
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
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"
)

const ExpectedGot = "Expected %s, got %s"

func ExampleStrCamel() {
	const (
		TestStrCamelInput  = "test_with_number_123"
		TestStrCamelInput2 = "test with number -123.456e-2"
	)

	output1 := StrCamel(TestStrCamelInput)
	fmt.Println(output1)

	output2 := StrCamel(TestStrCamelInput2)
	fmt.Println(output2)

	// Output:
	// testWithNumber123
	// testWithNumber123456e2
}

func ExampleStrConstant() {
	const (
		TestStrCamelInput  = "test_with_number_123"
		TestStrCamelInput2 = "test with number -123.456e-2"
	)

	output1 := StrConstant(TestStrCamelInput)
	fmt.Println(output1)

	output2 := StrConstant(TestStrCamelInput2)
	fmt.Println(output2)

	// Output:
	// TEST_WITH_NUMBER_123
	// TEST_WITH_NUMBER_123_456E_2
}

func ExampleStrKebab() {
	const (
		TestStrCamelInput  = "test_with_number_123"
		TestStrCamelInput2 = "test with number -123.456e-2"
	)

	output1 := StrKebab(TestStrCamelInput)
	fmt.Println(output1)

	output2 := StrKebab(TestStrCamelInput2)
	fmt.Println(output2)

	// Output:
	// test-with-number-123
	// test-with-number-123-456e-2
}

func ExampleStrPascal() {
	const (
		TestStrCamelInput  = "test_with_number_123"
		TestStrCamelInput2 = "test with number -123.456e-2"
	)

	output1 := StrPascal(TestStrCamelInput, false)
	fmt.Println(output1)

	output2 := StrPascal(TestStrCamelInput2, false)
	fmt.Println(output2)

	// Output:
	// TestWithNumber123
	// TestWithNumber123456e2
}

func ExampleStrSnake() {
	const (
		TestStrCamelInput  = "test_with_number_123"
		TestStrCamelInput2 = "test with number -123.456e-2"
	)

	output1 := StrSnake(TestStrCamelInput)
	fmt.Println(output1)

	output2 := StrSnake(TestStrCamelInput2)
	fmt.Println(output2)

	// Output:
	// test_with_number_123
	// test_with_number_123_456e_2
}

func TestStrCamel(t *testing.T) {
	for name, tc := range testfixtures.StrCamelTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrCamel(tc.Input)

			if output != tc.Expected {
				t.Errorf(ExpectedGot, tc.Expected, output)
			}
		})
	}
}

func TestStrConstant(t *testing.T) {
	for name, tc := range testfixtures.StrConstantTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrConstant(tc.Input)

			if output != tc.Expected {
				t.Errorf(ExpectedGot, tc.Expected, output)
			}
		})
	}
}

func TestStrKebab(t *testing.T) {
	for name, tc := range testfixtures.StrKebabTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrKebab(tc.Input)

			if output != tc.Expected {
				t.Errorf(ExpectedGot, tc.Expected, output)
			}
		})
	}
}

func TestStrPascal(t *testing.T) {
	for name, tc := range testfixtures.StrPascalTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrPascal(tc.Input, tc.AcronymCaps)

			if output != tc.Expected {
				t.Errorf(ExpectedGot, tc.Expected, output)
			}
		})
	}
}

func TestStrSnake(t *testing.T) {
	for name, tc := range testfixtures.StrSnakeTestTable {
		t.Run(name, func(t *testing.T) {
			output := StrSnake(tc.Input)

			if output != tc.Expected {
				t.Errorf(ExpectedGot, tc.Expected, output)
			}
		})
	}
}

func BenchmarkStrCamel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrCamelTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = StrCamel(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrCamelParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrCamelTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrCamel(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkStrConstant(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrConstantTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = StrConstant(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrConstantParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrConstantTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrConstant(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkStrKebab(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrKebabTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = StrKebab(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrKebabParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrKebabTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrKebab(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkStrPascal(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrPascalTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = StrPascal(tc.Input, tc.AcronymCaps) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrPascalParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrPascalTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrPascal(tc.Input, tc.AcronymCaps) // lint:allow_unhandled
				}
			})
		})
	}
}

func BenchmarkStrSnake(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrSnakeTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				_ = StrSnake(tc.Input) // lint:allow_unhandled
			}
		})
	}
}

func BenchmarkStrSnakeParallel(b *testing.B) {
	b.ReportAllocs()

	for name, tc := range testfixtures.StrSnakeTestTable {
		b.Run(name, func(b *testing.B) {
			b.ResetTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_ = StrSnake(tc.Input) // lint:allow_unhandled
				}
			})
		})
	}
}

func FuzzStrCamel(f *testing.F) {
	for _, tc := range testfixtures.StrCamelTestTable {
		f.Add(tc.Input)
		f.Add(tc.Expected)
	}

	f.Fuzz(
		func(t *testing.T, in string) {
			_ = StrCamel(in) // lint:allow_unhandled
		},
	)
}

func FuzzStrConstant(f *testing.F) {
	for _, tc := range testfixtures.StrConstantTestTable {
		f.Add(tc.Input)
		f.Add(tc.Expected)
	}

	f.Fuzz(
		func(t *testing.T, in string) {
			_ = StrConstant(in) // lint:allow_unhandled
		},
	)
}

func FuzzStrKebab(f *testing.F) {
	for _, tc := range testfixtures.StrKebabTestTable {
		f.Add(tc.Input)
		f.Add(tc.Expected)
	}

	f.Fuzz(
		func(t *testing.T, in string) {
			_ = StrKebab(in) // lint:allow_unhandled
		},
	)
}

func FuzzStrPascal(f *testing.F) {
	for _, tc := range testfixtures.StrPascalTestTable {
		f.Add(tc.Input)
		f.Add(tc.Expected)
	}

	f.Fuzz(
		func(t *testing.T, in string) {
			_ = StrPascal(in, true) // lint:allow_unhandled
		},
	)
}

func FuzzStrSnake(f *testing.F) {
	for _, tc := range testfixtures.StrSnakeTestTable {
		f.Add(tc.Input)
		f.Add(tc.Expected)
	}

	f.Fuzz(
		func(t *testing.T, in string) {
			_ = StrSnake(in) // lint:allow_unhandled
		},
	)
}
