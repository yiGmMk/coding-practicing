import math


class Solution:
    def judeSquareSum(c: int) -> bool:
        l, r = 0,  (int)(math.sqrt(c))  # 默认是浮点型
        while l <= r:
            v = l*l+r*r
            if v == c:
                return True
            elif v > c:
                r -= 1
            else:
                l += 1
        return False


testCases = [
    {"num": 5, "res": True},
    {"num":3,"res": False},
]

s = Solution
for i, v in enumerate(testCases):
    got = s.judeSquareSum(v["num"])
    if got != v["res"]:
        print("[%d].got:%d,expect:%d", i, got, v["res"])
