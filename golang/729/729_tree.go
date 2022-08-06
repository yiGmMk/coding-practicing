package jzoffer

/*
题目的意思就是说两个数组之间不能有交叉的范围
而且0 <= start < end，即[start, end)
那我们可以自建一个二叉搜索树，依次插入树中
比较当前节点的max(注意max取不到，右边的数为开区间)是否小于等于给定的start，如果满足，往右子树找
比较当前节点的min是否大于等于给定的end(注意end取不到，右边的数为开区间)，如果满足，往左子树找
最后看是否能插入一个新节点
如果插入不了，返回false
方法二为红黑树
floor为向下取，即查找小于或等给定的参数
ceiling为向上取，即查找大于或等给定的参数

作者：XiaoWei_Algorithm
链接：https://leetcode.cn/problems/my-calendar-i/solution/by-xiaowei_algorithm-qmyv/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
type MyCalendar2 struct {
	root *node
}

func Constructor2() *MyCalendar2 {
	return &MyCalendar2{
		root: &node{
			min: -2,
			max: -1,
		},
	}
}

/*
				 min   max
	               -2 -1
	                   10 20
					0 10  20 30
*/
func (m *MyCalendar2) Book(start, end int) bool {
	cur := m.root
	for {
		if cur.min >= end {
			if cur.left == nil {
				cur.left = &node{min: start, max: end}
				return true
			}
			cur = cur.left
		} else if cur.max <= start {
			if cur.right == nil {
				cur.right = &node{min: start, max: end}
				return true
			}
			cur = cur.right
		} else {
			return false
		}
	}
}
