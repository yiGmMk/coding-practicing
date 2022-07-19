/*
 * @lc app=leetcode.cn id=938 lang=cpp
 *
 * [938] 二叉搜索树的范围和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     TreeNode *left;
 *     TreeNode *right;
 *     TreeNode(int x) : val(x), left(NULL), right(NULL) {}
 * };
 */
//#include"leetcode.h"
// class Solution {
// public:
//     void search(TreeNode*root,vector<int> &vec)
//     {
//         if (!root)
//         {
//             return;
//         }
//         else
//         {
//             vec.push_back(root->val);
//             search(root->left,vec);
//             search(root->right,vec);
//         }

//     }
//     // int rangeSumBST(TreeNode* root, int L, int R) {
//     //     vector<int> vec;

//     //     search(root,vec);
//     //     int iRes=0;
//     //     for (size_t i = 0; i < vec.size(); i++)
//     //     {
//     //         if (vec[i]>=L && vec[i]<=R)
//     //         {
//     //              iRes+=vec[i];
//     //         }

//     //     }
//     //     return iRes;
//     // }
//     void BinarySearch(TreeNode* root,int &res,int L,int R)
//     {
//         if (!root)
//         {
//             return;
//         }
//         if (root->val>=L && root->val<=R)
//         {
//             res+=root->val;
//             BinarySearch(root->right,res,L,R);
//             BinarySearch(root->left,res,L,R);
//         }
//         else if (root->val<L)
//         {
//            BinarySearch(root->right,res,L,R);
//         }
//         else if (root->val>R)
//         {
//             BinarySearch(root->left,res,L,R);
//         }

//     }
//     int rangeSumBST(TreeNode* root, int L, int R) {
//         int iRes=0;
//         BinarySearch(root,iRes,L,R);

//         return iRes;
//     }
// };

// 方法一：深度优先搜索
// 我们对树进行深度优先搜索，对于当前节点 node，如果 node.val 小于等于 L，那么只需要继续搜索它的右子树；如果 node.val 大于等于 R，那么只需要继续搜索它的左子树；如果 node.val 在区间 (L, R) 中，则需要搜索它的所有子树。

// 我们在代码中用递归和迭代的方法分别实现了深度优先搜索。

// 作者：LeetCode
// 链接：https://leetcode-cn.com/problems/range-sum-of-bst/solution/er-cha-sou-suo-shu-de-fan-wei-he-by-leetcode/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
#include "leetcode.h"
class Solution
{
public:
    int ans;
    int rangeSumBST(TreeNode *root, int L, int R)
    {
        ans = 0;
        dfs(root, L, R);
        return ans;
    }

    void dfs(TreeNode *node, int L, int R)
    {
        if (node != nullptr)
        {
            if (L <= node->val && node->val <= R)
                ans += node->val;
            if (L < node->val)
                dfs(node->left, L, R);
            if (node->val < R)
                dfs(node->right, L, R);
        }
    }
};
// @lc code=end

int main()
{
    Solution s;
    TreeNode *root = new TreeNode(10);
    root->left = new TreeNode(5);
    root->right = new TreeNode(15);
    root->left->left = new TreeNode(3);
    root->left->right = new TreeNode(7);
    root->right->right = new TreeNode(18);
    int res = s.rangeSumBST(root, 7, 15);
    cout << "res from tree:" << res << endl;
    return 0;
}