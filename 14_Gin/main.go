package main

import "github.com/gin-gonic/gin"

func main() {
	// 生成了一个实例，这个实例即 WSGI 应用程序。
	r := gin.Default()

	// 路由
	withNoParm(r)
	withParm(r)
	withQuery(r)
	withPost(r)
	withPostAndQuery(r)
	withPostMap(r)
	withRedirect(r)
	withGroupRoute(r)

	// 上传文件
	uploadSingleFile(r)
	uploadMultipleFiles(r)

	// HTML模板
	withHtmlTemplate(r)

	// 中间件
	withMiddleware(r)

	// 让应用运行在本地服务器上，默认监听端口是 _8080_，可以传入参数设置端口
	r.Run() // listen and serve on 0.0.0.0:8080
	// r.Run(":8081")
}
