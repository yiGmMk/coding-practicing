package base

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"
	"unicode"

	"github.com/samber/lo"
	"github.com/sergi/go-diff/diffmatchpatch"
)

type KeyValue struct {
	Key   string
	Value string
}

type ByKey []KeyValue

// for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }

func TestCreateOrAppendFile(t *testing.T) {
	fileName := "./test.txt"
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("TestCreateOrAppendFile panic,recover:%+v", r)
		}
		os.RemoveAll(fileName)
	}()
	err := CreateOrAppendFile(fileName, []string{time.Now().Format("2006-01-02 13:04:05"), "hello", "world", "\n"})
	log.Println(lo.Must(err, nil))

	err = CreateOrAppendFile(fileName, []string{time.Now().Format("2006-01-02 13:04:05"), "hello", "world"})
	log.Println(lo.Must(err, nil))
}

// 文件不存在则创建,存在则追加
// create file if not exist, else append to tail
func CreateOrAppendFile(filename string, contents []string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	//file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("open file error:%w", err)
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	for _, content := range contents {
		err = enc.Encode(content)
		if err != nil {
			return fmt.Errorf("encode error:%w", err)
		}
	}
	return nil
}

func TestReduce(t *testing.T) {
	dir := "./file"
	es, err := os.ReadDir(dir)
	if err != nil {
		log.Println(err)
	}

	files := []string{}
	for _, e := range es {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if name == "" {
			continue
		}
		if !unicode.IsDigit(rune(name[0])) {
			continue
		}
		files = append(files, dir+"/"+name)
	}

	// 输出结果
	outPut := "./out.txt"
	kvs, err := Reduce(files, outPut)
	if err != nil {
		log.Println(err)
	}
	if len(kvs) <= 0 {
		return
	}
	outFile, err := os.OpenFile(outPut, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}
	defer func() {
		outFile.Close()
		os.RemoveAll(outPut)
	}()

	writeOutput2File(outFile, kvs)

	// diff file
	f1, f2 := outPut, "./file/mr-correct-file"

	err = fileDiff(f1, f2)
	if err != nil {
		log.Println("diff file failed:", err)
	}
	DiffFile(f1, f2)
}

func reducef(key string, values []string) string {
	// return the number of occurrences of this word.
	return strconv.Itoa(len(values))
}

func writeOutput2File(file *os.File, kvs []KeyValue) {
	i := 0
	k := 0
	for i < len(kvs) {
		j := i + 1
		for j < len(kvs) && kvs[j].Key == kvs[i].Key {
			j++
		}
		values := []string{}
		for k := i; k < j; k++ {
			values = append(values, kvs[k].Value)
		}
		output := reducef(kvs[i].Key, values)

		// this is the correct format for each line of Reduce output.
		fmt.Fprintf(file, "%v %v\n", kvs[i].Key, output)

		i = j
		k++
	}
}

// reduce file
func Reduce(files []string, outfile string) ([]KeyValue, error) {
	kvs := []KeyValue{}
	if len(files) <= 0 {
		return kvs, nil
	}

	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			return nil, fmt.Errorf("open file failed,%w", err)
		}
		defer f.Close()

		dec := json.NewDecoder(f)
		for {
			var kv KeyValue
			err := dec.Decode(&kv)
			if err != nil {
				break
			}
			kvs = append(kvs, kv)
		}
	}

	if len(kvs) <= 0 {
		return kvs, nil
	}

	sort.Sort(ByKey(kvs))

	return kvs, nil
}

// diff two file to check if is the same
func fileDiff(f1, f2 string) error {
	c1, err := ioutil.ReadFile(f1)
	if err != nil {
		return err
	}
	c2, err := ioutil.ReadFile(f2)
	if err != nil {
		return err
	}
	dmp := diffmatchpatch.New()

	diffs := dmp.DiffMain(string(c1), string(c2), false)

	log.Println("file differences:", dmp.DiffPrettyText(diffs))
	return nil
}
