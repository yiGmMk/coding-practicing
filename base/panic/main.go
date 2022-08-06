package main

import (
	"log"
	"time"
)

// panic can only be recovered in the deferred function
// and in a goroutine it can recovered by deferred function in the same goroutine
func main() {
	go func() {
		defer func() {
			log.Println("1.start recover out of goroutine 2")
			if err := recover(); err != nil {
				log.Println("2.panic goroutine 1", err)
			}
			log.Println("3.finish def in goroutine 1")
		}()
		go func() {
			defer func() {
				if r := recover(); r != nil {
					log.Println("4.panic in goroutine 2", r)
				}
			}()
			log.Println("5.panic before recover in goroutine 2")
			panic(time.Now())
		}()

		time.Sleep(time.Second * 1)
		log.Println("6.finish goroutine 1")
	}()
	log.Println("7.befer sleep in main")
	time.Sleep(time.Second * 5)
	log.Println("8.finish in main")
}

/*output:
2022/07/27 10:11:27 7.befer sleep in main
2022/07/27 10:11:27 5.panic before recover in goroutine 2
2022/07/27 10:11:27 4.panic in goroutine 2 2022-07-27 10:11:27.362826084 +0800 CST m=+0.000354330
2022/07/27 10:11:28 6.finish goroutine 1
2022/07/27 10:11:28 1.start recover out of goroutine 2
2022/07/27 10:11:28 3.finish def in goroutine 1
2022/07/27 10:11:32 8.finish in main
*/
