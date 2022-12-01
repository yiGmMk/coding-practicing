package jzoffer

import (
	"testing"
)

type Handler interface {
	Book(start, end int) bool
}

func TestCases(t *testing.T) {
	type inNode struct {
		start, end int
	}
	ts := []struct {
		nodes []struct {
			val inNode
			res bool
		}
	}{
		// 测试用例1
		{
			nodes: []struct {
				val inNode
				res bool
			}{
				{inNode{start: 10, end: 20}, true},
				{inNode{start: 15, end: 25}, false},
				{inNode{start: 20, end: 30}, true},
				{inNode{start: 5, end: 15}, false},
			},
		},
		// 测试用例2
		{
			// [47,50],[33,41],[39,45],[33,42],[25,32],[26,35],[19,25],[3,8],[8,13],[18,27]]
			// [true,    true,   false, false,  true,   false,   true,   true, true, false]
			nodes: []struct {
				val inNode
				res bool
			}{
				{inNode{start: 47, end: 50}, true},
				{inNode{start: 33, end: 41}, true},
				{inNode{start: 39, end: 45}, false},
				{inNode{start: 33, end: 42}, false},
				{inNode{start: 25, end: 32}, true},
				{inNode{start: 26, end: 35}, false},
				{inNode{start: 19, end: 25}, true},
				{inNode{start: 3, end: 8}, true},
				{inNode{start: 8, end: 13}, true},
				{inNode{start: 18, end: 27}, false},
			},
		},
	}
	var handdler Handler
	for _, tt := range ts {
		handdler = Constructor2()
		for _, n := range tt.nodes {
			if handdler.Book(n.val.start, n.val.end) != n.res {
				t.Errorf("error,node:%+v", n)
			}
		}
	}
}
