package thread_test

import (
	"fmt"
)

// print multiples of 5 up until 50 -> 10 iterations
// go routine 1 -> creates a multiple
// go routine 2 -> will print the multiple

func PrintMultiple(ch chan int) {
	for i := range <-ch {
		fmt.Println(" Channel=", i)
	}
}

func CreateMultiple(chBool chan bool, ch chan int) {
	iterations := 10
	for i := 1; i <= iterations; i++ {
		multiple := 5 * i
		fmt.Println("updating multiple", multiple)
		ch <- 5
	}
	// wg.Done()
	chBool <- true
	// close(ch)
}
