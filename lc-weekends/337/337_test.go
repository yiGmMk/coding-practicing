package jzoffer

import (
	"fmt"
	"sort"
	"testing"
)

// 1
func evenOddBit(n int) []int {
	str := fmt.Sprintf("%b", n)
	odd, even := 0, 0
	id := 0
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] != '1' {
			id++
			continue
		}
		if id%2 == 0 {
			even++
		} else {
			odd++
		}
		id++
	}
	return []int{even, odd}
}

func TestXxx(t *testing.T) {
	// evenOddBit(2)

	/*
	   [
	   	[24,11,22,17,4],
	   	[21,16,5, 12,9],
	   	[6, 23,10,3, 18],
	   	[15,20,1, 8, 13],
	   	[0, 7, 14,19,2]
	   ]
	*/

	/*checkValidGrid([][]int{
		{24, 11, 22, 17, 4},
		{21, 16, 5, 12, 9},
		{6, 23, 10, 3, 18},
		{15, 20, 1, 8, 13},
		{0, 7, 14, 19, 2},
	})*/

}

func checkValidGrid(grid [][]int) bool {
	n, m := len(grid), len(grid[0])
	pos := make(map[int][]int, n*m)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			pos[grid[i][j]] = []int{i, j}
		}
	}
	// 必须从左上角出发
	if pos[0][0] != 0 || pos[0][1] != 0 {
		return false
	}
	// 00,12
	check := func(i, j int) bool {
		x, y := pos[i][0], pos[i][1]
		x1, y1 := pos[j][0], pos[j][1]

		if abs(x-x1) == 2 && abs(y-y1) == 1 {
			return true
		}

		if abs(x-x1) == 1 && abs(y-y1) == 2 {
			return true
		}
		return false
	}
	for i := 1; i < n*m; i++ {
		if !check(i, i-1) {
			return false
		}
	}
	return true
}

func Test3(t *testing.T) {
	fmt.Println("---------------------")
	//fmt.Println(beautifulSubsets([]int{2, 4, 6}, 2), 4)
	fmt.Println("---------------------")
	// [10,4,5,7,2,1], 3  23
	fmt.Println(beautifulSubsets([]int{10, 4, 5, 7, 2, 1}, 3), 23)
	fmt.Println("---------------------")
	fmt.Println(beautifulSubsets([]int{4, 2, 5, 9, 10, 3}, 1), 23)
	fmt.Println("---------------------")

}

// 子数组
func beautifulSubsets(nums []int, k int) int {
	sort.Ints(nums)
	res := 0
	n := len(nums)
	var backtrack func(start int, subset []int)
	backtrack = func(start int, subset []int) {
		if len(subset) > 1 {
			for i := 0; i < len(subset); i++ {
				if abs(subset[i]-subset[len(subset)-1]) == k {
					return
				}
			}
			//fmt.Println(nums, subset)
			res++
		}
		for i := start; i < n; i++ {
			subset = append(subset, nums[i])
			backtrack(i+1, subset)
			subset = subset[:len(subset)-1]
		}
	}
	backtrack(0, []int{})
	return res + n
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

/*
作者：endlesscheng
链接：https://leetcode.cn/problems/the-number-of-beautiful-subsets/solution/tao-lu-zi-ji-xing-hui-su-pythonjavacgo-b-fcgs/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
func beautifulSubsets1(nums []int, k int) int {
	ans := -1                    // 去掉空集
	cnt := make([]int, 1001+k*2) // 用数组实现比哈希表更快
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			ans++
			return
		}
		dfs(i + 1)       // 不选
		x := nums[i] + k // 避免负数下标
		if cnt[x-k] == 0 && cnt[x+k] == 0 {
			cnt[x]++ // 选
			dfs(i + 1)
			cnt[x]-- // 恢复现场
		}
	}
	dfs(0)
	return ans
}
