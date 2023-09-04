package testfixtures

// TruncateLabelTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// https://github.com/golang/go/wiki/TableDrivenTests
var TruncateLabelTestTable = map[string]struct {
	Prefix    string
	Label     string
	MaxLength int64
	Expected  string
}{
	// ---------------------------------------------------------------------
	// Very simple tests meant to illustrate the logic.

	"simple0": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 0,
		Expected:  "",
	},
	"simple3": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 3, // lint:allow_raw_number
		Expected:  "…",
	},
	"simple5": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 5, // lint:allow_raw_number
		Expected:  "…: …",
	},
	"simple7": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 7, // lint:allow_raw_number
		Expected:  "…: Nam…",
	},
	"simple8": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 8, // lint:allow_raw_number
		Expected:  "…: Name…",
	},
	"simple10": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 10, // lint:allow_raw_number
		Expected:  "N…: Name…",
	},
	"simple20": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 20, // lint:allow_raw_number
		Expected:  "NW-TES…: Name of my…",
	},
	"simple30": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 30, // lint:allow_raw_number
		Expected:  "NW-TEST-FIR…: Name of my moni…",
	},
	"simple32": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 32, // lint:allow_raw_number
		Expected:  "NW-TEST-FIRS…: Name of my monit…",
	},
	"simple33": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 33, // lint:allow_raw_number
		Expected:  "NW-TEST-FIRST: Name of my monitor",
	},
	"simple64": {
		Prefix:    "NW-TEST-FIRST",
		Label:     "Name of my monitor",
		MaxLength: 64, // lint:allow_raw_number
		Expected:  "NW-TEST-FIRST: Name of my monitor",
	},

	// ---------------------------------------------------------------------
	// truncate() has a "balancing" algorithm between the prefix and the label.
	// These are tests where the prefix is longer than the label.

	"longPrefix0": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 0,
		Expected:  "",
	},
	"longPrefix3": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 3, // lint:allow_raw_number
		Expected:  "…",
	},
	"longPrefix5": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 5, // lint:allow_raw_number
		Expected:  "…: …",
	},
	"longPrefix7": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 7, // lint:allow_raw_number
		Expected:  "NW-…: …",
	},
	"longPrefix8": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 8, // lint:allow_raw_number
		Expected:  "NW-A…: …",
	},
	"longPrefix10": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 10, // lint:allow_raw_number
		Expected:  "NW-AWS…: …",
	},
	"longPrefix30": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 30, // lint:allow_raw_number
		Expected:  "NW-AWS_ACCOUNT_NAME-PLATFO…: …",
	},
	"longPrefix64": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 64, // lint:allow_raw_number
		Expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-…: K8S Pods No…",
	},
	"longPrefix80": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 80, // lint:allow_raw_number
		Expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_…: K8S Pods Not Runnin…",
	},
	"longPrefix100": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 100,                                                                                                    // lint:allow_raw_number lint:ignore_length
		Expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-…: K8S Pods Not Running Deployme…", // lint:ignore_length
	},
	"longPrefix120": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 120,                                                                                                                  // lint:allow_raw_number lint:ignore_length
		Expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL: K8S Pods Not Running Deployment Check", // lint:ignore_length
	},
	"longPrefix128": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 128,                                                                                                                  // lint:allow_raw_number lint:ignore_length
		Expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL: K8S Pods Not Running Deployment Check", // lint:ignore_length
	},
	"longPrefix256": {
		Prefix:    "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 256,                                                                                                                  // lint:allow_raw_number lint:ignore_length
		Expected:  "NW-AWS_ACCOUNT_NAME-PLATFORM_NAME-APP_NAME-INFRA-RUNTEAM_NAME-PROD-CRITICAL: K8S Pods Not Running Deployment Check", // lint:ignore_length
	},

	// ---------------------------------------------------------------------
	// truncate() has a "balancing" algorithm between the prefix and the label.
	// These are tests where the label is longer than the prefix.

	"longLabel0": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 0,
		Expected:  "",
	},
	"longLabel3": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 3, // lint:allow_raw_number
		Expected:  "…",
	},
	"longLabel5": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 5, // lint:allow_raw_number
		Expected:  "…: …",
	},
	"longLabel7": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 7, // lint:allow_raw_number
		Expected:  "…: K8S…",
	},
	"longLabel8": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 8, // lint:allow_raw_number
		Expected:  "…: K8S…",
	},
	"longLabel10": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 10, // lint:allow_raw_number
		Expected:  "…: K8S Po…",
	},
	"longLabel30": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 30, // lint:allow_raw_number
		Expected:  "…: K8S Pods Not Running Deplo…",
	},
	"longLabel64": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 64, // lint:allow_raw_number
		Expected:  "NW-ZZZ-CLOU…: K8S Pods Not Running Deployment Check K8S Pods No…",
	},
	"longLabel80": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 80, // lint:allow_raw_number
		Expected:  "NW-ZZZ-CLOUD-TEST-A…: K8S Pods Not Running Deployment Check K8S Pods Not Runnin…",
	},
	"longLabel100": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 100,                                                                                                    // lint:allow_raw_number lint:ignore_length
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-P…: K8S Pods Not Running Deployment Check K8S Pods Not Running Deployme…", // lint:ignore_length
	},
	"longLabel120": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 120,                                                                                                                  // lint:allow_raw_number lint:ignore_length
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check", // lint:ignore_length
	},
	"longLabel128": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 128,                                                                                                                  // lint:allow_raw_number lint:ignore_length
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check", // lint:ignore_length
	},
	"longLabel256": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check",
		MaxLength: 256,                                                                                                                  // lint:allow_raw_number lint:ignore_length
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check K8S Pods Not Running Deployment Check", // lint:ignore_length
	},

	// ---------------------------------------------------------------------
	// truncate() has a "balancing" algorithm between the prefix and the label.
	// These are tests where the prefix and the label are the same length.

	"balanced0": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 0,
		Expected:  "",
	},
	"balanced3": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 3, // lint:allow_raw_number
		Expected:  "…",
	},
	"balanced5": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 5, // lint:allow_raw_number
		Expected:  "…: …",
	},
	"balanced7": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 7, // lint:allow_raw_number
		Expected:  "NW…: K…",
	},
	"balanced8": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 8, // lint:allow_raw_number
		Expected:  "NW…: K8…",
	},
	"balanced10": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 10, // lint:allow_raw_number
		Expected:  "NW-…: K8S…",
	},
	"balanced30": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 30, // lint:allow_raw_number
		Expected:  "NW-ZZZ-CLOUD-…: K8S Pods Not…",
	},
	"balanced64": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 64, // lint:allow_raw_number
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PR…: K8S Pods Not Running Deploymen…",
	},
	"balanced80": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 80, // lint:allow_raw_number
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check",
	},
	"balanced100": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 100, // lint:allow_raw_number
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check",
	},
	"balanced120": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 120, // lint:allow_raw_number
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check",
	},
	"balanced128": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 128, // lint:allow_raw_number
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check",
	},
	"balanced256": {
		Prefix:    "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT",
		Label:     "K8S Pods Not Running Deployment Check",
		MaxLength: 256, // lint:allow_raw_number
		Expected:  "NW-ZZZ-CLOUD-TEST-APP-CLOUD-PROD-CRIT: K8S Pods Not Running Deployment Check",
	},
}
