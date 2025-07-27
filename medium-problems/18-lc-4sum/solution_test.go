package lc4sum

import (
	"reflect"
	"testing"
)

func TestReverseVowels(t *testing.T) {

	t.Run("3[a]2[bc]", func(t *testing.T) {
		got := fourSum([]int{1, 0, -1, 0, -2, 2}, 0)
		want := [][]int{[]int{-1, 0, 0, 1}, []int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
