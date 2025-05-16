package aws

import "fmt"

/**
case1 : when ! is 0
x' ==> '01', y' ==> '10'
x'*x + y'*y = out1

case2 : when ! is 1
x'' ==> '01', y'' ==> '10'
x''*x + y''*y = out2

min(out1, out2) = output

return output
*/

func getMinErrors(input string) int {
	return 0
}

func runThis() {
	input := "101!1"

	fmt.Println("input=", input)
}
