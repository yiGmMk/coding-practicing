
public class Solution {
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
    }

    public ListNode mergeInBetween(ListNode list1, int a, int b, ListNode list2) {
        ListNode preA = list1;
        ListNode aftB = list1;
        for (int i = 0; i <= b; i++) {
            if (i < a - 1) {
                preA = preA.next;
            }
            aftB = aftB.next;
        }
        preA.next = list2;
        ListNode p = list2;
        for (; p != null && p.next != null; p = p.next) {
        }
        p.next = aftB;
        return list1;
    }
}
