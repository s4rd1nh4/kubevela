/*
Copyright 2021 The KubeVela Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package resourcetracker

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"k8s.io/utils/pointer"
)

func TestResourceTreePrintOption_getWidthForDetails(t *testing.T) {
	r := require.New(t)
	options := &ResourceTreePrintOptions{}
	r.Equal(math.MaxInt, options._getWidthForDetails(nil))
	options.MaxWidth = pointer.Int(50 + applyTimeWidth)
	r.Equal(30, options._getWidthForDetails([]int{10, 10}))
	r.Equal(math.MaxInt, options._getWidthForDetails([]int{20, 20}))
}

func TestResourceTreePrintOptions_wrapDetails(t *testing.T) {
	r := require.New(t)
	options := &ResourceTreePrintOptions{}
	detail := "test-key: test-val\ttest-data: test-val\ntest-next-line: text-next-value  test-long-key: test long long long long value  test-append: test-append-val"
	r.Equal(
		[]string{
			"test-key: test-val test-data: test-val",
			"test-next-line: text-next-value",
			"test-long-key: test long long long long ",
			"value  test-append: test-append-val",
		},
		options._wrapDetails(detail, 40))
}
