package main

import (
	"fmt"
)

func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

func Filter(in <-chan int, out chan<- int, prime int) {

	for {
		i := <-in
		fmt.Println("prime=", prime, " i=", i)
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	in := make(chan int)
	go Generate(in)

	for i := 0; i < 10; i++ {
		prime := <-in
		fmt.Println(prime)
		out := make(chan int)
		go Filter(in, out, prime)
		in = out
	}
}
