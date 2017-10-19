package main

import (
	"fmt"
	"math/rand"
)

/*
func rand_generator_1() int {
	return rand.Int()
}*/

func rand_generator_2() chan int {
	out := make(chan int)

	go func() {
		for {
			out <- rand.Int()
		}
	}()

	return out
}

func rand_generator_3() chan int {
	rand_generator_1 := rand_generator_2()
	rand_generator_2 := rand_generator_2()

	out := make(chan int)

	go func() {
		for {
			out <- <-rand_generator_1
		}
	}()

	go func() {
		for {
			out <- <-rand_generator_2
		}
	}()

	return out
}

func main() {
	rand_service_handler := rand_generator_3()
	fmt.Printf("%d\n", <-rand_service_handler)
}
