from typing import List


class Solution:
    def checkXMatrix(self, grid: List[List[int]]) -> bool:
        for i, row in enumerate(grid):
            for j, v in enumerate(row):
                if (v == 0) == (j == i or i+j == len(v)-1):
                    return False

        return True
