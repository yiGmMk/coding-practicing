package jzoffer

import (
	"container/heap"
	"reflect"
	"testing"
)

var fFix = func(nums1 []int, nums2 []int, k int) [][]int {
	out := [][]int{}
	h := minHeap{}
	for i := 0; i < len(nums1) && i < k; i++ {
		for j := 0; j < len(nums2) && j < k; j++ {
			if h.Len() < k {
				heap.Push(&h, node{x: nums1[i], y: nums2[j], sum: nums1[i] + nums2[j]})
			} else if nums1[i]+nums2[j] <= h.data[0].sum {
				heap.Pop(&h)
				heap.Push(&h, node{x: nums1[i], y: nums2[j], sum: nums1[i] + nums2[j]})
				heap.Fix(&h, h.Len()-1)
			}
		}
	}
	for h.Len() > 0 && len(out) < k {
		n := heap.Pop(&h).(node)
		out = append(out, []int{n.x, n.y})
	}
	return out
}

func TestMinSum(t *testing.T) {
	ts := []struct {
		n1  []int
		n2  []int
		k   int
		res [][]int
	}{

		{
			n1: []int{-10, -4, 0, 0, 6},
			n2: []int{3, 5, 6, 7, 8, 100},
			k:  10,
			res: [][]int{
				{-10, 3},
				{-10, 5},
				{-10, 6},
				{-10, 7},
				{-10, 8},
				{-4, 3},
				{-4, 5},
				{-4, 6},
				{0, 3},
				{0, 3},
			},
		},
		{
			n1: []int{1, 2, 4, 5, 6},
			n2: []int{3, 5, 7, 9},
			k:  3,
			res: [][]int{
				{1, 3},
				{2, 3},
				{1, 5},
			},
		},
		{
			n1: []int{1, 1, 2},
			n2: []int{1, 2, 3},
			k:  2,
			res: [][]int{
				{1, 1},
				{1, 1},
			},
		},
		{
			n1: []int{1, 1, 2},
			n2: []int{1, 2, 3},
			k:  10,
			res: [][]int{
				{1, 1},
				{1, 1},
				{1, 2},
				{1, 2},
				{2, 1},
				{1, 3},
				{2, 2},
				{1, 3},
				{2, 3},
			},
		},
		{
			n1: []int{1, 2},
			n2: []int{3},
			k:  3,
			res: [][]int{
				{1, 3},
				{2, 3},
			},
		},
		{
			n1: []int{1, 7, 11},
			n2: []int{2, 4, 6},
			k:  3,
			res: [][]int{
				{1, 2},
				{1, 4},
				{1, 6},
			},
		},
	}
	var f = fFix
	for _, tt := range ts {
		res := f(tt.n1, tt.n2, tt.k)
		if !reflect.DeepEqual(res, tt.res) {
			t.Errorf("not equal,in_1:%+v,in_2:%+v,expect:%+v,got:%+v\n", tt.n1, tt.n2, tt.res, res)
		}
	}
}

func fkSmallestPairs(nums1, nums2 []int, k int) (ans [][]int) {
	m, n := len(nums1), len(nums2)
	h := HP{nil, nums1, nums2}
	for i := 0; i < k && i < m; i++ {
		h.data = append(h.data, pair{i, 0})
	}
	for h.Len() > 0 && len(ans) < k {
		p := heap.Pop(&h).(pair)
		i, j := p.i, p.j
		ans = append(ans, []int{nums1[i], nums2[j]})
		if j+1 < n {
			heap.Push(&h, pair{i, j + 1})
		}
	}
	return
}
