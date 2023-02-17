
from typing import List


class Solution:

    def minSumOfLengths(self, arr: List[int], target: int) -> int:
        INF = len(arr) + 1
        # the ith index represents the smallest length subarray we've found ending <= i that sums to target
        best_at_i = [INF]*len(arr)
        best = INF  # output
        curr_sum = 0  # current sum between left and right

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


s = Solution()
print(s.minSumOfLengths([3,2,2,4,3], 3))
print(s.minSumOfLengths([1, 2, 1, 3, 1], 4))
