package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"testing"
)

func TestMrc(t *testing.T) {
	files := []string{"./pg-being_ernest.txt"}
	mapReduce(files)
}

func TestMcr2(t *testing.T) {
	n := 10
	files := []string{"./pg-being_ernest.txt"}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Println(err)
			continue
		}
		content, err := io.ReadAll(f)
		if err != nil {
			log.Println(err)
			continue
		}
		kvs := Map(file, string(content))

		sort.Sort(ByKey(kvs))

		kvMap := make(map[int][]KeyValue)
		for _, kv := range kvs {
			key := ihash(kv.Key) % n
			kvMap[key] = append(kvMap[key], kv)
		}
		for k, kvs := range kvMap {
			outf := "mr-out-" + strconv.Itoa(k)
			outFile, err := os.OpenFile(outf, os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				continue
			}
			defer func() {
				_ = outFile.Close()
			}()
			enc := json.NewEncoder(outFile)

			for _, kv := range kvs {
				_ = enc.Encode(&kv)
			}
		}
	}

}

func reduce(filename string) []KeyValue {
	file, err := os.Open(filename)
	if err != nil {
		log.Println(err)
		return []KeyValue{}
	}
	dec := json.NewDecoder(file)
	kvs := []KeyValue{}
	for {
		kv := KeyValue{}
		err := dec.Decode(&kv)
		if err != nil {
			log.Println(err)
			break
		}
		kvs = append(kvs, kv)
	}
	return kvs
}
