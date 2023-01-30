import java.util.Arrays;

class Solution189 {
    public void rotate(int[] nums, int k) {
        int n = nums.length;
        int[] newArr = new int[n];
        for (int i = 0; i < n; i++) {
            newArr[(i + k) % n] = nums[i];
        }
        System.arraycopy(newArr, 0, nums, 0, n);
    }

    public static void rotate1(int[] nums, int k) {
        k = k % nums.length;
        reverse(nums, 0, nums.length - 1);
        reverse(nums, 0, k - 1);
        reverse(nums, k, nums.length - 1);
    }

    public static void reverse(int[] nums, int l, int r) {
        int tmp;
        while (l < r) {
            tmp = nums[l];
            nums[l] = nums[r];
            nums[r] = tmp;
            l++;
            r--;
        }
    }

    public static void main(String[] args) {
        class TestCase {
            int[] nums;
            int k;
            int[] expect;

            public TestCase(int[] nums, int k, int[] expect) {
                this.nums = nums;
                this.k = k;
                this.expect = expect;
            }
        }
        TestCase[] tcs = {
                new TestCase(
                        new int[] { 1, 2, 3, 4, 5, 6, 7 },
                        3,
                        new int[] { 5, 6, 7, 1, 2, 3, 40 }) };

        for (TestCase v : tcs) {
            rotate1(v.nums, v.k);
            if (!Arrays.equals(v.nums, v.expect)) {
                System.out.printf("rotate err:%s", Arrays.toString(v.nums));
            }
        }
    }
}