/*
 * @lc app=leetcode.cn id=33 lang=cpp
 *
 * [33] 搜索旋转排序数组
 */

// @lc code=start
//#include"leetcode.h"

class Solution {
public:
    int search(vector<int>& nums, int target) {
        int l=0,r=nums.size()-1,mid=0,size=nums.size();
        while (l<=r)
        {
            mid=(l+r)/2;
            if (nums[mid]==target)
            {
                return mid;
            }
            if(nums[0]<=nums[mid])//0~mid有序
            { 
                 if (nums[0]<=target && target<=nums[mid])
                 {
                    r=mid-1;
                 }else
                 {
                      l=mid+1;
                 }
            }
            else{//mid~（size-1）有序
                if (nums[mid]<=target && target<=nums[size-1])
                {
                    l=mid+1;
                }else
                {
                    r=mid-1;
                }
            }        
        }
        return -1;        
    }
};
// @lc code=end

