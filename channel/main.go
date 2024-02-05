package main

import (
	"fmt"
	"sync"
	"time"
)

/*
golang 协程退出的3种方法
*/
func worker3(id int, wg *sync.WaitGroup, stopChan chan struct{}) {
	defer wg.Done()
	for {
		select {
		case <-stopChan: //接收退出信号
			fmt.Printf("stop id: %d", id)
		default:
			fmt.Println("working")
			time.Sleep(1 * time.Second)
		}
	}
}

// 使用sync.WaitGroup控制协程退出
/*
1、stopCh 通道用于通知协程退出。当关闭 stopCh 时，所有监听这个通道的协程都会接收到信号，
并优雅地停止执行。本质上成了和1.使用通道（Channel）一样.
2、sync.WaitGroup的真正作用是卡在wg.Wait()处，直到wg.Done()被执行(wg.Add()增加的值被减为0),
才会继续往下执行. 比如往往用于防止goroutine还没执行完,主协程就退出了
*/
func main() {
	var wg sync.WaitGroup
	stopChan := make(chan struct{}) //空结构体channel

	workerCount := 3
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go worker3(i, &wg, stopChan)
	}

	time.Sleep(3 * time.Second) //模拟工作
	close(stopChan)             //发送退出信号给所有协程

	wg.Wait() //等待所有协程退出
	fmt.Println("All worker stopped")
}

// func worker2(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("exit")
// 			return
// 		default:
// 			fmt.Println("working")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// // 使用context 退出
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	go worker2(ctx)
// 	time.Sleep(3 * time.Second)
// 	cancel() //通知协程退出
// 	time.Sleep(1 * time.Second)
// }

// func worker1(exitchan chan bool) {
// 	for {
// 		select {
// 		case <-exitchan:
// 			fmt.Println("exit")
// 			return
// 		default:
// 			fmt.Println("working")
// 			time.Sleep(1 * time.Second)
// 		}
// 	}
// }

// // 使用通道退出
// // func main() {
// // 	exitchan := make(chan bool)
// // 	go worker1(exitchan)

// // 	time.Sleep(3 * time.Second)
// // 	exitchan <- true //通知协程退出
// // 	time.Sleep(1 * time.Second)
// // }
