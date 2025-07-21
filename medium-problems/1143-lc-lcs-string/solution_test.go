package lcs_string

import "testing"

func TestSubarraySum(t *testing.T) {

	t.Run("text1=\"abcde\" \n text2=\"ace\"", func(t *testing.T) {
		input1 := "abcde"
		input2 := "ace"
		got := longestCommonSubsequence(input1, input2)
		want := 3

		if got != want {
			t.Errorf("got %q, wanted %q", got, want)
		}
	})

}
