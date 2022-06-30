package jzoffer

func inorderSuccessor3(root, p *TreeNode) *TreeNode {
	cur, result := root, (*TreeNode)(nil)
	for cur != nil {
		if cur.Val > p.Val {
			result = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return result
}

// 作者：angela-x
// 链接：https://leetcode.cn/problems/P5rCT8/solution/shi-jian-fu-za-du-wei-ceng-gao-de-jie-fa-fheg/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
