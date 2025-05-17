package main

import (
	"fmt"
	"time"
)

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

func ChannelTestCache(channelSize, size int) {
	// 创建一个有3个缓冲区的channel，且channel中元素类型为int
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
	// ChannelTestCache(3, 4)
	// ChannelTestClose()
	// ChannelTestRange()
	ChannelTestSelect()
}
