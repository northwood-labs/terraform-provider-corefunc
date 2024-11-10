// Copyright 2023-2024, Northwood Labs
// Copyright 2023-2024, Ryan Parman <rparman@northwood-labs.com>
//
// Licensed under the Apache License, Version 2.0 (the "License";
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

// NetCidrContainsTestTable is used by both the standard Go tests and also the
// Terraform acceptance tests.
// <https://github.com/golang/go/wiki/TableDrivenTests>
var NetCidrContainsTestTable = map[string]struct { // lint:no_dupe
	ContainerCidr     string
	ContainedIPOrCidr string
	Expected          bool
	ExpectedErr       bool
}{
	"192.168.2.0/20 contains 192.168.2.1": {
		// IPv4, contained (IP).
		ContainerCidr:     "192.168.2.0/20",
		ContainedIPOrCidr: "192.168.2.1",
		Expected:          true,
		ExpectedErr:       false,
	},
	"192.168.2.0/20 contains 192.168.2.0/22": {
		// IPv4, contained (CIDR).
		ContainerCidr:     "192.168.2.0/20",
		ContainedIPOrCidr: "192.168.2.0/22",
		Expected:          true,
		ExpectedErr:       false,
	},
	"192.168.2.0/20 contains 192.126.2.1": {
		// IPv4, not contained.
		ContainerCidr:     "192.168.2.0/20",
		ContainedIPOrCidr: "192.126.2.1",
		Expected:          false,
		ExpectedErr:       false,
	},
	"192.168.2.0/20 contains 192.126.2.0/18": {
		// IPv4, not contained (CIDR).
		ContainerCidr:     "192.168.2.0/20",
		ContainedIPOrCidr: "192.126.2.0/18",
		Expected:          false,
		ExpectedErr:       false,
	},
	"fe80::/48 contains fe80::1": {
		// IPv6, contained.
		ContainerCidr:     "fe80::/48",
		ContainedIPOrCidr: "fe80::1",
		Expected:          true,
		ExpectedErr:       false,
	},
	"fe80::/48 contains fe81::1": {
		// IPv6, not contained.
		ContainerCidr:     "fe80::/48",
		ContainedIPOrCidr: "fe81::1",
		Expected:          false,
		ExpectedErr:       false,
	},
	"192.168.2.0/20 contains fe80::1": {
		// Address family mismatch: IPv4 containing_prefix, IPv6
		// contained_ip_or_prefix (IP).
		ContainerCidr:     "192.168.2.0/20",
		ContainedIPOrCidr: "fe80::1",
		Expected:          false,
		ExpectedErr:       true,
	},
	"192.168.2.0/20 contains fe80::/24": {
		// Address family mismatch: IPv4 containing_prefix, IPv6
		// contained_ip_or_prefix (prefix).
		ContainerCidr:     "192.168.2.0/20",
		ContainedIPOrCidr: "fe80::/24",
		Expected:          false,
		ExpectedErr:       true,
	},
	"fe80::/48 contains 192.168.2.1": {
		// Address family mismatch: IPv6 containing_prefix, IPv4
		// contained_ip_or_prefix (IP).
		ContainerCidr:     "fe80::/48",
		ContainedIPOrCidr: "192.168.2.1",
		Expected:          false,
		ExpectedErr:       true,
	},
	"fe80::/48 contains 192.168.2.0/20": {
		// Address family mismatch: IPv6 containing_prefix, IPv4
		// contained_ip_or_prefix (prefix).
		ContainerCidr:     "fe80::/48",
		ContainedIPOrCidr: "192.168.2.0/20",
		Expected:          false,
		ExpectedErr:       true,
	},
	"not-a-cidr contains 192.168.2.1": {
		// Input error: invalid CIDR address.
		ContainerCidr:     "not-a-cidr",
		ContainedIPOrCidr: "192.168.2.1",
		Expected:          false,
		ExpectedErr:       true,
	},
	"192.168.2.0/20 contains not-an-address": {
		// Input error: invalid IP address.
		ContainerCidr:     "192.168.2.0/20",
		ContainedIPOrCidr: "not-an-address",
		Expected:          false,
		ExpectedErr:       true,
	},
}
