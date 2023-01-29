public class Solution extends VersionControl {
    public int firstBadVersion(int n) {
        int left=0,right=n;
        while (left+1<right){
            int mid=left+(right-left)/2;
            if (isBadVersion(mid)) {
                right=mid;
            }else{
                left=mid;
        }
    }
    return  right;
    }
}