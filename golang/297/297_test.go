package jzoffer

import (
	"strings"
	"testing"
)

var ts = [][]string{
	{"1", "2", "3", "null", "null", "4", "5", "6", "7"},
}

func TestBuild(t *testing.T) {
	cc := &BTreeBfsCodec{}
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
	cc := &CodecDfs{}
	for _, tc := range ts {
		in := strings.Join(tc, ",")
		root := cc.Deserialize(in)
		res := cc.Serialize(root)
		if in != res {
			t.Errorf("in: %v, res: %v not equal", in, res)
		}
	}
}

func TestBfsLeetcode(t *testing.T) {
	cc := &BfsLeetcode{}
	for _, tc := range ts {
		in := strings.Join(tc, ",")
		root := cc.Deserialize(in)
		res := cc.Serialize(root)
		if in != res {
			t.Errorf("in: %v, res: %v not equal", in, res)
		}
	}
}
