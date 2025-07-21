package lcs_string

/**
    a c e
  0 0 0 0 // extra cells to hold nil values for comparison
a 0 1 1 1 // compare each text1 with text2 in terms of row and column
b 0 1 1 1 // increase the cell value by its previous diangonal cell
c 0 1 2 2 // otherwise put max of left and top cell
d 0 1 2 2
e 0 1 2 3

*/

func longestCommonSubsequence(text1 string, text2 string) int {
	m := len(text1) + 1 // text1
	n := len(text2) + 1 // text2
	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if text1[i-1] == text2[j-1] { // increase the cell value by its previous diangonal cell
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[m-1][n-1]
}
