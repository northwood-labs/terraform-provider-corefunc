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

	"github.com/northwood-labs/terraform-provider-corefunc/corefunc/types"
)

/*
DeepMerge takes a configuration, and a list of items to merge [item1, item2,
...], and returns a single item that is a merge of all items, based on the merge
configuration.

Discussion:

* https://developer.hashicorp.com/terraform/language/functions/merge
* https://opentofu.org/docs/language/functions/merge/
* https://github.com/opentofu/opentofu/issues/790
* https://github.com/hashicorp/terraform/issues/31815
* https://github.com/hashicorp/terraform/issues/24987#issuecomment-719089236

----

  - str (string): The string with which to count bytes/runes.

----

DeepMerge(config, ["a", "b", "c", "d"])
#=> "d"

DeepMerge(config, [1, 2, 3, 4])
#=> 4
*/
func DeepMerge(config types.MergeConfig, elems []any) (any, error) {
	if len(elems) == 0 {
		return nil, nil
	}

	firstType := fmt.Sprintf("%T", elems[0])

	for i := range elems {
		elem := elems[i]

		if fmt.Sprintf("%T", elem) != firstType {
			return nil, fmt.Errorf("all elements must be of the same type")
		}
	}

	return elems[len(elems)-1], nil
}

// ### list_of_items_to_merge
// items of different types are handled based on the type_mismatch_strategy:
// - complex_wins: object and map win over lists which win over literals; the losers are removed from the merge;
// - abort: abort the terraform plan or apply

// if all items are literals, the rightmost item is used

// ### list_merge_strategy
// if the remaining items are lists, the list_merge_strategy specified (which defaults to "replace") is applied:
// - replace: the list of the right most argument applies
// - concat: the sequence of lists is concatenated in the same order as arguments
// - merge: each list is appended with null so they all have the same length,
//   then the deepmerge() is applied to each slice through the lists

// if the remaining items are maps or objects (the most common case), if more
// than one defines the same key or attribute, the deepmerge() is applied to the
// sequence of associated values across all the items

// deepmerge([a, b, c, d]):

// 1. if a, b, c, d are all literals (string, bool, number), then d is return

// 2. if all items are lists, then they are all merged according to
//    list_merge_strategy

// 3. if a and b are lists and c and d are literals, then c and d are discarded
//    and a and b are merged according to list_merge_strategy

// 4. if a and b are maps or objects then c and d are discarded and for each key
//    in a, b, deepmerge() is called on the corresponding list of associated
//    values, and objects that don't have the key are not included. So if a has
//    key1 and key2, and b has key1 and key3, then the set of keys of the result is
//    (key1, key2, key3) and there will be 3 calls to deepmerge():
//    deepmerge([a.key1, b.key1]), deepmerge([a.key2]), and deepmerge([b.key3])
//    (the merge-config is passed to these but left out here for simplicity)

// 5. if all objects are maps or objects then the previous rule is applied to
//    all items instead of just a and b
