package main

import (
	"fmt"
	"time"
)

func timer(d time.Duration) <-chan int {
	c := make(chan int)

	go func() {
		time.Sleep(d)
		c <- 1
	}()

	return c
}

func main() {
	for i := 0; i < 24; i++ {
		c := timer(1 * time.Second)
		v := <-c
		fmt.Println(v)
	}
}
