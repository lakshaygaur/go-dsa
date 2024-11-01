package lcmaxavgsubarray1

import "testing"

func TestMaxAvgSubarray(t *testing.T) {
	testcase := []int{1, 12, -5, -6, 50, 3}
	k := 4
	t.Run("[1 ,12,-5,-6,50,3]", func(t *testing.T) {
		got := findMaxAverage(testcase, k)
		want := 12.75

		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	})
}
