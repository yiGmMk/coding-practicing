package trace

import (
	"fmt"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"testing"
)

// 一个简单的例子去演示如何在代码中使用trace包
// 获取到trace.out后通过go tool trace trace.out即可在web页面查看
func TestTrace(t *testing.T) {
	// 1. 创建trace持久化的文件句柄
	f, err := os.Create("trace.out")
	if err != nil {
		log.Fatalf("failed to create trace output file: %v", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("failed to close trace file: %v", err)
		}
	}()

	// 2. trace绑定文件句柄
	if err := trace.Start(f); err != nil {
		log.Fatalf("failed to start trace: %v", err)
	}
	defer trace.Stop()

	// 下面就是你的监控的程序
	// 我简单写了一个文件读写
	var wg sync.WaitGroup
	wg.Add(2)

	// 一个协程用来读文件
	go func() {
		defer wg.Done()
		content, err := os.ReadFile(`mxc.txt`)
		fmt.Println(content, err)
	}()

	// 写文件协程
	go func() {
		defer wg.Done()
		fmt.Println("hello trace")
	}()
	wg.Wait()
}
