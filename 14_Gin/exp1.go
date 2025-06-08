package main

// 安装：go install github.com/gin-gonic/gin@latest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 无参数get请求
func withNoParm(r *gin.Engine) {
	// curl localhost:8080
	// 声明了一个路由，告诉 Gin 什么样的URL能触发传入的函数，这个函数返回我们想要显示在用户浏览器中的信息
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Who are you?")
	})
}

// 带参数get请求
func withParm(r *gin.Engine) {
	// curl localhost:8080/user/张三
	r.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
}

// 带参数get请求
func withQuery(r *gin.Engine) {
	// curl "localhost:8080/users?name=jack&role=stu"
	r.GET("/users", func(c *gin.Context) {
		// name参数必须要有，否则报错
		name := c.Query("name")
		// role参数可选，如果没有，则使用默认值
		role := c.DefaultQuery("role", "teacher")
		c.String(http.StatusOK, "%s is a %s", name, role)
	})
}

// 带参数post请求
func withPost(r *gin.Engine) {
	// curl localhost:8080/form -X POST -d 'username=jack&password=123'
	r.POST("/form", func(c *gin.Context) {
		// username参数必须要有，否则报错
		username := c.PostForm("username")
		// password参数可选，如果没有，则使用默认值
		password := c.DefaultPostForm("password", "000000")

		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
		})
	})
}

// 带参数post请求
func withPostAndQuery(r *gin.Engine) {
	// curl "http://localhost:8080/posts?id=9876&page=7"  -X POST -d 'username=geektutu&password=1234'
	r.POST("/posts", func(c *gin.Context) {
		// 获取查询参数id
		id := c.Query("id")
		// 获取查询参数page，如果没有则默认为0
		page := c.DefaultQuery("page", "0")
		// 获取POST表单参数username
		username := c.PostForm("username")
		// 获取POST表单参数password，如果没有则默认为000000
		password := c.DefaultPostForm("username", "000000")

		// 返回JSON格式的响应
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"page":     page,
			"username": username,
			"password": password,
		})
	})
}

func withPostMap(r *gin.Engine) {
	// curl -g "http://localhost:8080/post?ids[Jack]=001&ids[Tom]=002" -X POST -d 'names[a]=Sam&names[b]=David'
	r.POST("/post", func(c *gin.Context) {
		// 获取URL中的参数ids，返回一个map
		ids := c.QueryMap("ids")
		// 获取POST请求中的参数names，返回一个map
		names := c.PostFormMap("names")

		// 返回JSON格式的数据，包含ids和names
		c.JSON(http.StatusOK, gin.H{
			"ids":   ids,
			"names": names,
		})
	})
}

// 定义一个函数，用于设置重定向
func withRedirect(r *gin.Engine) {
	// curl -i http://localhost:8080/redirect
	// 当访问/redirect时，重定向到/index
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/index")
	})

	// curl -i http://localhost:8080/goindex
	// 当访问/goindex时，将路径设置为根路径，并处理请求
	r.GET("/goindex", func(c *gin.Context) {
		c.Request.URL.Path = "/"
		r.HandleContext(c)
	})
}

func withGroupRoute(r *gin.Engine) {
	defaultHandler := func(c *gin.Context) {
		// 返回一个JSON响应，其中包含请求的完整路径
		c.JSON(http.StatusOK, gin.H{
			"path": c.FullPath(),
		})
	}

	// group routes 分组路由
	// curl http://localhost:8080/v1/posts
	// group: v1
	v1 := r.Group("/v1")
	{
		v1.GET("/posts", defaultHandler)
		v1.GET("/series", defaultHandler)
	}
	// curl http://localhost:8080/v2/posts
	// group: v2
	v2 := r.Group("/v2")
	{
		v2.GET("/posts", defaultHandler)
		v2.GET("/series", defaultHandler)
	}

}
