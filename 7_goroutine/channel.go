package main

import (
	"fmt"
	"time"
)

// 无缓冲通道
func ChannelTestNoCache() {
	// 创建一个channel，且channel中元素类型为int
	ch := make(chan int)
	go func() {
		defer fmt.Println("子goroutine结束")
		fmt.Println("子goroutine正在运行")
		ch <- 666 // 向channel中写入数据
	}()
	num := <-ch
	fmt.Println("num=", num)
	fmt.Println("ChannelTestNoCache goroutine 结束")
}

// 有缓冲通道测试
/*
channelSize: 缓冲区大小
size: 发送元素个数
	- channelSize >= size:
		发送完所有元素后，主goroutine会继续等待子goroutine结束
	- channelSize < size，有可能会导致死锁
		假设 channelSize = 3，size = 5，代码执行顺序如下：
			子协程发送前 3 个元素：
				子协程将元素 0、1、2 放入缓冲区（缓冲区容量 3，已满）。
			子协程阻塞在第 4 个元素：
				当子协程尝试发送元素 3 时，缓冲区已满，发送操作被阻塞。此时子协程暂停执行，等待缓冲区有空间（即有接收操作）。
			主协程休眠 2 秒后开始接收：
				主协程从通道接收元素。假设主协程接收速度很快，一次性接收了 3 个元素（0、1、2），此时缓冲区变为空。
			关键点：子协程是否恢复执行？
				如果主协程在接收 3 个元素后立即继续接收第 4 个元素：
					此时子协程会恢复执行，将元素 3 放入缓冲区并继续发送元素 4。程序不会死锁。
				如果主协程在接收 3 个元素后（缓冲区为空），没有立即发起下一次接收：
					子协程仍处于阻塞状态，因为Go 的通道发送操作不会自动 “检测” 缓冲区是否有空间，而是需要接收操作主动触发。此时主协程和子协程可能陷入双向阻塞：
					子协程等待主协程接收元素以腾出缓冲区空间；
					主协程可能在等待子协程发送更多元素（例如，主协程执行 num := <-ch 时，通道中没有元素
*/
func ChannelTestCache(channelSize, size int) {
	// 创建一个有channelSize个缓冲区的channel，且channel中元素类型为int
	ch := make(chan int, channelSize)
	go func() {
		defer fmt.Println("子goroutine结束")
		fmt.Println("子goroutine正在运行")
		for i := 0; i < size; i++ {
			ch <- i
			fmt.Println("子goroutine正在发送元素=", i, "len(ch)=", len(ch), "cap(ch)=", cap(ch))
		}
	}()
	time.Sleep(time.Second * 2)
	for i := 0; i < size; i++ {
		num := <-ch
		fmt.Println("num=", num)
	}
	fmt.Println("ChannelTestCache goroutine 结束")
}

func ChannelTestClose() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// 如果没有数据发送就关闭通道，防止主goroutine阻塞造成死锁
		close(ch)
	}()

	for {
		if data, ok := <-ch; ok {
			fmt.Println(data)
		} else {
			break
		}
	}
	fmt.Println("main goroutine 结束")
}

func ChannelTestRange() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// 如果没有数据发送就关闭通道，防止主goroutine阻塞造成死锁
		close(ch)
	}()

	for data := range ch {
		fmt.Println(data)
	}
	fmt.Println("main goroutine 结束")
}

// 多个通道的操作，需要使用select
func ChannelWrite(ch, quit chan int) {
	x, y := 1, 1

	for {
		select {
		//ch通道可写时，执行该case
		case ch <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func ChannelTestSelect() {
	ch := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-ch)
		}
		quit <- 0
	}()
	ChannelWrite(ch, quit)
}

func main() {
	// ChannelTestNoCache()
	// ChannelTestCache(3, 3)
	ChannelTestCache(3, 10)
	// ChannelTestClose()
	// ChannelTestRange()
	// ChannelTestSelect()
}
