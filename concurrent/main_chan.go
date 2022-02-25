package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 3)
	for i := 0; i < 10; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			log.Println(i)
			time.Sleep(time.Second)
			<-ch
		}(i)
	}
	wg.Wait()
}

// 创建缓冲区大小为 3 的 channel
// 开启协程前，调用 ch <- struct{}{}，若缓存区满，则阻塞。 协程任务结束，调用 <-ch 释放缓冲区。
