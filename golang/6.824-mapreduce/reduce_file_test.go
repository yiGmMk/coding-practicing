package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/zeromicro/go-zero/core/mr"
)

func TestCheckFile(t *testing.T) {

	has := strings.HasPrefix("mr-out-", "mr-out-1.txt")
	fmt.Println(has)

	dir := "."
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("read dir failed:%+v", err)
	}

	kvs := []KeyVal{}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		if !strings.HasPrefix(file.Name(), "mr-out-") {
			continue
		}

		file, err := os.Open(file.Name())
		if err != nil {
			log.Fatalln("open file failed:", err)
			return
		}
		defer file.Close()
		bf := bufio.NewReader(file)

		for {
			line, err := bf.ReadString('\n')
			if err != nil {
				log.Println("read line failed:", err)
				break
			}
			vs := strings.Split(line, " ")
			if len(vs) != 2 {
				log.Println("invalid line:", line)
				continue
			}
			val, err := strconv.Atoi(strings.Trim(vs[1], "\n"))
			if err != nil {
				log.Println("invalid line:", line, "cann't convert to int", err.Error())
				continue
			}
			kvs = append(kvs, KeyVal{Key: vs[0], Value: val})
		}
	}

	sort.Sort(KeyValByKey(kvs))

	resultFile := "result.txt"
	rf, err := os.OpenFile(resultFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("open result file failed:", err)
	}
	defer rf.Close()
	for _, kv := range kvs {
		fmt.Fprintf(rf, "%v %v\n", kv.Key, kv.Value)
	}
}

func TestMapReduce(t *testing.T) {

	has := strings.HasPrefix("mr-out-", "mr-out-1.txt")
	fmt.Println(has)

	dir := "."
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatalf("read dir failed:%+v", err)
	}

	kvs := []KeyVal{}
	mr.MapReduce(func(source chan<- interface{}) {
		for _, file := range files {
			if file.IsDir() {
				continue
			}
			if strings.HasPrefix(file.Name(), "mr-out-") {
				source <- file.Name()
			}
		}
	}, func(item interface{}, writer mr.Writer, cancel func(error)) {
		fileName, ok := item.(string)
		if !ok || fileName == "" {
			return
		}
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatalln("open file failed:", err)
			return
		}
		defer file.Close()
		bf := bufio.NewReader(file)

		for {
			line, err := bf.ReadString('\n')
			if err != nil {
				log.Println("read line failed:", err)
				break
			}
			vs := strings.Split(line, " ")
			if len(vs) != 2 {
				log.Println("invalid line:", line)
				continue
			}
			val, err := strconv.Atoi(strings.Trim(vs[1], "\n"))
			if err != nil {
				log.Println("invalid line:", line, "cann't convert to int", err.Error())
				continue
			}
			kvs = append(kvs, KeyVal{Key: vs[0], Value: val})
		}
	}, func(pipe <-chan interface{}, writer mr.Writer, cancel func(error)) {

	})
	sort.Sort(KeyValByKey(kvs))

	resultFile := "result.txt"
	rf, err := os.OpenFile(resultFile, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("open result file failed:", err)
	}
	defer rf.Close()
	for _, kv := range kvs {
		fmt.Fprintf(rf, "%v %v\n", kv.Key, kv.Value)
	}
}
