package jzoffer

func pivotIndexOn(nums []int) int {
	total := 0
	for _, v := range nums {
		total += v
	}
	sum := 0
	for i, v := range nums {
		// left = sum
		// right = total - sum - v
		// left == right => 2*sum+v == total
		if 2*sum+v == total {
			return i
		}
		sum += v
	}
	return -1
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/tvdfij/solution/zuo-you-liang-bian-zi-shu-zu-de-he-xiang-5j4r/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
