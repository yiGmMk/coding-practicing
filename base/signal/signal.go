package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func catchAsyncSignal(c chan os.Signal) {
	for {
		s := <-c
		fmt.Println("收到异步信号:", s)
	}
}

func triggerSyncSignal() {
	time.Sleep(3 * time.Second)
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("恢复panic:", e)
			return
		}
	}()

	var a, b = 1, 0
	fmt.Println(a / b)
}

func main() {
	var wg sync.WaitGroup
	c := make(chan os.Signal, 1)
	signal.Notify(c,
		syscall.SIGFPE,
		syscall.SIGINT,
		// syscall.SIGKILL) //kill -9 强制退出,拦截不了  SA1016: syscall.SIGKILL cannot be trapped (did you mean syscall.SIGTERM?)
		syscall.SIGTERM) // kill -15
	wg.Add(2)
	go func() {
		catchAsyncSignal(c)
		wg.Done()
	}()

	go func() {
		triggerSyncSignal()
		wg.Done()
	}()

	wg.Wait()
}
