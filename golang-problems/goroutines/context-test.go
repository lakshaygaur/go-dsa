package thread_test

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type CtxDetails struct {
	deadline time.Time
	id       int
}

func ReadContext(ctx context.Context, wg *sync.WaitGroup) {
	for {
		defer wg.Done()
		select {
		case ch := <-ctx.Value("details").(chan CtxDetails):
			// Handle the context cancellation or timeout
			fmt.Println("Reading thread details for id: ", ch.id)
			if time.Now().After(ch.deadline) {
				fmt.Println("Context deadline exceeded!!")
				return
			}
		default:
			fmt.Println("No details provided yet, waiting...")
		}
		time.Sleep(1 * time.Second)
	}
}

func ShareContext() {
	var wg sync.WaitGroup
	ctx := context.Background()
	details1 := CtxDetails{
		deadline: time.Now().Add(5 * time.Second),
		id:       1,
	}
	ch := make(chan CtxDetails)
	childCtx := context.WithValue(ctx, "details", ch)
	wg.Add(1)
	go ReadContext(childCtx, &wg)

	// 2nd thread
	childCtx = context.WithValue(ctx, "details", ch)
	wg.Add(1)
	go ReadContext(childCtx, &wg)
	time.Sleep(2 * time.Second)
	ch <- details1
	time.Sleep(2 * time.Second)
	ch <- CtxDetails{
		deadline: time.Now().Add(3 * time.Second),
		id:       2,
	}
	wg.Wait()
}
