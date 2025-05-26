package uniqueoccurences

import (
	"reflect"
	"testing"
)

func TestUniqueOccurences(t *testing.T) {

	t.Run("arr = [1,2,2,1,1,3]", func(t *testing.T) {
		got := uniqueOccurrences([]int{1, 2, 2, 1, 1, 3})
		want := true

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})

	t.Run("arr =[1,2]", func(t *testing.T) {
		got := uniqueOccurrences([]int{1, 2})
		want := false

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
