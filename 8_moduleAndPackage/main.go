package main

import (
	"fmt"

	"github.com/moduleExp/Hello"
	utils "github.com/string_utils"

	"github.com/test"
)

func main() {
	//使用同级目录下的go文件,同级目录的go文件必须处在同一个包中
	fmt.Println(HelloWorld())
	//使用不同目录下的go文件，同时不使用mod文件，导入时导入：模块名/包名
	Hello.HelloWorld("zhangsan")
	//使用不同目录下的go文件，使用mod文件，但是mod中定义的模块名与目录名不同，导入时导入：模块并使用别名
	str := "zhangsan"
	res := utils.StringReverse(str)
	fmt.Println(res)

	rr := utils.AddAndGreet(1, 2)
	fmt.Println(rr)
	//使用不同目录下的go文件，使用mod文件，同时mod中定义的模块名与目录名相同，导入时导入：模块名
	test.TestAdd()
}
