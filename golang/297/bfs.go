package jzoffer

import (
	"strconv"
	"strings"
)

/*
BFS解题分析
这道题使用BFS的解题更为简单写，默认BFS时，需要判断当前节点是否有左、右子树
而此时，我们无需判断左右子树，当node为None时添加一个None到队列中即可。
DFS解题分析
在使用DFS时，势必将根节点放在第一个位置，可以方便我们反序列化，所以应该使用前序遍历。
遍历的时候同样注意，如果遇到空节点，需要将为空的节点也记录下来
即当一个子节点的左右节点都是None时，表示它其实是一个叶子节点。
反序列化时，同样通过前序遍历来实现，但注意默认的字符串转化为列表后，我们应该将从左到右遍历才满足条件
Python我们可以反转列表或者使用deque的双端队列
Java我们可以链表来实现
具体BFS、DFS解题如下：
---------------bfs------------
class Codec:
    def serialize(self, root):
        if not root:
            return ""
        dq = deque([root])
        res = []
        while dq:
            node = dq.popleft()
            if node:
                res.append(str(node.val))
                dq.append(node.left)
                dq.append(node.right)
            else:
                res.append('None')
        return ','.join(res)

    def deserialize(self, data):
        if not data:
            return []
        dataList = data.split(',')
        root = TreeNode(int(dataList[0]))
        dq = deque([root])
        i = 1
        while dq:
            node = dq.popleft()
            if dataList[i] != 'None':
                node.left = TreeNode(int(dataList[i]))
                dq.append(node.left)
            i += 1
            if dataList[i] != 'None':
                node.right = TreeNode(int(dataList[i]))
                dq.append(node.right)
            i += 1
        return root
-------------------DFS解题--------------------------
from collections import deque
class Codec:
    def serialize(self, root):
        if not root:
            return 'None'
        root.left = self.serialize(root.left)
        root.right = self.serialize(root.right)
        return f"{root.val},{root.left},{root.right}"

    def deserialize(self, data):
        dq = deque(data.split(','))
        def dfs(li):
            val = li.popleft()
            if val == "None":
                return None
            root = TreeNode(int(val))
            root.left = dfs(li)
            root.right = dfs(li)
            return root
        return dfs(dq)
*/
type BTreeCodec interface {
	Serialize(root *TreeNode) string
	Deserialize(data string) *TreeNode
}

type BTreeBfsCodec struct {
}

func (bfs BTreeBfsCodec) Serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	queue := []*TreeNode{root}
	res := []string{}
	ql := len(queue)
	for ql > 0 {
		null := 0
		notNull := 0
		for i := 0; i < ql; i++ {
			node := queue[0]
			queue = queue[1:]
			if node != nil {
				res = append(res, strconv.Itoa(node.Val))
				queue = append(queue, node.Left)
				queue = append(queue, node.Right)
				notNull++
			} else {
				res = append(res, "null")
				null++
			}
		}
		if null == ql && notNull == 0 {
			res = res[:len(res)-null]
		}
		ql = len(queue)
	}
	return strings.Join(res, ",")
}

func (bfs BTreeBfsCodec) Deserialize(data string) *TreeNode {
	strs := strings.Split(data, ",")
	if len(strs) == 0 {
		return nil
	}
	root, _ := bfs.newNode(strs[0])
	queue := []*TreeNode{root}
	i, ql := 1, len(queue)
	for ql > 0 && i < len(strs) {
		node := queue[0]
		queue = queue[1:]
		if strs[i] != "null" {
			node.Left, _ = bfs.newNode(strs[i])
			queue = append(queue, node.Left)
		}
		i += 1
		if strs[i] != "null" {
			node.Right, _ = bfs.newNode(strs[i])
			queue = append(queue, node.Right)
		}
		i += 1
		ql = len(queue)
	}
	return root
}

func (bfs BTreeBfsCodec) newNode(val string) (*TreeNode, error) {
	iVal, err := strconv.Atoi(val)
	if err != nil {
		return nil, err
	}
	return &TreeNode{Val: iVal}, nil
}
