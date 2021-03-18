#include<unordered_set>
#include<vector>
#include<memory.h>
#include<iostream>
#include<string>
#include<queue>
#include<algorithm>
#include<map>

using namespace std;
struct TreeNode {
      int val;
      TreeNode *left;
      TreeNode *right;
      TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};