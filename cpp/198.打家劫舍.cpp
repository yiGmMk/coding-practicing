/*
 * @lc app=leetcode.cn id=198 lang=cpp
 *
 * [198] 打家劫舍
 */

// @lc code=start
#include"leetcode.h"
class Solution {
public:
    // int rob(vector<int>& nums) {
    //     if (nums.empty())
    //     {
    //         return 0;
    //     }
    //     if (nums.size()==1)
    //     {
    //         return nums[0];
    //     }
        
    //     if (nums.size()==2)
    //     {
    //         return std::max(nums[0],nums[1]);
    //     }
    //     vector<int> dp;
    //     dp.push_back(nums[0]);
    //     dp.push_back(std::max(nums[0],nums[1]));
    //     size_t size=nums.size();
    //     for (size_t i = 2; i < size; i++)
    //     {
    //         dp.push_back(max(dp[i-2]+nums[i],dp[i-1]));
    //     }
    //     return dp.back();
    // }

    int rob(vector<int>& nums) {
        if (nums.empty())
        {
            return 0;
        }
        if (nums.size()==1)
        {
            return nums[0];
        }
        
        if (nums.size()==2)
        {
            return std::max(nums[0],nums[1]);
        }
        int pre=nums[0],next=std::max(nums[0],nums[1]),iMax=0;
        size_t size=nums.size();
        for (size_t i = 2; i < size; i++)
        {
            iMax=std::max(pre+nums[i],next);
            pre=next;
            next=iMax;
        }
        return iMax;
    }
};

int main()
{
    vector<int>sss={2,7,9,3,1};
    Solution s;
    s.rob(sss);
    return 1;
}
// @lc code=end

