// Copyright 2023-2024, Ryan Parman
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

import "fmt"

/*
IntLeftPad pads an integer on the left side with zeroes until it reaches the
desired width. Result is a string.

----

  - num (int64): The number to pad.

  - padWidth (int): The total width of the padded string. If `num` is equivalent
    length (base10) or longer, then no padding will be applied. If the integer
    is anything other than base10, it will be converted to base10.
*/
func IntLeftPad(num int64, padWidth int) string {
	return fmt.Sprintf("%0"+fmt.Sprint(padWidth)+"d", num)
}
