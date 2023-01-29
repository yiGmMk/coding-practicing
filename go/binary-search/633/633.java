class Solution {
    public boolean judgeSquareSum(int c) {
        long l = 0;
        long r = (long) Math.sqrt(c);
        while (l <= r) {
            long v = l * l + r * r;
            if (v > c) {
                r--;
            } else if (v == c) {
                return true;
            } else {
                l++;
            }
        }
        return false;
    }
}
