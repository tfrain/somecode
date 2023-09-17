package interview

import (
	"fmt"
)

// 发送一个序列 2, 3, 4, ... 到 channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// 把 channel 'in'中的数据复制到 channel 'out'。
// 删除那些可以被'prime'整除的数据。
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // 从 channel 'in'中获取数据。
		// fmt.Printf("filter i: %v \n", i)
		// fmt.Printf("filter prime: %v \n", prime)
		if i%prime != 0 {
			out <- i // 把'i'发送到 channel 'out'中。
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func main() {
	ch := make(chan int) // 创建一个新的 channel。
	go Generate(ch)      // 启动 goroutine.
	for i := 0; i < 5; i++ {
		prime := <-ch
		fmt.Printf("main prime: %v \n", prime)
		ch1 := make(chan int)
		go Filter(ch, ch1, prime)
		ch = ch1
	}
}
