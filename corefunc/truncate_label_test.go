package corefunc

import (
	"fmt"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var (
	// Used by both the standard Go tests and also the Terraform acceptance tests.
	// https://github.com/golang/go/wiki/TableDrivenTests
	tests = map[string]struct {
		prefix    string
		label     string
		maxLength int64
		expected  string
	}{
		// ---------------------------------------------------------------------
		// Very simple tests meant to illustrate the logic.

		"simple0": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 0,
			expected:  "",
		},
		"simple3": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 3,
			expected:  "…",
		},
		"simple5": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 5,
			expected:  "…: …",
		},
		"simple7": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 7,
			expected:  "…: Nam…",
		},
		"simple8": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 8,
			expected:  "…: Name…",
		},
		"simple10": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 10,
			expected:  "N…: Name…",
		},
		"simple20": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 20,
			expected:  "NW-TES…: Name of my…",
		},
		"simple30": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 30,
			expected:  "NW-TEST-FIR…: Name of my moni…",
		},
		"simple32": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 32,
			expected:  "NW-TEST-FIRS…: Name of my monit…",
		},
		"simple33": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 33,
			expected:  "NW-TEST-FIRST: Name of my monitor",
		},
		"simple64": {
			prefix:    "NW-TEST-FIRST",
			label:     "Name of my monitor",
			maxLength: 64,
			expected:  "NW-TEST-FIRST: Name of my monitor",
		},

		// ---------------------------------------------------------------------
		// truncate() has a "balancing" algorithm between the prefix and the label.
		// These are tests where the prefix is longer than the label.

		"longPrefix0": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 0,
			expected:  "",
		},
		"longPrefix3": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 3,
			expected:  "…",
		},
		"longPrefix5": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 5,
			expected:  "…: …",
		},
		"longPrefix7": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 7,
			expected:  "NW-…: …",
		},
		"longPrefix8": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 8,
			expected:  "NW-A…: …",
		},
		"longPrefix10": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 10,
			expected:  "NW-AWS…: …",
		},
		"longPrefix30": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 30,
			expected:  "NW-AWS_ACCOUNT_NAME-PLATFO…: …",
		},
		"longPrefix64": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 64,
			expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-…: K8S Pods No…",
		},
		"longPrefix80": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 80,
			expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_…: K8S Pods Not Runnin…",
		},
		"longPrefix100": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 100,
			expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-…: K8S Pods Not Running Deployme…",
		},
		"longPrefix120": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 120,
			expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL: K8S Pods Not Running Deployment Check",
		},
		"longPrefix128": {
			prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 128,
			expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL: K8S Pods Not Running Deployment Check",
		},

		// ---------------------------------------------------------------------
		// truncate() has a "balancing" algorithm between the prefix and the label.
		// These are tests where the label is longer than the prefix.

		"longLabel0": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 0,
			expected:  "",
		},
		"longLabel3": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 3,
			expected:  "…",
		},
		"longLabel5": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 5,
			expected:  "…: …",
		},
		"longLabel7": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 7,
			expected:  "…: K8S…",
		},
		"longLabel8": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 8,
			expected:  "…: K8S…",
		},
		"longLabel10": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 10,
			expected:  "…: K8S Po…",
		},
		"longLabel30": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 30,
			expected:  "…: K8S Pods Not Running Deplo…",
		},
		"longLabel64": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 64,
			expected:  "NW-ZZZ-CLOU…: K8S Pods Not Running Deployment Check K8S Pods No…",
		},
		"longLabel80": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 80,
			expected:  "NW-ZZZ-CLOUD-TEST-A…: K8S Pods Not Running Deployment Check K8S Pods Not Runnin…",
		},
		"longLabel100": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 100,
			expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-P…: K8S Pods Not Running Deployment Check K8S Pods Not Running Deployme…",
		},
		"longLabel120": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 120,
			expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		},
		"longLabel128": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
			maxLength: 128,
			expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		},

		// ---------------------------------------------------------------------
		// truncate() has a "balancing" algorithm between the prefix and the label.
		// These are tests where the prefix and the label are the same length.

		"balanced0": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 0,
			expected:  "",
		},
		"balanced3": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 3,
			expected:  "…",
		},
		"balanced5": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 5,
			expected:  "…: …",
		},
		"balanced7": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 7,
			expected:  "NW…: K…",
		},
		"balanced8": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 8,
			expected:  "NW…: K8…",
		},
		"balanced10": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 10,
			expected:  "NW-…: K8S…",
		},
		"balanced30": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 30,
			expected:  "NW-ZZZ-CLOUD-…: K8S Pods Not…",
		},
		"balanced64": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 64,
			expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PR…: K8S Pods Not Running Deploymen…",
		},
		"balanced80": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 80,
			expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check",
		},
		"balanced100": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 100,
			expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check",
		},
		"balanced120": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 120,
			expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check",
		},
		"balanced128": {
			prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
			label:     "K8S Pods Not Running Deployment Check",
			maxLength: 128,
			expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check",
		},
	}
)

func TestTruncate(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			actual := TruncateLabel(tc.maxLength, tc.prefix, tc.label)
			diff := cmp.Diff(tc.expected, actual)
			if diff != "" {
				t.Fatalf(diff)
			}
		})
	}
}

func BenchmarkTruncate0Parallel(b *testing.B) {
	prefix := "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label := "K8S Pods Not Running Deployment Check"

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TruncateLabel(0, prefix, label)
		}
	})
}

func BenchmarkTruncate3Parallel(b *testing.B) {
	prefix := "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label := "K8S Pods Not Running Deployment Check"

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TruncateLabel(3, prefix, label)
		}
	})
}

func BenchmarkTruncate5Parallel(b *testing.B) {
	prefix := "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label := "K8S Pods Not Running Deployment Check"

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TruncateLabel(5, prefix, label)
		}
	})
}

func BenchmarkTruncate10Parallel(b *testing.B) {
	prefix := "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label := "K8S Pods Not Running Deployment Check"

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TruncateLabel(10, prefix, label)
		}
	})
}

func BenchmarkTruncate32Parallel(b *testing.B) {
	prefix := "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label := "K8S Pods Not Running Deployment Check"

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TruncateLabel(32, prefix, label)
		}
	})
}

func BenchmarkTruncate64Parallel(b *testing.B) {
	prefix := "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label := "K8S Pods Not Running Deployment Check"

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TruncateLabel(64, prefix, label)
		}
	})
}

func BenchmarkTruncate128Parallel(b *testing.B) {
	prefix := "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label := "K8S Pods Not Running Deployment Check"

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TruncateLabel(128, prefix, label)
		}
	})
}

func BenchmarkTruncate256Parallel(b *testing.B) {
	prefix := "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT"
	label := "K8S Pods Not Running Deployment Check"

	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			TruncateLabel(256, prefix, label)
		}
	})
}

func FuzzTruncate(f *testing.F) {
	for _, tc := range tests {
		f.Add( // Use f.Add to provide a seed corpus
			fmt.Sprintf("%s:::::%s", tc.prefix, tc.label),
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
