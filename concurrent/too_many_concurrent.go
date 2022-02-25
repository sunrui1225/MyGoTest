package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < math.MaxInt32; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
}

//panic: too many concurrent operations on a single file or socket (max 1048575)
//goroutine 2020817 [running]:
//internal/poll.(*fdMutex).rwlock(0xc0000b0060, 0x10c9a00, 0x11)
//报错是 fmt.Printf 函数引起,内核(kernel)利用文件描述符(file descriptor)来访问文件,简而言之，系统的资源被耗尽了。
// 去掉呢？那程序很可能会因为内存不足而崩溃。程至少需要消耗 2KB 的空间
