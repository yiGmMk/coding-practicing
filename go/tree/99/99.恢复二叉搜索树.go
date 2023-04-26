/*
 * @lc app=leetcode.cn id=99 lang=golang
 *
 * [99] 恢复二叉搜索树
 *
 * https://leetcode.cn/problems/recover-binary-search-tree/description/
 *
 * algorithms
 * Medium (60.27%)
 * Likes:    855
 * Dislikes: 0
 * Total Accepted:    133.1K
 * Total Submissions: 220.8K
 * Testcase Example:  '[1,3,null,null,2]'
 *
 * 给你二叉搜索树的根节点 root ，该树中的 恰好 两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树 。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：root = [1,3,null,null,2]
 * 输出：[3,1,null,null,2]
 * 解释：3 不能是 1 的左孩子，因为 3 > 1 。交换 1 和 3 使二叉搜索树有效。
 *
 *
 * 示例 2：
 *
 *
 * 输入：root = [3,1,4,null,null,2]
 * 输出：[2,1,4,null,null,3]
 * 解释：2 不能在 3 的右子树中，因为 2 < 3 。交换 2 和 3 使二叉搜索树有效。
 *
 *
 *
 * 提示：
 *
 *
 * 树上节点的数目在范围 [2, 1000] 内
 * -2^31 <= Node.val <= 2^31 - 1
 *
 *
 *
 *
 * 进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用 O(1) 空间的解决方案吗？
 *
 */
package jzoffer

import (
	"math"

	"github.com/yiGmMk/leetcode/golang/util"
)

type TreeNode = util.TreeNode

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// TODO:
/*
显示中序遍历:
func recoverTree(root *TreeNode)  {
    nums := []int{}
    var inorder func(node *TreeNode)
    inorder = func(node *TreeNode) {
        if node == nil {
            return
        }
        inorder(node.Left)
        nums = append(nums, node.Val)
        inorder(node.Right)
    }
    inorder(root)
    x, y := findTwoSwapped(nums)
    recover(root, 2, x, y)
}

func findTwoSwapped(nums []int) (int, int) {
    index1, index2 := -1, -1
    for i := 0; i < len(nums) - 1; i++ {
        if nums[i + 1] < nums[i] {
            index2 = i + 1
            if index1 == -1 {
                index1 = i
            } else {
                break
            }
        }
    }
    x, y := nums[index1], nums[index2]
    return x, y
}

func recover(root *TreeNode, count, x, y int) {
    if root == nil {
        return
    }
    if root.Val == x || root.Val == y {
        if root.Val == x {
            root.Val = y
        } else {
            root.Val = x
        }
        count--
        if count == 0 {
            return
        }
    }
    recover(root.Right, count, x, y)
    recover(root.Left, count, x, y)
}
*/

func recoverTree1(root *TreeNode) {
	var nodes []*TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		nodes = append(nodes, node)
		dfs(node.Right)
	}
	// 中序遍历=> 从小到大排序
	dfs(root)
	// 找出逆序的节点
	pre, cur := nodes[0].Val, math.MinInt
	var n1, n2 *TreeNode
	for i := 1; i < len(nodes); i++ {
		cur = nodes[i].Val
		if cur < pre {
			n1 = nodes[i]
			if n2 == nil {
				n2 = nodes[i-1]
			} else {
				break
			}
			//	fmt.Println("found val:", cur, pre)
		}
		//  fmt.Println(i, ":", pre, cur)
		pre = cur

	}
	// 恢复
	if n1 != nil && n2 != nil {
		n1.Val, n2.Val = n2.Val, n1.Val
	}
	return
}

func recoverTree(root *TreeNode) {
	nodes := []*TreeNode{}
	swap := func(x, y *TreeNode) {
		x.Val, y.Val = y.Val, x.Val
	}
	var x, y, pred *TreeNode
	for len(nodes) > 0 || root != nil {
		for root != nil {
			nodes = append(nodes, root)
			root = root.Left
		}
		size := len(nodes)
		root = nodes[size-1]
		nodes = nodes[:size-1]
		if pred != nil && root.Val < pred.Val {
			y = root
			if x == nil {
				x = pred
			} else {
				break
			}
		}
		pred = root
		root = root.Right
	}
	swap(x, y)
}

/*
class Solution {
    public void recoverTree(TreeNode root) {
        Deque<TreeNode> stack = new ArrayDeque<TreeNode>();
        TreeNode x = null, y = null, pred = null;

        while (!stack.isEmpty() || root != null) {
            while (root != null) {
                stack.push(root);
                root = root.left;
            }
            root = stack.pop();
            if (pred != null && root.val < pred.val) {
                y = root;
                if (x == null) {
                    x = pred;
                } else {
                    break;
                }
            }
            pred = root;
            root = root.right;
        }

        swap(x, y);
    }

    public void swap(TreeNode x, TreeNode y) {
        int tmp = x.val;
        x.val = y.val;
        y.val = tmp;
    }
}

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/recover-binary-search-tree/solution/hui-fu-er-cha-sou-suo-shu-by-leetcode-solution/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/

// @lc code=end
