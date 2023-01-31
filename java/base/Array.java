package base;

import java.util.Arrays;

public class Array {

    public static void bubbleSort(int[] nums) {
        // 每一轮循环后，最大的一个数被交换到末尾，因此，下一轮循环就可以“刨除”最后的数，每一轮循环都比上一轮循环的结束位置靠前一位
        // System.out.println("begin:" + Arrays.toString(nums));
        for (int i = 0; i < nums.length - 1; i++) {
            for (int j = 0; j < nums.length - 1; j++) {
                if (nums[j] > nums[j + 1]) {
                    int tmp = nums[j];
                    nums[j] = nums[j + 1];
                    nums[j + 1] = tmp;
                }
            }
        }
        // System.out.println("end :" + Arrays.toString(nums));
    }

    public static int[][] arrays() {
        int[][] nums = {
                { 1 },
                { 10, 20, },
                { 5, 11, 6, }
        };
        return nums;
    }

    public static void main(String[] args) {
        int[] ns = { 111, 42, 9, 16, 25 };
        System.out.println(ns.toString());
        System.out.println("origin arrsy:" + Arrays.toString(ns));

        int[] copy = Arrays.copyOfRange(ns, 0, ns.length);
        Arrays.sort(ns);
        System.out.println("sorted array:" + Arrays.toString(ns) + "copy" + Arrays.toString(copy));

        System.out.println("\nbubble sort :" + Arrays.toString(ns));
        System.arraycopy(copy, 0, ns, 0, ns.length);
        bubbleSort(ns);
        System.out.println("sorted array:" + Arrays.toString(ns) + "copy" + Arrays.toString(copy));

        int[][] arrays = arrays();
        System.out.println("\narrays:" + Arrays.toString(arrays) + " deep: " + Arrays.deepToString(arrays));
    }
}
