package zlog

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

// zap 日志测试,使用runtime.Caller获取程序调用栈和源代码文件相关信息
// zap.AddCaller,库本身支持的获取行号,文件路径的方案
func TestZap(t *testing.T) {
	tempFile, err := os.CreateTemp("", "zlog-*.txt")
	if err != nil {
		t.Fatalf("create tmp file:%v", err)
	}
	path := tempFile.Name()
	defer func() {
		_ = tempFile.Close()
		_ = os.RemoveAll(path)
	}()
	t.Run("without caller", func(t *testing.T) {
		ZlogInit(path, true)

		Info("1", zap.String("tag", "dev"), zap.Int("num", 123))
		Debug("2", zap.String("tag", "dev"), zap.Int("num", 123))
		Error("3", zap.String("tag", "dev"), zap.Int("num", 123))
		Warn("4", zap.String("tag", "dev"), zap.Int("num", 123))
	})

	t.Run("caller", func(t *testing.T) {
		logger.WithOptions(zap.AddCaller()).Info("5", zap.String("tag", "dev"), zap.Int("num", 123))
	})

	t.Run("rotate", func(t *testing.T) {
		i := 0
		for {
			logger.Info("test rotate", zap.Int("id", i), zap.String("uuid", uuid.NewString()))
			i++
			if i == 1000 {
				break
			}
		}
	})
}

func TestCaller(t *testing.T) {
	call := func() []zap.Field {
		return getCallerInfoForLog()
	}

	fmt.Println(call())
}
