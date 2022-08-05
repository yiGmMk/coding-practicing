package base

import (
	"log"
	"os"
	"testing"
)

// 文件保存,6.824中提到的优化措施
// 文件内容先保存到临时文件中,确认后再修改文件名
// 通过ioutil.TempFile();os.Rename()实现
func TestFileSave(t *testing.T) {
	dir, file := ".", "test.txt"
	// 创建临时文件
	f, err := os.CreateTemp(dir, file)
	if err != nil {
		log.Panicf("create temp file failed,%v", err)
		return
	}
	defer func() {
		_ = os.Remove(f.Name())
	}()
	err = os.Rename(f.Name(), file)
	if err != nil {
		log.Panicf("rename file failed,%v", err)
		return
	}
	defer func() {
		_ = os.Remove(file)
	}()
}
