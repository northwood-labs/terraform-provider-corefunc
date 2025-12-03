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

package types // lint:allow_package_name

import "time"

const (
	// -------------------------------------------------------------------------
	// Standard library time layouts

	// ANSIC comes from the standard library time package.
	ANSIC TimeLayout = time.ANSIC

	// DateOnly comes from the standard library time package.
	DateOnly TimeLayout = time.DateOnly

	// DateTime comes from the standard library time package.
	DateTime TimeLayout = time.DateTime

	// Kitchen comes from the standard library time package.
	Kitchen TimeLayout = time.Kitchen

	// Layout comes from the standard library time package.
	Layout TimeLayout = time.Layout

	// RFC1036 comes from the standard library time package.
	RFC1036 TimeLayout = time.RFC822

	// RFC1036Z comes from the standard library time package.
	RFC1036Z TimeLayout = time.RFC822Z

	// RFC1123 comes from the standard library time package.
	RFC1123 TimeLayout = time.RFC1123

	// RFC1123Z comes from the standard library time package.
	RFC1123Z TimeLayout = time.RFC1123Z

	// RFC2822 comes from the standard library time package.
	RFC2822 TimeLayout = time.RFC1123

	// RFC2822Z comes from the standard library time package.
	RFC2822Z TimeLayout = time.RFC1123Z

	// RFC3339 comes from the standard library time package.
	RFC3339 TimeLayout = time.RFC3339

	// RFC3339Nano comes from the standard library time package.
	RFC3339Nano TimeLayout = time.RFC3339Nano

	// RFC822 comes from the standard library time package.
	RFC822 TimeLayout = time.RFC822

	// RFC822Z comes from the standard library time package.
	RFC822Z TimeLayout = time.RFC822Z

	// RFC850 comes from the standard library time package.
	RFC850 TimeLayout = time.RFC850

	// RubyDate comes from the standard library time package.
	RubyDate TimeLayout = time.RubyDate

	// Stamp comes from the standard library time package.
	Stamp TimeLayout = time.Stamp

	// StampMicro comes from the standard library time package.
	StampMicro TimeLayout = time.StampMicro

	// StampMilli comes from the standard library time package.
	StampMilli TimeLayout = time.StampMilli

	// StampNano comes from the standard library time package.
	StampNano TimeLayout = time.StampNano

	// TimeOnly comes from the standard library time package.
	TimeOnly TimeLayout = time.TimeOnly

	// UnixDate comes from the standard library time package.
	UnixDate TimeLayout = time.UnixDate

	// W3C comes from the standard library time package.
	W3C TimeLayout = time.RFC3339

	// -------------------------------------------------------------------------
	// Additional time layouts ported from PHP

	// COOKIE was ported from PHP.
	COOKIE TimeLayout = "Monday, 02-Jan-2006 15:04:05 MST"

	// ISO8601 was ported from PHP.
	ISO8601 TimeLayout = "2006-01-02T15:04:05Z0700"

	// ISO8601Nano was ported from PHP.
	ISO8601Nano TimeLayout = "2006-01-02T15:04:05.999999999Z0700"

	// RFC7231 was ported from PHP.
	RFC7231 TimeLayout = "Mon, 02 Jan 2006 15:04:05 GMT"
)

// TimeLayout represents a time layout string for use with time parsing and formatting.
type TimeLayout string
