/*
 * @lc app=leetcode.cn id=1477 lang=golang
 *
 * [1477] 找两个和为目标值且不重叠的子数组
 *
 * https://leetcode.cn/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/description/
 *
 * algorithms
 * Medium (30.78%)
 * Likes:    124
 * Dislikes: 0
 * Total Accepted:    8.3K
 * Total Submissions: 26.8K
 * Testcase Example:  '[3,2,2,4,3]\n3'
 *
 * 给你一个整数数组 arr 和一个整数值 target 。
 *
 * 请你在 arr 中找 两个互不重叠的子数组 且它们的和都等于 target 。可能会有多种方案，请你返回满足要求的两个子数组长度和的 最小值 。
 *
 * 请返回满足要求的最小长度和，如果无法找到这样的两个子数组，请返回 -1 。
 *
 *
 *
 * 示例 1：
 *
 * 输入：arr = [3,2,2,4,3], target = 3
 * 输出：2
 * 解释：只有两个子数组和为 3 （[3] 和 [3]）。它们的长度和为 2 。
 *
 *
 * 示例 2：
 *
 * 输入：arr = [7,3,4,7], target = 7
 * 输出：2
 * 解释：尽管我们有 3 个互不重叠的子数组和为 7 （[7], [3,4] 和 [7]），但我们会选择第一个和第三个子数组，因为它们的长度和 2
 * 是最小值。
 *
 *
 * 示例 3：
 *
 * 输入：arr = [4,3,2,6,2,3,4], target = 6
 * 输出：-1
 * 解释：我们只有一个和为 6 的子数组。
 *
 *
 * 示例 4：
 *
 * 输入：arr = [5,5,4,4,5], target = 3
 * 输出：-1
 * 解释：我们无法找到和为 3 的子数组。
 *
 *
 * 示例 5：
 *
 * 输入：arr = [3,1,1,1,5,1,2,1], target = 3
 * 输出：3
 * 解释：注意子数组 [1,2] 和 [2,1] 不能成为一个方案因为它们重叠了。
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= arr.length <= 10^5
 * 1 <= arr[i] <= 1000
 * 1 <= target <= 10^8
 *
 *
 */
package jzoffer

import (
	"sort"
)

// TODO

// @lc code=start
/*
有几个测试用例过不了,从前往后滑窗,将结果集过滤,导致,有可能从后往前滑窗得出更短的结果
	{
		nums:   []int{2, 1, 3, 3, 2, 3, 1},
		target: 6,
		expect: 5,
	},
// 从前往后+从后往前 两次遍历还是有些过不了
*/
func minSumOfLengthsWrongSolution(arr []int, target int) int {
	cal := func(arr []int) int {
		ans, l, sum := []int{}, 0, 0
		for r, v := range arr {
			if v == target {
				ans = append(ans, 1)
				l = r + 1
				sum = 0
				continue
			}
			sum += v
			for sum > target {
				sum -= arr[l]
				l++
			}
			if sum == target {
				ans = append(ans, r-l+1)
				sum = 0
				l = r + 1 //不能重叠
			}
		}
		if len(ans) < 2 {
			return -1
		}
		sort.Ints(ans)
		return ans[0] + ans[1]
	}

	p := cal(arr)
	reverse(arr)
	n := cal(arr)
	return min(p, n)
}

