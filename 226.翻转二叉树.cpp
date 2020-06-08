/*
 * @lc app=leetcode.cn id=226 lang=cpp
 *
 * [226] 翻转二叉树
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
    TreeNode* invertTree(TreeNode* root) {
        if(!root)
            return NULL;
        TreeNode *pNodeL=root->right;        
        TreeNode *pNodeR=root->left;
        root->left=invertTree(pNodeL);
        root->right=invertTree(pNodeR);
        return root;
    }
};


int main()
{

    TreeNode  *pNode=new TreeNode(3);

    Solution s;
    s.invertTree(pNode);


    delete pNode;
    return true;
}
// @lc code=end

