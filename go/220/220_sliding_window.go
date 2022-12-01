package jzoffer

/* abs(nums[i] - nums[j])<=t,同时又满足 abs(i - j)<=k  。
 * 滑动窗口 + 有序集合 实现
 * 维护一个大小为k的滑动窗口,每次看是否存在一个元素在[x-t,x+t]范围内,存在则返回true
 */

/*
class Solution {
public:
    bool containsNearbyAlmostDuplicate(vector<int>& nums, int k, int t) {
        int n = nums.size();
        set<int> rec;
        for (int i = 0; i < n; i++) {
            auto iter = rec.lower_bound(max(nums[i], INT_MIN + t) - t);
            if (iter != rec.end() && *iter <= min(nums[i], INT_MAX - t) + t) {
                return true;
            }
            rec.insert(nums[i]);
            if (i >= k) {
                rec.erase(nums[i - k]);
            }
        }
        return false;
    }
};

作者：LeetCode-Solution
链接：https://leetcode.cn/problems/7WqeDu/solution/zhi-he-xia-biao-zhi-chai-du-zai-gei-ding-94ei/
来源：力扣（LeetCode）
著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
*/
import (
	"math"

	"github.com/duke-git/lancet/v2/lancetconstraints"
	"github.com/liyue201/gostl/ds/set"
)

func max[T lancetconstraints.Number](x T, y T) T {
	if x > y {
		return x
	}
	return y
}

func min[T lancetconstraints.Number](x T, y T) T {
	if x < y {
		return x
	}
	return y
}

func containsNearbyAlmostDuplicateSlidingWindow(nums []int, k int, t int) bool {
	s := set.New(set.WithGoroutineSafe())
	for i, v := range nums {
		iter := s.LowerBound(max(v, math.MinInt+t) - t)
		if iter != s.Last() &&
			iter.IsValid() &&
			iter.Value().(int) <= min(v, math.MaxInt-t)+t {
			return true
		}

		s.Insert(v)
		if i >= k {
			s.Erase(nums[i-k])
		}
	}
	return false
}
