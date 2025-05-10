package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Id   int
	Name string
	Age  int
}

func (u User) Call() {
	fmt.Println("user is called...")
	fmt.Printf("%v\n", u)
}

func DoFiledAndMethod(input interface{}) {
	inputType := reflect.TypeOf(input)
	fmt.Println("type:", inputType)
	inputValue := reflect.ValueOf(input)
	fmt.Println("value:", inputValue)
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)
		value := inputValue.Field(i)
		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
	}
}

func main() {
	user := User{1, "Tom", 18}
	DoFiledAndMethod(user)
}
