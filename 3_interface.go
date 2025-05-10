package main

import "fmt"

func funcName(a interface{}) string {
	value, ok := a.(string)
	if !ok {
		fmt.Println("Type assertion failed")
		return ""
	}
	fmt.Println("Type assertion successful, the value is: ", value)
	return value
}

func InterfaceTest() {
	// var a int = 10
	var a string = "hello"
	funcName(a)
}

func main() {
	InterfaceTest()
}
