/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 *
 * https://leetcode.cn/problems/rotate-array/description/
 *
 * algorithms
 * Medium (44.26%)
 * Likes:    1697
 * Dislikes: 0
 * Total Accepted:    601.7K
 * Total Submissions: 1.4M
 * Testcase Example:  '[1,2,3,4,5,6,7]\n3'
 *
 * 给你一个数组，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。
 *
 *
 *
 * 示例 1:
 *
 *
 * 输入: nums = [1,2,3,4,5,6,7], k = 3
 * 输出: [5,6,7,1,2,3,4]
 * 解释:
 * 向右轮转 1 步: [7,1,2,3,4,5,6]
 * 向右轮转 2 步: [6,7,1,2,3,4,5]
 * 向右轮转 3 步: [5,6,7,1,2,3,4]
 *
 *
 * 示例 2:
 *
 *
 * 输入：nums = [-1,-100,3,99], k = 2
 * 输出：[3,99,-1,-100]
 * 解释:
 * 向右轮转 1 步: [99,-1,-100,3]
 * 向右轮转 2 步: [3,99,-1,-100]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= nums.length <= 10^5
 * -2^31 <= nums[i] <= 2^31 - 1
 * 0 <= k <= 10^5
 *
 *
 *
 *
 * 进阶：
 *
 *
 * 尽可能想出更多的解决方案，至少有 三种 不同的方法可以解决这个问题。
 * 你可以使用空间复杂度为 O(1) 的 原地 算法解决这个问题吗？
 *
 *
 *
 *
 *
 *
 *
 *
 */
package jzoffer

// TODO:
// @lc code=start

// 1. The other line of thought is a tad bit complicated but essentially
//  it builds on the idea of placing each element in its original position while
//  keeping track of the element originally in that position. Basically, at every step,
//  we place an element in its rightful position and keep track of the element already there or
//  the one being overwritten in an additional variable. We can't do this in one linear pass and
//  the idea here is based on cyclic-dependencies between elements.

func rotate1(nums []int, k int) {
	k = k % len(nums)
	res := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
	copy(nums, res) // 直接return res不行,res这样改了导致与最初始的切片底层数组脱离
	return
}

// 我们可以使用额外的数组来将每个元素放至正确的位置。用 n 表示数组的长度，我们遍历原数组，
// 将原数组下标为 i 的元素放至新数组下标为 (i+k) mod n的位置，最后将新数组拷贝至原数组即可。
// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode-solution-nipk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func rotate2(nums []int, k int) {
	newNums := make([]int, len(nums))
	for i, v := range nums {
		newNums[(i+k)%len(nums)] = v
	}
	copy(nums, newNums)
}

/* 题解来源  https://leetcode.cn/u/angel_monica/
思路：此题可以采用头插法，一个一个的移动。但是有种更加简单的选择数组的方式。我们可以采用翻转的方式，比如12345经过翻转就变成了54321，这样已经做到了把前面的数字放到后面去，但是还没有完全达到我们的要求，比如，我们只需要把12放在后面去，目标数组就是34512，与54321对比发现我们就只需要在把分界线前后数组再进行翻转一次就可得到目标数组了。所以此题只需要采取三次翻转的方式就可以得到目标数组，首先翻转分界线前后数组，再整体翻转一次即可。此题面试常考，大家可以记一下此方法。
class Solution {
public:
    void reverse(vector<int>& nums,int begin,int end)
    {
        int temp;
        while(begin<end){
            temp = nums[begin];
            nums[begin] = nums[end];
            nums[end] = temp;
            begin++;
            end--;
        }
    }
    void rotate(vector<int>& nums, int k) {
        if(nums.size()<2) return;
        k %=nums.size();
        reverse(nums,0,nums.size()-1);
        reverse(nums,0,k-1);
        reverse(nums,k,nums.size()-1);
    }
};
*/
// 旋转数组
func rotate(nums []int, k int) {
	if len(nums) < 2 {
		return
	}

	reverse := func(nums []int, l, r int) { // l,r是下标
		var temp int
		for l < r {
			temp = nums[l]
			nums[l] = nums[r]
			nums[r] = temp
			l++
			r--
		}
	}
	k = k % len(nums)
	reverse(nums, 0, len(nums)-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, len(nums)-1)
}

// 环状替换,一个数替换(n和k的最大公约数次会循环,n是数组长度)
// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/rotate-array/solution/xuan-zhuan-shu-zu-by-leetcode-solution-nipk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func rotate3(nums []int, k int) {
	n := len(nums)
	k %= n
	for start, count := 0, gcd(k, n); start < count; start++ {
		pre, cur := nums[start], start
		for ok := true; ok; ok = cur != start {
			next := (cur + k) % n
			nums[next], pre, cur = pre, nums[next], next
		}
	}
}

// 最大公约数
func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}

// @lc code=end
