import java.util.Deque;
import java.util.LinkedList;

// https://zhuanlan.zhihu.com/p/414104829
// main是静态方法,静态方法是不能直接访问非静态类的。故ListNode需要是静态内部类或改为外部类

class ListNode {
    int val;
    ListNode next;

    ListNode() {
    }

    ListNode(int val) {
        this.val = val;
    }

    ListNode(int val, ListNode next) {
        this.val = val;
        this.next = next;
    }

    public String toString() {
        StringBuilder sb = new StringBuilder();
        sb.append("[");
        sb.append(this.val);
        for (ListNode p = this.next; p != null; p = p.next) {
            sb.append(",");
            sb.append(p.val);
        }
        sb.append("]");
        return sb.toString();
    }
}

class Solution19 {
    // static class ListNode {
    // int val;
    // ListNode next;

    // ListNode() {
    // }

    // ListNode(int val) {
    // this.val = val;
    // }

    // ListNode(int val, ListNode next) {
    // this.val = val;
    // this.next = next;
    // }
    // }

    public static ListNode removeNthFromEnd(ListNode head, int n) {
        if (head == null) {
            return head;
        }
        ListNode dummy = new ListNode(0, head);
        // Queue是队列，只能一头进，另一头出。
        // 如果把条件放松一下，允许两头都进，两头都出，这种队列叫双端队列（Double Ended Queue），学名Deque
        Deque<ListNode> stack = new LinkedList<ListNode>();
        ListNode cur = dummy;
        while (cur != null) {
            stack.push(cur);
            cur = cur.next;
        }
        for (int i = 0; i < n; ++i) {
            stack.pop();
        }
        ListNode prev = stack.peek();
        prev.next = prev.next.next;
        ListNode ans = dummy.next;
        return ans;
    }

    public static void main(String[] args) {
        ListNode head = new ListNode(1, new ListNode(2, new ListNode(3, new ListNode(4, new ListNode(5)))));
        System.out.println("list:" + head.toString());
        head = removeNthFromEnd(head, 2);
        System.out.println("list,remove 2nd:" + head.toString());
        head = removeNthFromEnd(head, 2);
        System.out.println("list,remove 2nd:" + head.toString());
        head = removeNthFromEnd(head, 2);
        System.out.println("list,remove 2nd:" + head.toString());
        head = removeNthFromEnd(head, 2);
        System.out.println("list,remove 2nd:" + head.toString());
    }
}

// 作者：LeetCode-Solution 链接：https://
// leetcode.cn/problems/remove-nth-node-from-end-of-list/solution/shan-chu-lian-biao-de-dao-shu-di-nge-jie-dian-b-61/
// 来源：力扣（LeetCode）著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。