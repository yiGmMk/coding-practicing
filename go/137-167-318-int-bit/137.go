package jzoffer

// @lc code=start
func singleNumberOn(nums []int) int {
	im := make(map[int]int, len(nums)/3)
	for _, v := range nums {
		im[v]++
	}
	var res int
	for k, v := range im {
		if v == 1 {
			res = k
		}
	}
	return res
}

// @lc code=end
