package main

import (
	"fmt"
	"runtime"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("newTask: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 使用显示函数创建子routinue
func goroutineTest1() {
	go newTask()
	i := 0
	for {
		i++
		fmt.Printf("main: %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 使用匿名函数创建子routinue
func goroutineTest2() {
	// 创建匿名函数并作为goroutine执行
	go func() {
		defer fmt.Print("goroutine A end\n")
		func() {
			defer fmt.Print("goroutine B end\n")
			// 调用runtime.Goexit()终止当前goroutine
			runtime.Goexit()
			fmt.Print("goroutine B\n")
		}()
		fmt.Print("goroutine A\n")
	}()

	for {
		time.Sleep(1 * time.Second)
	}
}

// 使用匿名函数创建子routinue
func goroutineTest3() {
	// 创建匿名函数并作为goroutine执行
	go func(a, b int) {
		time.Sleep(1 * time.Second)
		fmt.Printf("goroutine: %d + %d = %d\n", a, b, a+b)
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
func main() {
	// goroutineTest1()
	// goroutineTest2()
	goroutineTest3()
}
