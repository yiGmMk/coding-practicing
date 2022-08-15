package base

import (
	"log"
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"golang.org/x/sync/errgroup"
)

// fan-out
func TestFanOut(t *testing.T) {

}

// fan-in
func TestFanIn(t *testing.T) {

}

func TestChan(t *testing.T) {
	g := errgroup.Group{}
	wg := sync.WaitGroup{}
	ic := make(chan string)
	for i := 0; i < 10; i++ {
		id := i
		wg.Add(1)
		g.Go(func() error {
			for i := 0; i < rand.Intn(100); i++ {
				ic <- strconv.Itoa(id) + ":" + strconv.Itoa(i)
			}
			wg.Done()
			return nil
		})
	}
	g.Go(func() error {
		wg.Wait()
		close(ic)
		return nil
	})
	g.Go(func() error {
		for i := range ic {
			log.Println(i)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		t.Error(err)
	}
	log.Println("end")
}
