
from typing import List


class Solution:
    def rotate(self, nums: List[int], k: int) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        k = k % len(nums)
        nums[:] = nums[len(nums)-k:]+nums[:len(nums)-k]

    # 作者：genening
    # 链接：https://leetcode.cn/problems/rotate-array/solution/wei-shi-yao-by-genening-fkz1/
    # 来源：力扣（LeetCode）
    # 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
    def rotate1(self, nums: List[int], k: int) -> None:
        length = len(nums)
        k %= length
        nums[:] = nums[::-1]
        nums[:k] = nums[:k][::-1]
        nums[k:] = nums[k:][::-1]


s = Solution
nums = [1, 2, 3, 4, 5, 6, 7]
s.rotate(s, nums=nums, k=3)
if nums != [5, 6, 7, 1, 2, 3, 4]:
    print("rotate err:", nums)
else:
    print("rotated:", nums)


nums = [1, 2, 3, 4, 5, 6, 7]
s.rotate1(s, nums=nums, k=3)
if nums != [5, 6, 7, 1, 2, 3, 4]:
    print("rotate err:", nums)
else:
    print("rotated:", nums)
