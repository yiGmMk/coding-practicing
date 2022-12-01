package jzoffer

import (
	"strconv"
	"strings"
)

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/serialize-and-deserialize-binary-tree/solution/er-cha-shu-de-xu-lie-hua-yu-fan-xu-lie-hua-by-le-2/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

type CodecDfs struct{}

func (dfs CodecDfs) Serialize(root *TreeNode) string {
	sb := &strings.Builder{}
	var fDfs func(*TreeNode)
	fDfs = func(node *TreeNode) {
		if node == nil {
			sb.WriteString("null,")
			return
		}
		sb.WriteString(strconv.Itoa(node.Val))
		sb.WriteByte(',')
		fDfs(node.Left)
		fDfs(node.Right)
	}
	fDfs(root)
	return sb.String()
}

func (dfs CodecDfs) Deserialize(data string) *TreeNode {
	sp := strings.Split(data, ",")
	var build func() *TreeNode
	build = func() *TreeNode {
		if len(sp) == 0 {
			return nil
		}
		if sp[0] == "null" {
			sp = sp[1:]
			return nil
		}
		val, _ := strconv.Atoi(sp[0])
		sp = sp[1:]
		return &TreeNode{Val: val, Left: build(), Right: build()}
	}
	return build()
}
