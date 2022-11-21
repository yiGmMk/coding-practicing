package base

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	consume := func(c <-chan bool) bool {
		timer := time.NewTimer(time.Second * 5)
		defer timer.Stop()
		select {
		case b := <-c:
			if b == false {
				log.Printf("recv false, continue")
				return true
			}
			log.Printf("recv true, return")
			return false
		case <-timer.C:
			log.Printf("timer expired")
			return true
		}
	}
	c := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)

	// 生产者
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1)
			c <- false
		}

		time.Sleep(time.Second * 1)
		c <- true
		wg.Done()
	}()
	// 消费者
	go func() {
		for {
			if b := consume(c); !b {
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
}

func TestTimerReset(t *testing.T) {
	consume := func(c <-chan bool, timer *time.Timer) bool {
		// Stop prevents the Timer from firing.
		// It returns true if the call stops the timer, false if the timer has already
		// expired or been stopped.
		if !timer.Stop() {
			select {
			case <-timer.C: //防止堵塞
				fmt.Println("time.C len:", len(timer.C))
			default:
			}

			// <-timer.C
		}
		timer.Reset(5 * time.Second)
		select {
		case b := <-c:
			if b == false {
				log.Printf("recv false, continue")
				return true
			}
			log.Printf("recv true, return")
			return false
		case <-timer.C:
			log.Printf("timer expired")
			return true
		}
	}

	c := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Second * 1) //
			c <- false
		}
		time.Sleep(time.Second * 1)
		c <- true
		wg.Done()
	}()
	go func() {
		timer := time.NewTimer(time.Second * 5)
		for {
			if b := consume(c, timer); !b {
				wg.Done()
				return
			}
		}
	}()
	wg.Wait()
}

// time timer使用
func TestTimerSince(t *testing.T) {
	start := time.Now()
	total := time.Microsecond * 1000
	timer := time.NewTimer(total)
	defer timer.Stop()
LOOP:
	for {
		duration := time.Since(start)
		timer.Reset(total - duration)
		log.Printf("duration: %d ms", duration.Microseconds())
		select {
		case <-timer.C:
			log.Println("time tick")
			break LOOP
		default:
			time.Sleep(time.Microsecond * 3)
		}
	}
	log.Println("timer done")
}
