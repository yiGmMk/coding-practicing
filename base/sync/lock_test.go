package sync

import (
	"fmt"
	"sync"
	"testing"
)

var a int

func TestRwLock(t *testing.T) {
	lock := sync.RWMutex{}
	lock.RLock()
	// defer lock.RUnlock() // 不能unlock两次
	a++
	lock.RUnlock()

	lock.Lock()
	defer lock.Unlock()
	a--
	fmt.Println("lock success", a)
}
