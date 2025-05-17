package main

import (
	"bytes"
	"html/template"
	"testing"
	"time"
)

/*
注意事项：
	- 函数名必须以 Benchmark 开头，后面一般跟待测试的函数名
	- 参数为 b *testing.B。
	- 执行基准测试时，需要添加 -bench 参数。
*/

/*
测试执行：
 $  go test -benchmem -bench .
goos: linux
goarch: amd64
pkg: github.com/unitTest/benchMark
cpu: 12th Gen Intel(R) Core(TM) i7-12650H
BenchmarkAdd-4          1000000000               0.4596 ns/op          0 B/op          0 allocs/op
PASS
ok      github.com/unitTest/benchMark   0.651s
*/

/*
基准测试报告每一列值对应的含义如下：

	type BenchmarkResult struct {
			N         int           // 迭代次数
			T         time.Duration // 基准测试花费的时间
			Bytes     int64         // 一次迭代处理的字节数
			MemAllocs uint64        // 总的分配内存的次数
			MemBytes  uint64        // 总的分配内存的字节数
	}
*/
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}

func BenchmarkMul(b *testing.B) {
	time.Sleep(time.Second)
	// 重置计时器，排除 time.Sleep 的影响
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Mul(1, 2)
	}
}

// 使用 RunParallel 测试并发性能
func BenchmarkParallelAdd(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			Add(1, 2)
		}
	})
}

// 基准测试，并行执行
func BenchmarkParallel(b *testing.B) {
	// 创建一个模板，并解析字符串
	templ := template.Must(template.New("test").Parse("Hello, {{.}}!"))
	// 并行执行基准测试
	b.RunParallel(func(pb *testing.PB) {
		// 创建一个缓冲区
		var buf bytes.Buffer
		// 循环执行基准测试，直到 pb.Next() 返回 false
		for pb.Next() {
			// 所有 goroutine 一起，循环一共执行 b.N 次
			// 重置缓冲区
			buf.Reset()
			// 执行模板，并将结果写入缓冲区
			templ.Execute(&buf, "World")
		}
	})
}
