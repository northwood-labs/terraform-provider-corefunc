// Copyright 2023-2024, Northwood Labs
// Copyright 2023-2024, Ryan Parman <rparman@northwood-labs.com>
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
	"net"

	"github.com/apparentlymart/go-cidr/cidr"
)

/*
CIDRContains checks to see if an IP address or CIDR block is contained within
another CIDR block.

Ported from OpenTofu. This function licensed as MPL-2.0.

----

  - containingCidr (string): A CIDR range to check as a containing range.

  - containedIpOrCidr (string): An IP address or CIDR range to check as a
    contained range.
*/
func CIDRContains(containingCidr, containedIPOrCidr string) (bool, error) {
	// The first argument must be a CIDR prefix.
	_, containing, err := net.ParseCIDR(containingCidr)
	if err != nil {
		return false, fmt.Errorf(
			"invalid containing CIDR `%s`: %w",
			containingCidr,
			err,
		)
	}

	// The second argument can be either an IP address or a CIDR prefix. We will
	// try parsing it as an IP address first.
	startIP := net.ParseIP(containedIPOrCidr)

	var endIP net.IP

	// If the second argument did not parse as an IP, we will try parsing it as
	// a CIDR prefix.
	if startIP == nil {
		_, contained, err := net.ParseCIDR(containedIPOrCidr)
		// If that also fails, we'll return an error.
		if err != nil {
			return false, fmt.Errorf(
				"invalid IP address or containing CIDR: %s",
				containedIPOrCidr,
			)
		}

		// Otherwise, we will want to know the start and the end IP of the
		// prefix, so that we can check whether both are contained in the
		// containing prefix.
		startIP, endIP = cidr.AddressRange(contained)
	}

	// We require that both addresses are of the same type, so that we can't
	// accidentally compare an IPv4 address to an IPv6 prefix. The underlying Go
	// function will always return false if this happens, but we want to return
	// an error instead so that the caller can distinguish between a
	// "legitimate" false result and an erroneous check.
	if (startIP.To4() == nil) != (containing.IP.To4() == nil) {
		return false, fmt.Errorf(
			"address family mismatch: %s vs. %s",
			containingCidr,
			containedIPOrCidr,
		)
	}

	// If the second argument was an IP address, we will check whether it is
	// contained in the containing prefix, and that's our result.
	result := containing.Contains(startIP)

	// If the second argument was a CIDR prefix, we will also check whether the
	// end IP of the prefix is contained in the containing prefix. Once CIDR is
	// contained in another CIDR if both the start and the end IP of the
	// contained CIDR are contained in the containing CIDR.
	if endIP != nil {
		result = result && containing.Contains(endIP)
	}

	return result, nil
}