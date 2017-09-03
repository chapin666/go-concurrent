package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 42
		fmt.Println("I'm a child goroutine")
	}()

	<-ch
	fmt.Println("over")
}
