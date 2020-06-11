/*
 * @lc app=leetcode.cn id=54 lang=cpp
 *
 * [54] 螺旋矩阵
 */

// @lc code=start
#include"leetcode.h"
class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>>& matrix) {
        if (matrix.size()==0 || matrix[0].size()==0)
        {
            return {};
        }
        int rows=matrix.size(),cols=matrix[0].size();
        vector<int>  vecRtn;
        int top=0,left=0,right=cols-1,bottom=rows-1;

        while (left<=right && top<=bottom)
        {
            for (size_t column =left;column< cols; column++)
            {
                vecRtn.push_back(matrix[top][column]);
            }
            for (size_t row = top+1; row < rows; row++)
            {
                vecRtn.push_back(matrix[row][right]);
            }
            if (left<right && top<bottom)
            {
                for (size_t column = right-1; column >left ;column--)
                {
                    vecRtn.push_back(matrix[bottom][column]);
                }
                for (size_t row = bottom-1; row >top ; row--)
                {
                    vecRtn.push_back(matrix[row][left]);
                }                
            }   
            left++;
            right--;
            top++;
            bottom--;   
        }
        return vecRtn;
    }
};
// @lc code=end

