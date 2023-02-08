package jzoffer

import (
	"testing"

	"github.com/samber/lo"
)

var _testcase = []struct {
	folder, expect []string
}{
	{
		folder: []string{"/a", "/a/b/c", "/a/b/d"},
		expect: []string{"/a"},
	},
	{
		folder: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
		expect: []string{"/a/b/c", "/a/b/ca", "/a/b/d"},
	},
	{
		folder: []string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"},
		expect: []string{"/a", "/c/d", "/c/f"},
	},
}

func Test1233(t *testing.T) {
	for i, v := range _testcase {
		got := removeSubfolders(v.folder)
		if l, r := lo.Difference(got, v.expect); len(l) > 0 || len(r) > 0 {
			t.Errorf("[%d].got:%s,expect:%s", i, got, v.expect)
		}
	}
}
