package thread_test

import (
	"fmt"
	"sync"
)

func myRoutine(ch1 chan int, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(" Thread #: ", id, "  reading data ", <-ch1)
}

func RunThreads() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(2)

	go myRoutine(ch, 1, &wg)
	go myRoutine(ch, 2, &wg)

	ch <- 5
	ch <- 10

	wg.Wait()
}
