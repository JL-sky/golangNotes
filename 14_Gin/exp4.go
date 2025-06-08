package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func withMiddleware(r *gin.Engine) {
	// 全局中间件：记录请求耗时
	r.Use(recordTime)

	// 单个路由中间件：权限验证
	r.GET("/hello", verifyToken, sayHello)

	// 路由组中间件
	apiGroup := r.Group("/api")
	apiGroup.Use(logGroupAccess)
	{
		apiGroup.GET("/user", getUserInfo)
		apiGroup.GET("/password", getPasswd)
	}
}

// 记录请求耗时的中间件
// 记录请求处理时间
func recordTime(c *gin.Context) {
	// 记录请求开始时间
	start := time.Now()
	// 继续处理请求
	c.Next()
	// 计算请求处理时间
	cost := time.Since(start)
	// 打印请求处理时间
	log.Printf("[GLOBAL] %s %s 耗时: %v",
		c.Request.Method, c.Request.URL.Path, cost)
}

// 验证Token的中间件
func verifyToken(c *gin.Context) {
	if token := c.Query("token"); token != "test" {
		c.AbortWithStatusJSON(401, gin.H{"msg": "权限不足"})
		return
	}
	c.Next()
}

// /hello路由的处理函数
func sayHello(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Hello, Middleware!"})
}

// 记录路由组访问的中间件
func logGroupAccess(c *gin.Context) {
	log.Println("[GROUP] 进入API组")
	c.Next()
}

// /api/user路由的处理函数
func getUserInfo(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "User API"})
}

// /api/password路由的处理函数
func getPasswd(c *gin.Context) {
	c.JSON(200, gin.H{"msg": "Password API"})
}
