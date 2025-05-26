package find_diff_in_arr

import (
	"reflect"
	"testing"
)

func TestFindDifferenceInArrays(t *testing.T) {

	t.Run("nums1 = [1,2,3], nums2 = [2,4,6]", func(t *testing.T) {
		got := findDifference([]int{1, 2, 3}, []int{2, 4, 6})
		want := [][]int{{1, 3}, {4, 6}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})
}
