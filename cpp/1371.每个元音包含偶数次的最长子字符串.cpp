/*
 * @lc app=leetcode.cn id=1371 lang=cpp
 *
 * [1371] 每个元音包含偶数次的最长子字符串
 */

// @lc code=start
#include"leetcode.h"
class Solution {
public:
    int findTheLongestSubstring(string s) {
        int ans = 0, status = 0, n = s.length();
        vector<int> pos(1 << 5, -1);//32 number
        pos[0] = 0;
        for (int i = 0; i < n; ++i) {
            if (s[i] == 'a') {
                status ^= 1<<0;
            } else if (s[i] == 'e') {
                status ^= 1<<1;
            } else if (s[i] == 'i') {
                status ^= 1<<2;
            } else if (s[i] == 'o') {
                status ^= 1<<3;
            } else if (s[i] == 'u') {
                status ^= 1<<4;
            }
            if (~pos[status]) {
                ans = max(ans, i + 1 - pos[status]);
            } else {
                pos[status] = i + 1;
            }
        }
        return ans;
    }
};

int main()
{
    Solution s;
    int iRtn=s.findTheLongestSubstring("eleetminicoworoep");//13
    return iRtn;
}
// @lc code=end

