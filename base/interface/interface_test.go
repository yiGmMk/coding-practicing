package base

import (
	"fmt"
	"io"
	"strings"
	"testing"
)

type TxtReader struct{}

func (*TxtReader) Read(p []byte) (n int, err error) {
	// ...
	return 0, nil
}

func NewTxtReader(path string) io.Reader {
	var r *TxtReader
	if strings.Contains(path, ".txt") {
		r = new(TxtReader)
	}
	return r
}

// interface在运行时的表示分为两部分,类型信息和值信息,都为nil,interface才是nil
func TestI(t *testing.T) {
	t.Run("", func(t *testing.T) {
		i := NewTxtReader("/home/tony/test.png")
		if i == nil {
			println("fail to init txt reader")
			return
		}
		println("init txt reader ok")
	})

	t.Run("", func(t *testing.T) {
		var r *TxtReader = nil
		var i io.Reader = r
		fmt.Printf("r==nil:%v, r:%+v, i==nil:%v, i:%+v\n", r == nil, r, i == nil, i)
	})
}
