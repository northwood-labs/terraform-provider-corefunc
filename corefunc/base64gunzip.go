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
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"io"
	"strings"
)

/*
Base64Gunzip is a function that decodes a Base64-encoded string, then
decompresses the result with gzip. Supports both padded and non-padded Base64
strings.

Uses the "standard" Base64 alphabet as defined in RFC 4648 §4.
<https://datatracker.ietf.org/doc/html/rfc4648#section-4>

There is a data limit of 10 MiB (10485760 bytes) for the decompressed data. This
is to avoid "decompression bomb" vulnerabilities.

This functionality is built into OpenTofu 1.8, but is missing in Terraform 1.9.
This also provides a 1:1 implementation that can be used with Terratest or other
Go code.

----

  - input (string): A string of gzipped then Base64-encoded data.
*/
func Base64Gunzip(input string) (string, error) {
	const maxDecodedLen = 10485760 // 10 MiB

	encoder := base64.StdEncoding.WithPadding(base64.NoPadding)

	if strings.HasSuffix(input, "=") {
		encoder = base64.StdEncoding
	}

	decodedLen := encoder.DecodedLen(len(input))
	decodedBytes := make([]byte, decodedLen)

	n, err := encoder.Decode(decodedBytes, []byte(input))
	if err != nil {
		return "", fmt.Errorf("failed to Base64-decode the data: %w", err)
	}

	decodedBytes = decodedBytes[:n]
	decodedBytesBuf := bytes.NewReader(decodedBytes)

	unzippedReader, err := gzip.NewReader(decodedBytesBuf)
	if err != nil {
		return "", fmt.Errorf("failed to decompress the gzipped data: %w", err)
	}

	var unzippedBytes bytes.Buffer

	limitedReader := io.LimitReader(unzippedReader, maxDecodedLen)

	_, err = io.Copy(&unzippedBytes, limitedReader)
	if err != nil {
		return "", fmt.Errorf("failed to read the decompressed gzipped data: %w", err)
	}

	err = unzippedReader.Close()
	if err != nil {
		return "", fmt.Errorf("failed to close the decompressed gzipped data: %w", err)
	}

	return unzippedBytes.String(), nil
}
