package jzoffer

import (
	"fmt"
	"reflect"
	"testing"
)

// todo:
func differenceOfSum(nums []int) int {
	a, b := 0, 0
	for _, v := range nums {
		a += v
		for v > 0 {
			b += v % 10
			v /= 10
		}
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	return abs(a - b)
}

func rangeAddQueries(n int, queries [][]int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	//val := make(map[int]int, n*n)
	for _, v := range queries {
		for i := v[0]; i <= v[2]; i++ {
			for j := v[1]; j <= v[3]; j++ {
				//val[i*n+j]++
				res[i][j]++
			}
		}
	}
	return res

}

// TODO:解法不对
func countGood(nums []int, k int) int64 {
	check := func(cnt map[int]int) bool {
		// 相等的数
		num := 0
		for _, v := range cnt {
			/// 2 2 2 // 2 + 1=3 // 3*2/*2
			if v < 2 {
				continue
			}
			if v == 2 {
				num++
				continue
			}
			num += (v * (v - 1)) / 2 //5*4/2
		}
		return num >= k
	}

	cnt, res := make(map[int]int, 0), int64(0)
	var dfs func(i int, cnt *map[int]int)
	dfs = func(i int, cnt *map[int]int) {
		if i == len(nums) {
			if check(*cnt) {
				res++
			}
			return
		}
		for j := i; j < len(nums); j++ {
			(*cnt)[nums[j]]++
			dfs(j+1, cnt)
			(*cnt)[nums[j]]--
		}
	}
	for i, v := range nums {
		cnt[v]++
		dfs(i+1, &cnt)
		cnt[v]--
	}
	return res
}

var _testCase = []struct {
	mat    [][]int
	k      int
	expect []int
}{
	{
		mat: [][]int{
			{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1},
		},
		k:      3,
		expect: []int{2, 0, 3},
	},
}

func Test1337(t *testing.T) {
	fmt.Println(countGood([]int{3, 3, 1, 2, 2}, 2))
	fmt.Println(countGood([]int{3, 1, 4, 3, 2, 2, 4}, 2))

	for i, v := range _testCase {
		got := kWeakestRows(v.mat, v.k)
		if !reflect.DeepEqual(got, v.expect) {
			t.Errorf("[%d].got:%v,expect:%v", i, got, v.expect)
		}
	}
}
