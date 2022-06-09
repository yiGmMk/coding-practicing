package jzoffer

import (
	"strings"
	"testing"
)

var ts = [][]string{
	[]string{"1", "2", "3", "null", "null", "4", "5", "6", "7"},
}

func TestBuild(t *testing.T) {
	var cc BTreeCodec
	cc = &BTreeBfsCodec{}
	for _, tc := range ts {
		in := strings.Join(tc, ",")
		root := cc.Deserialize(in)
		res := cc.Serialize(root)
		if in != res {
			t.Errorf("in: %v, res: %v not equal", in, res)
		}
	}
}

func TestDfs(t *testing.T) {
	var cc BTreeCodec
	cc = &CodecDfs{}
	for _, tc := range ts {
		in := strings.Join(tc, ",")
		root := cc.Deserialize(in)
		res := cc.Serialize(root)
		if in != res {
			t.Errorf("in: %v, res: %v not equal", in, res)
		}
	}
}
