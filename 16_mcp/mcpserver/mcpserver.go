package main

import (
	"context"
	"fmt"
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// 创建一个新的MCPServer实例
	s := server.NewMCPServer(
		"CalculatorServer", //服务器名称
		"1.0.0",            //服务器版本
	)

	// 添加工具
	calculatorTool := mcp.NewTool(
		"calculator", //工具名称",
		mcp.WithDescription("执行基本的算术运算"), //工具描述
		/*以下为参数设置*/
		mcp.WithString("operation", //参数名称
			mcp.Required(), //表示参数必填
			mcp.Description("要执行的算术运算类型"),  //参数描述
			mcp.Enum("multiply", "divide"), //参数的可选值
		),

		mcp.WithNumber("x", //参数名称
			mcp.Required(),           //表示参数必填
			mcp.Description("第一个数字"), //参数描述
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("第二个数字"),
		),
	)

	// 将工具添加到服务器
	s.AddTool(calculatorTool, calculator)

	// 启动SSE服务器
	//sseServer := server.NewSSEServer(s, server.WithBaseURL("http://localhost:8082"))
	//log.Printf("SSE server listening on :8082")
	//if err := sseServer.Start(":8082"); err != nil {
	//	log.Fatalf("Server error: %v", err)
	//}

	// 启动基于 stdio 的服务器
	if err := server.ServeStdio(s); err != nil {
		log.Printf("Server error: %v\n", err)
	}
}

// 定义一个业务处理函数，用于处理工具的请求
func calculator(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// 获取请求中的operation参数
	op, err := request.RequireString("operation")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	// 获取请求中的x参数
	x, err := request.RequireFloat("x")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	// 获取请求中的y参数
	y, err := request.RequireFloat("y")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	// 根据operation参数执行相应的算术运算
	var result float64
	switch op {
	case "add":
		result = x + y
	case "subtract":
		result = x - y
	case "multiply":
		result = x * y
	case "divide":
		if y == 0 {
			return mcp.NewToolResultError("cannot divide by zero"), nil
		}
		result = x / y
	}

	// 返回结果
	return mcp.NewToolResultText(fmt.Sprintf("%.2f", result)), nil

}
