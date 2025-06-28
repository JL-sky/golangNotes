package main

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("请提供至少一个参数。")
		return
	}
	// 输出程序名称
	// 遍历并输出所有参数
	a, _ := strconv.ParseFloat(os.Args[1], 64)
	b, _ := strconv.ParseFloat(os.Args[2], 64)
	// 这里的路径是上面编译的mcp-server可执行文件
	mcpClient, err := client.NewStdioMCPClient("/home/cheersj/pro/test/golangNotes/16_mcp/mcpserver/mcpserver", []string{})
	if err != nil {
		panic(err)
	}
	defer mcpClient.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "Client Demo",
		Version: "1.0.0",
	}

	initResult, err := mcpClient.Initialize(ctx, initRequest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("初始化成功，服务器信息: %s %s\n", initResult.ServerInfo.Name, initResult.ServerInfo.Version)

	// 调用工具
	toolRequest := mcp.CallToolRequest{
		Request: mcp.Request{
			Method: "tools/call",
		},
	}
	toolRequest.Params.Name = "calculator"
	toolRequest.Params.Arguments = map[string]any{
		"operation": "multiply",
		"x":         a,
		"y":         b,
	}

	result, err := mcpClient.CallTool(ctx, toolRequest)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%f * %f = %s\n", a, b, result.Content[0].(mcp.TextContent).Text)
}
