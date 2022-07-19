#include <string>
#include <map>
#include <algorithm>
#include <vector>
using namespace std;

// 使用pair记录单个word的起始位置和长度，存入vector方便排序。排序完再将每个词叠加返回。
// vector<pair<int,int>> ;
class Solution
{
public:
    string arrangeWords(string text)
    {
        string str = text;
        vector<pair<int, int>> strVec;
        // 1.char
        size_t st = 0, ed = 0;
        for (auto c : str)
        { //"Leetcode is cool"
            if (c == (char)' ')
            {
                strVec.push_back(pair<int, int>(st, ed - st));
                st = ed + 1; // skip char' '
            }
            ed++; //
        }
        strVec.push_back(pair<int, int>(st, ed - st)); // add last word

        sort(strVec.begin(), strVec.end(), [=](const pair<int, int> &a, const pair<int, int> &b)
             {
            if(a.second!=b.second)//length  
                return a.second < b.second;  
            else 
                return a.first<b.first; });
        string strRtn;
        str[0] = str[0] - ('A' - 'a');
        for (auto it = strVec.begin(); it != strVec.end(); it++)
        {
            strRtn += str.substr(it->first, it->second);
            if (it != strVec.end() - 1)
                strRtn += ' ';
        }
        strRtn[0] += ('A' - 'a');
        return strRtn;
    }
};