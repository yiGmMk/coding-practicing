
def isBadVersion(version: int) -> bool:
    return False


class Solution:
    def firstBadVersion(self, n: int) -> int:
        left, right = 0, n
        while (left+1 < right):
            mid = left+(right-left)//2
            if isBadVersion(mid):
                right = mid
            else:
                left = mid
        return right
