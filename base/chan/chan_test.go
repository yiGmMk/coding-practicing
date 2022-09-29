package chantest

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
	"testing"
)

type Pair struct {
	l, r int32
}

type Pairs []Pair

func (ps Pairs) Len() int { return len(ps) }

func (ps Pairs) Less(i, j int) bool  { return ps[i].l < ps[j].l }
func (ps Pairs) Next(i, j int) bool  { return ps[i].r < ps[j].l }
func (ps Pairs) Swap(i, j int)       { ps[i], ps[j] = ps[j], ps[i] }
func (ps Pairs) Equal(i, j int) bool { return ps[i].l == ps[j].l && ps[i].r == ps[j].r }

func getThreeNonOverlappingIntervals(starting []int32, ending []int32) int64 {
	// Write your code here
	// n*n*n
	if len(starting) < 3 {
		return 0
	}
	var ps Pairs
	for i := 0; i < len(starting) && i < len(ending); i++ {
		ps = append(ps, Pair{l: starting[i], r: ending[i]})
	}
	sort.Sort(ps)

	var res int64
	// overlap := make(map[int]map[int]struct{})
	// set := func(i, j int) {
	// 	if _, ok := overlap[i]; !ok {
	// 		overlap[i] = make(map[int]struct{})
	// 		overlap[i][j] = struct{}{}
	// 	} else {
	// 		overlap[i][j] = struct{}{}
	// 	}
	// 	if _, ok := overlap[j]; !ok {
	// 		overlap[j] = make(map[int]struct{})
	// 		overlap[j][i] = struct{}{}
	// 	} else {
	// 		overlap[j][i] = struct{}{}
	// 	}
	// }
	// isOverlap := func(i, j int) bool {
	// 	if _, ok := overlap[i]; !ok {
	// 		return false
	// 	}
	// 	_, ok := overlap[i][j]
	// 	return ok
	// }
	for i := 0; i < ps.Len(); i++ {
		if i > 0 && ps.Equal(i, i-1) {
			continue
		}
		k := ps.Len() - 1
		for j := i + 1; j < ps.Len(); j++ {
			if !ps.Next(i, j) {
				continue
			}
			if j > i+1 && ps.Equal(j, j-1) {
				continue
			}
			for j < k && ps.Next(j, k) {
				res++
			}
			// sort.Search(ps.Len(), func(i int) bool { return })
		}
	}
	return res
}

/*
 * Complete the 'Server' function below and missing types and global variables.
 *
 * The function is void.
 */
type in struct {
	val               int32
	oddChan, evenChan chan int32
}

var (
	serverChan = make(chan in, 1000)
)

func Server() {
	for v := range serverChan {
		if v.val%2 == 0 {
			v.evenChan <- v.val
		} else {
			v.oddChan <- v.val
		}
	}
}
func TestChan(t *testing.T) {
	stdout := os.Stdout
	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	var arr []int32

	for i := 0; i < rand.Int()%1000; i++ {
		arr = append(arr, int32(rand.Int()))
	}

	oddChan := make(chan int32)
	evenChan := make(chan int32)
	for idx := 0; idx < len(arr); idx++ {
		i := idx
		serverChan <- in{arr[i], oddChan, evenChan}
	}
	odds, evens := []int32{}, []int32{}
	wg := &sync.WaitGroup{}
	wg.Add(len(arr))
	go func() {
		for newOdd := range oddChan {
			odds = append(odds, newOdd)
			wg.Done()
		}
	}()
	go func() {
		for newEven := range evenChan {
			evens = append(evens, newEven)
			wg.Done()
		}
	}()
	go Server()
	wg.Wait()

	fmt.Fprintln(writer, "odds:")
	for _, resultItem := range odds {
		fmt.Fprintf(writer, "%d", resultItem)
		fmt.Fprintf(writer, "\n")
	}

	fmt.Fprintln(writer, "evens:")
	for i, resultItem := range evens {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(evens)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
