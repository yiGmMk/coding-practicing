/*
 * @lc app=leetcode.cn id=236 lang=cpp
 *
 * [236] 二叉树的最近公共祖先
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
// class Solution {
// public:
//     set<int> treeNodeSet;
//     vector<int> nodeVec;
//     //1.遍历二叉树，二叉树的值加入set，每一棵子树构架一个set查找是否存在p，q
//     void BinarySearch(TreeNode *root,set<int> & nodeSet){
//         if(!root)
//             return;
//         nodeSet.insert(root->val);
//         if(root->left)
//             BinarySearch(root->left,nodeSet);
//         if(root->right)
//             BinarySearch(root->right,nodeSet);
//     }

//     int BinarySearch(TreeNode *root,int p,int q){
//         if(!root)
//             return -1;
//         treeNodeSet.clear();
//         treeNodeSet.insert(root->val);
//         BinarySearch(root->left,treeNodeSet);
//         BinarySearch(root->right,treeNodeSet);
//         if(treeNodeSet.count(p)&&treeNodeSet.count(q))
//             return root->val;
//         else
//             return -1;
//     }

//      void BinarySearch(TreeNode *root,int p,int q,vector<int> & vec){
//         if(!root)
//             return;
//         vec.push_back(BinarySearch(root,p,q));
//         if(root->left)
//             BinarySearch(root->left,p,q,vec);
//         if(root->right)
//             BinarySearch(root->right,p,q,vec);
//     }
//     TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
//         BinarySearch(root,p->val,q->val,nodeVec);
//         sort(nodeVec.begin(),nodeVec.end());
//         TreeNode *pRes=new TreeNode(nodeVec.back());
//         return pRes;
//     }
// };

#include"leetcode.h"
class Solution {
public:
    
    TreeNode* ans;
    bool dfs(TreeNode* root, TreeNode* p, TreeNode* q) {
        if (root == nullptr) return false;
        bool lson = dfs(root->left, p, q);
        bool rson = dfs(root->right, p, q);
        if ((lson && rson) || ((root->val == p->val || root->val == q->val) && (lson || rson)))         {
            ans = root;
        } 
        return lson || rson || (root->val == p->val || root->val == q->val);
    }
    TreeNode* lowestCommonAncestor(TreeNode* root, TreeNode* p, TreeNode* q) {
        dfs(root, p, q);
        return ans;
    }
};
// @lc code=end

