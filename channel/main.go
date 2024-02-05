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
