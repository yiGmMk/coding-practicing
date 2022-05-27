package concurrency

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var lock = sync.Mutex{}
var i int

func Range() {
	defer func() {
		if r := recover(); r != nil {
			log.Println("panic", r)
		}
	}()
	done := make(chan struct{}, 50000)
	data := make(chan time.Time, 50000)
	out := make(chan struct{}, 50000)

	go func() {
		for i := 0; i < 40000; i++ {
			data <- time.Now()
		}
	}()
	DealWith(done, data, out)
	DealWith(done, data, out)
	DealWith(done, data, out)
	DealWith(done, data, out)
	for range out {
		i++
		if i == 40000 {
			close(done)
			break
		}
	}
	time.Sleep(time.Second)
}

func DealWith(done chan struct{}, data <-chan time.Time, out chan<- struct{}) {
	go func() {
		defer func() {
			if r := recover(); r != nil {
				log.Println("panic", r)
			}
		}()
		for {
			select {
			case val := <-data:
				//resp, _ := resty.New().R().Get("http://www.baidu.com")
				fmt.Printf("第%d:%s,%+v\n", i, val.Format("2006-01-02 15:04:05"), i)
				time.Sleep(time.Millisecond)
				out <- struct{}{}
			case <-done:
				fmt.Println("----------------go done--------------------")
				return
			}
		}
	}()
}
