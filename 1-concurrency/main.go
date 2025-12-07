package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	numLen := 10
	randCh := make(chan int, numLen)
	sqCh := make(chan int, numLen)

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		for i := 0; i < numLen; i++ {
			randCh <- rand.Intn(100)
		}
		close(randCh)
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for i := 0; i < numLen; i++ {
			n := <-randCh
			sqCh <- n * n
		}
		close(sqCh)
		wg.Done()
	}()
	


	for i := 0; i < numLen; i++ {
		fmt.Printf("%d ", <-sqCh)
	}

	wg.Wait()
}
