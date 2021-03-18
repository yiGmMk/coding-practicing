/*
 * @lc app=leetcode.cn id=1431 lang=cpp
 *
 * [1431] 拥有最多糖果的孩子
 */

// @lc code=start
#include"leetcode.h"
class Solution {
public:
    vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {
        vector<int> num;
        for(int x:candies )
            num.push_back(x);
        sort(num.begin(),num.end());

        unsigned int max=num.back();
        vector<bool> rtn;
        for(int x:candies)
        {
            rtn.push_back((x+extraCandies)>=max);
        }
        return rtn;
    }

    // vector<bool> kidsWithCandies(vector<int>& candies, int extraCandies) {
    //      int n = candies.size();
    //     int maxCandies = *max_element(candies.begin(), candies.end());
    //     vector<bool> ret;
    //     for (int i = 0; i < n; ++i) {
    //         ret.push_back(candies[i] + extraCandies >= maxCandies);
    //     }
    //     return ret;

    //     // 作者：LeetCode-Solution
    //     // 链接：https://leetcode-cn.com/problems/kids-with-the-greatest-number-of-candies/solution/yong-you-zui-duo-tang-guo-de-hai-zi-by-leetcode-so/
    //     // 来源：力扣（LeetCode）
    //     // 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
    // }
};
// @lc code=end

