package concurrence

import "fmt"

func main() {
	src := make(chan int)
	dest := make(chan int, 3)

	go func() {
		// 生产者协程，产生数据
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()

	go func() {
		// 消费者协程，消费数据
		defer close(dest)
		for i := range src {
			dest <- i * i
		}
	}()

	for i := range dest {
		fmt.Println(i)
	}
}
