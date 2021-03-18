#include "leetcode.h"
class Solution {
public:
    bool check(string &s,int low,int high)
    {
        for(int i=low,j=high;i<j;i++,j--)
        {
            if (s[i]!=s[j])
            {
                return false;
            }
            
        }
        return true;
    }
    bool validPalindrome(string s) {
        int size=s.size();
        for(int i=0,j=size-1;i<j;i++,j--)
        {
            if(s[i]==s[j])
            {
                continue;
            }
            else
            {
                return check(s,i,j-1) || check(s,i+1,j);
            }
            
        }
        return true;
    }
};

int main()
{
    Solution s;
    int irtn=0;
    return irtn=s.validPalindrome("abca");
}