func reverse(arr []int) {
	for l, r := 0, len(arr)-1; l < r; {
		arr[l], arr[r] = arr[r], arr[l]
		l++
		r--
	}
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

/*
题解: https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/solutions/775374/python-o-n-sliding-window-with-comments/?topicTags=sliding-window
class Solution:

	def minSumOfLengths(self, arr: List[int], target: int) -> int:
	    INF = len(arr) + 1
	    best_at_i = [INF]*len(arr) # the ith index represents the smallest length subarray we've found ending <= i that sums to target
	    best = INF # output
	    curr_sum = 0 # current sum between left and right

	    left = 0
	    for right in range(len(arr)):
	        # update the running sum
	        curr_sum += arr[right]

	        # arr is strictly positive, so shrink window until we're not above target
	        while curr_sum > target and left <= right:
	            curr_sum -= arr[left]
	            left += 1

	        if curr_sum == target:
	            # we have a new shortest candidate to consider
	            best = min(best, best_at_i[left-1] + right - left + 1)
	            best_at_i[right] = min(best_at_i[right-1], right - left + 1)
	        else:
	            # best we've seen is the previous best (overlaps to end if right == 0)
	            best_at_i[right] = best_at_i[right-1]

	    if best == INF:
	        return -1
	    return best
*/
func minSumOfLengthsDp1(arr []int, target int) int {
	inf := len(arr) + 1
	best, bestI := inf, make([]int, len(arr))
	for i := range bestI {
		bestI[i] = inf
	}

	left, sum := 0, 0
	for right := range arr {
		sum += arr[right]
		for sum > target && left <= right {
			sum -= arr[left]
			left++
		}

		if sum == target {
			if left > 0 {
				best = min(best, bestI[left-1]+right-left+1)
			}
			if left > 0 {
				bestI[right] = min(bestI[right-1], right-left+1)
			} else {
				bestI[right] = right - left + 1
			}
		} else {
			if right != 0 {
				bestI[right] = bestI[right-1]
			}
		}

	}
	if best == inf {
		return -1
	}
	return best
}

// https://leetcode.com/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/solutions/685486/java-o-n-time-two-pass-solution-using-hashmap/
// 利用数组中的数全是正数的特点,先用map出入前缀和,在i找左右两边和为target的子数组
// class Solution {
//     public int minSumOfLengths(int[] arr, int target) {
//         HashMap<Integer,Integer> hmap=new HashMap<>();
//         int sum=0,lsize=Integer.MAX_VALUE,result=Integer.MAX_VALUE;
//         hmap.put(0,-1);
//         for(int i=0;i<arr.length;i++){
//             sum+=arr[i];
//             hmap.put(sum,i); // stores key as sum upto index i, and value as i.
//         }
//         sum=0;
//         for(int i=0;i<arr.length;i++){
//             sum+=arr[i];
//             if(hmap.get(sum-target)!=null){
//                 lsize=Math.min(lsize,i-hmap.get(sum-target));      // stores minimum length of sub-array ending with index<= i with sum target. This ensures non- overlapping property.
//             }
// 			//hmap.get(sum+target) searches for any sub-array starting with index i+1 with sum target.
//             if(hmap.get(sum+target)!=null&&lsize<Integer.MAX_VALUE){
//                 result=Math.min(result,hmap.get(sum+target)-i+lsize); // updates the result only if both left and right sub-array exists.
//             }
//         }
//         return result==Integer.MAX_VALUE?-1:result;
//     }
// }

/*
作者：antione
链接：https://leetcode.cn/problems/find-two-non-overlapping-sub-arrays-each-with-target-sum/solution/yi-ci-bian-li-hua-dong-chuang-kou-jia-dong-tai-gui/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

	class Solution {
	    public int minSumOfLengths(int[] arr, int target) {
	        int n = arr.length;
	        int[] dp = new int[n];
	        // 注意不能设置为最大值，因为相加会溢出
	        Arrays.fill(dp, Integer.MAX_VALUE / 2);

	        int ans = Integer.MAX_VALUE;
	        for(int i = 0, j = 0, sum = 0; j < n; j++){
	            sum += arr[j];
	            while(i <= j && sum > target){
	                sum -= arr[i++];
	            }
	            // 找到满足条件的一个区间
	            if(sum == target){
	                dp[j] = j - i + 1;
	                if(i != 0){
	                    ans = Math.min(ans, dp[i-1] + j - i + 1);
	                }
	            }
	            if(j != 0)
	                dp[j] = Math.min(dp[j], dp[j-1]);
	        }

	        return ans > arr.length ? -1 : ans;
	    }
	}
*/
func minSumOfLengths(arr []int, target int) int {
	// 和上面那个题解思路一样
	inf := len(arr) + 1
	ans, dp := inf, make([]int, len(arr))
	for i := range dp {
		dp[i] = inf
	}

	left, sum := 0, 0
	for right := range arr {
		sum += arr[right]
		for sum > target && left <= right {
			sum -= arr[left]
			left++
		}
		if sum == target {
			dp[right] = right - left + 1
			if left != 0 {
				ans = min(ans, dp[left-1]+dp[right])
			}
		}
		if right != 0 {
			dp[right] = min(dp[right], dp[right-1])
		}
	}
	if ans == inf {
		return -1
	}
	return ans
}

// @lc code=end
