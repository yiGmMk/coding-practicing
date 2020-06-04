/*
 * @lc app=leetcode.cn id=104 lang=cpp
 *
 * [104] 二叉树的最大深度
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
#include"leetcode.h"

class Solution {
public:
    // int maxDepth(TreeNode* root) {
    //     if (!root)
    //     {
    //         return 0;
    //     }
    //     int deep=0;
    //     queue<TreeNode*> q;
    //     q.push(root); 
    //     while (!q.empty())
    //     {
    //         int iSize=q.size();
    //         for (size_t i = 0; i < iSize; i++)
    //         {
    //             TreeNode* node=q.front();
    //             q.pop();
    //             if (node->left)
    //             {
    //                 q.push(node->left);
    //             }
    //             if (node->right)
    //             {
    //                 q.push(node->right);
    //             }
    //         }
    //         deep++;            
    //     }
    //     return deep;
    // }

    //深度优先：递归版
    int maxDepth(TreeNode* root) {
        if(root==NULL) return 0;
        int l=maxDepth(root->left)+1;
        int r=maxDepth(root->right)+1;
        return l>r?l:r;   
    }
};
// @lc code=end

