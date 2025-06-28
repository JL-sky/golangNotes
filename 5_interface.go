package main

import (
	"fmt"
)

func funcName(a interface{}) string {
	value, ok := a.(string)
	if !ok {
		fmt.Println("Type assertion failed")
		return ""
	}
	fmt.Println("Type assertion successful, the value is: ", value)
	return value
}

func AnyTest(arg any) {
	switch value := arg.(type) {
	case int:
		fmt.Println("The value is an int >>> ", value)
	case string:
		fmt.Println("The value is a string >>> ", value)
	}
}

func InterfaceTest() {
	// var a int = 10
	var a string = "hello"
	funcName(a)
}

func main() {
	// InterfaceTest()
	AnyTest(10)
	AnyTest("hello")
}
