package jzoffer

import "testing"

func Test547(t *testing.T) {
	// Testcase
	// [[1,0,0,1],[0,1,1,0],[0,1,1,1],[1,0,1,1]]
	// Expected Answer
	// 1
	if findCircleNum([][]int{
		{1, 0, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
		{1, 0, 1, 1}}) != 1 {
		t.Fail()
	}
}
