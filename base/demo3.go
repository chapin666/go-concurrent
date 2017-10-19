package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		fmt.Printf("%d", i)
	}

	wg.Add(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("%d", i)
		}

		defer wg.Done()
	}()

	wg.Wait()
}
