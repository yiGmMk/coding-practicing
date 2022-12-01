package jzoffer

// 遍历字符串，对每个字符，都看作回文的中心，向两端延申进行判断直到非回文。
// 回文的中心可能是一个字符，也可能是两个字符。
// 注意双指针可能越界。
// 代码

// class Solution {
//     public int countSubstrings(String s) {
//         if (s == null || s.length() == 0) {
//             return 0;
//         }
//         int count = 0;
//         //字符串的每个字符都作为回文中心进行判断，中心是一个字符或两个字符
//         for (int i = 0; i < s.length(); ++i) {
//             count += countPalindrome(s, i, i);
//             count += countPalindrome(s, i, i+1);
//         }
//         return count;
//     }

//     //从字符串的第start位置向左，end位置向右，比较是否为回文并计数
//     private int countPalindrome(String s, int start, int end) {
//         int count = 0;
//         while (start >= 0 && end < s.length() && s.charAt(start) == s.charAt(end)) {
//             count++;
//             start--;
//             end++;
//         }
//         return count;
//     }
// }

// 作者：ling-han-i
// 链接：https://leetcode.cn/problems/a7VOhD/solution/jian-zhi-offer-zhuan-xiang-tu-po-ban-gua-6bln/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
