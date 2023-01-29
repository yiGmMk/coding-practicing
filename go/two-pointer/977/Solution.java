import java.util.Arrays;

public class Solution {
    public static int[] sortedSquares(int[] nums) {
        int l = 0;
        int r = nums.length - 1;
        int[] res = new int[nums.length];
        for (int pos = r; l <= r; pos--) {
            if (nums[l] * nums[l] > nums[r] * nums[r]) {
                res[pos] = nums[l] * nums[l];
                l++;
            } else {
                res[pos] = nums[r] * nums[r];
                r--;
            }
        }
        return res;
    }

    public static int[] sortedSquares1(int[] nums) {
        int[] res = new int[nums.length];
        for (int i = 0; i < nums.length; i++) {
            res[i] = nums[i] * nums[i];
        }
        Arrays.sort(res);
        return res;
    }

    public static void main(String[] args) {
        int[] nums = { -4, -1, 0, 3, 10 };
        int[] res = sortedSquares(nums);
        for (int i = 0; i < res.length; i++) {
            System.out.print(res[i]);
        }
    }
}
