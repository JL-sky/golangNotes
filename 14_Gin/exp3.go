package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type student struct {
	Name string
	Age  int8
}

// 定义一个函数，用于加载HTML模板
func withHtmlTemplate(r *gin.Engine) {
	// curl localhost:8080/arr
	// 加载templates目录下的所有HTML文件
	r.LoadHTMLGlob("templates/*")

	// 创建两个student结构体实例
	stu1 := &student{Name: "Geektutu", Age: 20}
	stu2 := &student{Name: "Jack", Age: 22}
	// 定义一个GET请求的路由，当访问/arr时，返回arr.tmpl模板，并将title和stuArr作为参数传递给模板
	r.GET("/arr", func(c *gin.Context) {
		c.HTML(http.StatusOK, "demo.html", gin.H{
			"title":  "Gin",
			"stuArr": [2]*student{stu1, stu2},
		})
	})
}
