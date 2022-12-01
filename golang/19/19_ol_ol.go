package jzoffer

/* 头结点
 * 尾结点
 * 中间节点
 */
func removeNthFromEndOLOL(head *ListNode, n int) *ListNode {
	ns := []*ListNode{}
	p := head
	for p != nil {
		ns = append(ns, p)
		p = p.Next
	}
	num := len(ns)
	if num <= 1 {
		return nil
	}
	if n == num { //删除头结点
		if len(ns) > 1 {
			return ns[1]
		}
	} else if n == 1 { //尾结点
		if num > 1 {
			ns[num-2].Next = nil // 倒数第二个
			return ns[0]
		}
	} else { //删中间的
		// len=5 n=2  1 2 3 4 5
		// len=2 n=1  1 2
		ns[num-n-1].Next = ns[num-n+1]
		return ns[0]
	}
	return p
}
