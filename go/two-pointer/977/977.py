
from typing import List


class Solution(object):
    def sortedSquares(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        return sorted(num*num for num in nums)

    def sortedSquares1(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [0]*n
        l, r, pos = 0, n-1, n-1
        while l <= r:
            vl = nums[l] * nums[l]
            vr = nums[r] * nums[r]
            if vl > vr:
                ans[pos] = vl
                l += 1
            else:
                ans[pos] = vr
                r -= 1
            pos -= 1
        return ans


testCase = [
    {
        "nums": [-4, -1, 0, 3, 10],
        "res":  [0, 1, 9, 16, 100],
    },
    {
        "nums": [-7, -3, 2, 3, 11],
        "res":  [4, 9, 9, 49, 121],
    },
]

s = Solution
for i, v in enumerate(testCase):
    got = s.sortedSquares1(s, nums=v["nums"])
    got[0] = 1
    if got != v["res"]:
        print(i, "got:", got, "expect:", v["res"])
