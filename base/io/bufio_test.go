package base

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestBufio(t *testing.T) {
	t.Parallel()
	file := "./tmp/bufio.txt"

	t.Run("write", func(t *testing.T) {
		f, err := os.OpenFile(file, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
		if err != nil {
			t.Fatal(err)
		}
		defer func() {
			_ = f.Sync()
			f.Close()
		}()

		data := []byte("practicing,go \n") // ("I love golang!\n") //
		// 带缓冲的io
		bio := bufio.NewWriterSize(f, 32)

		// 将15字节写入bio缓冲区，缓冲区缓存15字节，bufio.txt中内容仍为空
		bio.Write(data)

		// 将15字节写入bio缓冲区，缓冲区缓存30字节，bufio.txt中内容仍为空
		bio.Write(data)

		// 将15字节写入bio缓冲区后，bufio将32字节写入bufio.txt，
		// bio缓冲区中仍然缓存（15*3-32）字节
		bio.Write(data)

		// 将bio缓冲区中的所有缓存数据均写入bufio.txt
		bio.Flush()

		fmt.Println("write end")
	})

	time.Sleep(time.Second)
	t.Run("read", func(t *testing.T) {
		defer func() {
			os.RemoveAll(file)
		}()

		f, err := os.Open(file)
		if err != nil {
			t.Error(err)
		}

		bio := bufio.NewReaderSize(f, 64)
		fmt.Printf("当前缓冲区可读数据:%d字节\n", bio.Buffered())

		var i int = 1
		for {
			data := make([]byte, 15)
			n, err := bio.Read(data)
			if err == io.EOF {
				fmt.Printf("第%d次读数据,到达文件尾\n", i)
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			fmt.Printf("第%d次,读出数据:%q,长度:%d,当前缓冲区可读数据:%d字节\n", i, data, n, bio.Buffered())
			fmt.Printf("当前缓冲区可读数据:%d字节\n", bio.Buffered())
			i++
		}
	})
}

// 当前缓冲区可读数据:0字节
// 第1次,读出数据:"practicing,go \n",长度:15,当前缓冲区可读数据:30字节
// 当前缓冲区可读数据:30字节
// 第2次,读出数据:"practicing,go \n",长度:15,当前缓冲区可读数据:15字节
// 当前缓冲区可读数据:15字节
// 第3次,读出数据:"practicing,go \n",长度:15,当前缓冲区可读数据:0字节
// 当前缓冲区可读数据:0字节
// 第4次读数据,到达文件尾

//第一次读出15字节数据后，当前缓冲区缓存数据数居然是30字节。
//也就是说第一次bufio.Reader.Read操作实际上从文件中读取了45字节，其中前15字节数据通过字节切片传递出来，
//剩余的30字节则缓存在bio维护的内部缓冲区中了。第二次、第三次读操作均为从该缓冲区中读取数据，不会触发文件I/O操作了。

func TestZipWithBuf(t *testing.T) {
	file := "./tmp/zip.gz"
	defer func() {
		os.RemoveAll(file)
	}()
	t.Run("zip", func(t *testing.T) {
		f, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()

		z := gzip.NewWriter(f)
		defer z.Close() // zw.Close方法调用会将压缩变换后的数据刷新到文件实例中

		_, err = z.Write([]byte("通过包裹类型实现数据压缩/解压缩"))
		if err != nil {
			t.Errorf("写入压缩数据失败,%+v", err)
		}
		fmt.Println("写入成功")
	})

	t.Run("read zip file", func(t *testing.T) {
		f, err := os.Open(file)
		if err != nil {
			t.Fatal(err)
		}
		z, err := gzip.NewReader(f)
		if err != nil {
			t.Fatal(err)
		}

		i := 1
		for {
			buf := make([]byte, 32)
			if _, err = z.Read(buf); err != nil {
				if err == io.EOF {
					fmt.Printf("%d.读取成功:%q\n到达文件尾部\n", i, buf)
				} else {
					fmt.Printf("%d.读取失败", i)
				}
				return
			}
			fmt.Printf("%d.读取成功:%q\n", i, buf)
			i++
		}
	})
}
