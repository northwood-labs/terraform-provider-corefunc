package corefunc

import (
	"fmt"
	"strings"
	"testing"

	"github.com/northwood-labs/terraform-provider-corefunc/testfixtures"

	"github.com/google/go-cmp/cmp"
)

func TestTruncateLabel(t *testing.T) {
	for name, tc := range testfixtures.TruncateLabelTestTable {
		t.Run(name, func(t *testing.T) {
			actual := TruncateLabel(tc.MaxLength, tc.Prefix, tc.Label)
			diff := cmp.Diff(tc.Expected, actual)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func BenchmarkTruncateLabel(b *testing.B) {
	b.ReportAllocs()
	for name, tc := range testfixtures.TruncateLabelTestTable {
		b.ResetTimer()
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				TruncateLabel(tc.MaxLength, tc.Prefix, tc.Label)
			}
		})
	}
}

func BenchmarkTruncateLabelParallel(b *testing.B) {
	b.ReportAllocs()
	for name, tc := range testfixtures.TruncateLabelTestTable {
		b.ResetTimer()
		b.Run(name, func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					TruncateLabel(tc.MaxLength, tc.Prefix, tc.Label)
				}
			})
		})
	}
}

func FuzzTruncateLabel(f *testing.F) {
	for _, tc := range testfixtures.TruncateLabelTestTable {
		f.Add( // Use f.Add to provide a seed corpus
			fmt.Sprintf("%s:::::%s", tc.Prefix, tc.Label),
		)
	}

	f.Fuzz(
		func(t *testing.T, input string) {
			if !strings.Contains(input, ":::::") {
				t.Skip("skipping test; input missing ':::::' delimiter")
			}

			pieces := strings.Split(input, ":::::")

			_ = TruncateLabel(0, pieces[0], pieces[1])
			_ = TruncateLabel(3, pieces[0], pieces[1])
			_ = TruncateLabel(5, pieces[0], pieces[1])
			_ = TruncateLabel(7, pieces[0], pieces[1])
			_ = TruncateLabel(8, pieces[0], pieces[1])
			_ = TruncateLabel(10, pieces[0], pieces[1])
			_ = TruncateLabel(20, pieces[0], pieces[1])
			_ = TruncateLabel(30, pieces[0], pieces[1])
			_ = TruncateLabel(32, pieces[0], pieces[1])
			_ = TruncateLabel(33, pieces[0], pieces[1])
			_ = TruncateLabel(64, pieces[0], pieces[1])
			_ = TruncateLabel(80, pieces[0], pieces[1])
			_ = TruncateLabel(100, pieces[0], pieces[1])
			_ = TruncateLabel(120, pieces[0], pieces[1])
			_ = TruncateLabel(128, pieces[0], pieces[1])
		},
	)
}
