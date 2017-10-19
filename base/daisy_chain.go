package main

import "fmt"

func xrange() chan int {
	ch := make(chan int)

	go func() { // 开启一个goroutine
		for i := 2; ; i++ { // 从2开始自增
			ch <- i // 管道索要数据才把i添加进管道
		}
	}()

	return ch
}

// 筛选器
func filter(in chan int, number int) chan int {
	// 输入一个整数队列，过滤number的倍数，不是number倍数输出管道
	out := make(chan int)

	go func() {
		for {
			i := <-in // 从生成器取出一个数字

			if i%number != 0 {
				out <- i
			}
		}
	}()

	return out
}

func main() {
	const max = 100  // 找出100以内的所有素数
	nums := xrange() // 初始化一个整数生成器
	number := <-nums // 从生成器中取出一个初始化整数（2）

	for number <= max { // number为筛子，当筛子超过max结束筛选
		fmt.Println(number)
		nums = filter(nums, number) // 筛掉number的倍数
		number = <-nums             // 更新筛子
	}
}
