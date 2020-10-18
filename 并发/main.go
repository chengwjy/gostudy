package main
import "fmt"
import "sync"

var wg sync.WaitGroup

// 互斥锁
var lock sync.Mutex
func main(){ // 开启一个主goroutine来执行main函数
	// Go语言并发通过goroutine实现
	// goroutine类似于线程，属于用户态的线程
	// 我们可以根据需要创建成千上万个goroutine并发工作
	// goroutine是Go语言的运行时调度完成，而线程是由操作系统调度完成
	// Go语言还提供channel在多个goroutine间通信
	// goroutine 和 channel是Go语言秉承的CSP并发模式的重要实现基础
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1) // 计数牌+1
		go hello(i, &count) // 开启一个goroutine来执行hello这个函数
	}
	
	wg.Wait() // 阻塞，等所有小兵都干完活才收兵
	fmt.Println("count", count) // 100000
}

func hello(num int, count *int){
	for i := 0; i < 100; i++ {
		// 加锁
		lock.Lock()
		*count++

		// 释放锁
		lock.Unlock()
		fmt.Println("hello", num, i)
	}
	wg.Done() // 通知wg把计数牌-1
	
}