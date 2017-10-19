package main

import "fmt"

// 整数生成器
func generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i
	}
}

// 过滤器
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}

func main() {
	ch := make(chan int)

	go generate(ch)

	ch1 := make(chan int)
	for i := 0; i < 10; i++ {
		prime := <-ch
		fmt.Println(prime)
		go filter(ch, ch1, prime)
		ch = ch1
	}

}
