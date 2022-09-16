package sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestRwLock(t *testing.T) {
	lock := sync.RWMutex{}
	lock.RLock()
	// defer lock.RUnlock() // 不能unlock两次
	lock.RUnlock()

	lock.Lock()
	defer lock.Unlock()

	fmt.Println("lock success")
}
