package main

import (
	"errors"
	"fmt"
	"os"
)

func fileTest() {
	_, err := os.ReadFile("test.txt")
	if err != nil {
		panic(err)
	}
}

func hello(name string) error {
	if len(name) == 0 {
		return errors.New("name is empty")
	}
	fmt.Println("Hello, " + name)
	return nil
}

/*
1.在 get 函数中，使用 defer 定义了异常处理的函数，
	在协程退出前，会执行完 defer 挂载的任务。
	因此如果触发了 panic，控制权就交给了 defer。
2.在 defer 的处理逻辑中，使用 recover，使程序恢复正常，
	并且将返回值设置为 -1，在这里也可以不处理返回值，
	如果不处理返回值，返回值将被置为默认值 0
*/

func get(index int) (ret int) {
	defer func() {
		fmt.Println("deferred call")
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
			ret = -1
		}
	}()
	arr := [3]int{1, 2, 3}
	return arr[index]
}

func test() {
	// fileTest()
	/*
			err := hello("world")
		if err != nil {
			fmt.Println(err)
		}
		err = hello("")
		if err != nil {
			fmt.Println(err)
		}
	*/
	fmt.Println(get(5))
}

func main() {
	test()
}
