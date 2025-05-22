package main

import (
	"context"
	"fmt"
	"time"
)

// 客户端接口
type Client interface {
	GetNamespace() string
	GetServiceName() string
	GetSet() string
	GetTimeout() int
}

type Option func(o *Options)

// 模拟registry.Node类型
type Node struct {
	Address string
}

// 客户端配置
type Options struct {
	ServiceName       string
	Namespace         string
	SetName           string
	Protocol          string
	Network           string
	SerializationType int
	Timeout           time.Duration
	SelectorNode      *Node
	RetryTimes        int
	endpoint          string
}

// 客户端构建器
type ClientBuilder struct {
	opts Options
}

// 创建新构建器并设置默认值
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		opts: Options{
			Protocol:   "trpc",
			Network:    "tcp4",
			Timeout:    5 * time.Second,
			RetryTimes: 3,
		},
	}
}

// 应用选项
func (b *ClientBuilder) Apply(opts []Option) {
	for _, opt := range opts {
		opt(&b.opts)
	}
}

// 构建客户端
func (b *ClientBuilder) Build() Client {
	if b.opts.ServiceName == "" {
		panic("service name is required")
	}
	return &ClientImpl{ // 假设 ClientImpl 实现了 Client 接口
		namespace:   b.opts.Namespace,
		serviceName: b.opts.ServiceName,
		setName:     b.opts.SetName,
		timeout:     int(b.opts.Timeout / time.Millisecond),
	}
}

// 客户端实现（具名结构体方案）
type ClientImpl struct {
	namespace   string
	serviceName string
	setName     string
	timeout     int
}

func (c *ClientImpl) GetNamespace() string   { return c.namespace }
func (c *ClientImpl) GetServiceName() string { return c.serviceName }
func (c *ClientImpl) GetSet() string         { return c.setName }
func (c *ClientImpl) GetTimeout() int        { return c.timeout }

// 选项设置函数
func WithServiceName(s string) Option {
	return func(o *Options) {
		o.ServiceName = s
		o.endpoint = s
	}
}

func WithNamespace(s string) Option {
	return func(o *Options) {
		o.Namespace = s
	}
}

func WithSetName(s string) Option {
	return func(o *Options) {
		o.SetName = s
	}
}

// newOpts 函数
func newOpts(ctx context.Context, n *Node, o Client) []Option {
	namespace := o.GetNamespace()
	if namespace == "Development" && "Production" == "Production" { // 示例条件
		namespace = "Production"
	}
	return []Option{
		WithServiceName(o.GetServiceName()),
		WithNamespace(namespace),
		WithSetName(o.GetSet()),
		// 其他选项...
	}
}

func main() {
	ctx := context.Background()
	node := &Node{Address: "127.0.0.1:8080"}

	// 使用具名结构体实现 Client 接口
	clientConfig := &ClientImpl{
		namespace:   "Development",
		serviceName: "user-service",
		setName:     "main",
		timeout:     5000,
	}

	// 创建选项列表
	options := newOpts(ctx, node, clientConfig)

	// 应用选项到构建器
	builder := NewClientBuilder()
	builder.Apply(options)

	// 构建客户端
	client := builder.Build()

	// 验证结果
	fmt.Printf("ServiceName: %s\n", builder.opts.ServiceName)   // user-service
	fmt.Printf("Namespace: %s\n", builder.opts.Namespace)       // Production
	fmt.Printf("SetName: %s\n", builder.opts.SetName)           // main
	fmt.Printf("Client Namespace: %s\n", client.GetNamespace()) // Production
}
