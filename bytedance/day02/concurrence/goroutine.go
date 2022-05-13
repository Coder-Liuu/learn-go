package concurrence

import (
	"fmt"
	"sync"
)

func hello(i int) {
	fmt.Println("current is ", i)
}

func HelloGoRoutine() {
	// 定义等待变量
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		// 计数加一
		wg.Add(1)
		// 开启一个协程
		go func(j int) {
			// 函数执行完毕后，进行wg减一
			defer wg.Done()
			hello(j)
		}(i)
	}
	// 等待全部执行完毕
	wg.Wait()
}
