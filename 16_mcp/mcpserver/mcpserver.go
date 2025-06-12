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
	s := server.NewMCPServer("CalculatorServer", "1.0.0")

	// 添加工具
	calculatorTool := mcp.NewTool("calculate",
		mcp.WithDescription("执行基本的算术运算"),
		mcp.WithString("operation",
			mcp.Required(),
			mcp.Description("要执行的算术运算类型"),
			mcp.Enum("multiply", "divide"),
		),
		mcp.WithNumber("x",
			mcp.Required(),
			mcp.Description("第一个数字"),
		),
		mcp.WithNumber("y",
			mcp.Required(),
			mcp.Description("第二个数字"),
		),
	)

	// 将工具添加到服务器
	s.AddTool(calculatorTool, func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
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
	})

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
