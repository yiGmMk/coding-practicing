package jzoffer

// 基于桶排序
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

/* abs(nums[i] - nums[j])<=t,同时又满足 abs(i - j)<=k  。
我们也可以使用利用桶排序的思想解决本题。我们按照元素的大小进行分桶，维护一个滑动窗口内的元素对应的元素。
对于元素 x，其影响的区间为 [x - t, x + t]。于是我们可以设定桶的大小为 t + 1。
如果两个元素同属一个桶，那么这两个元素必然符合条件。
如果两个元素属于相邻桶，那么我们需要校验这两个元素是否差值不超过 t。
如果两个元素既不属于同一个桶，也不属于相邻桶，那么这两个元素必然不符合条件。
具体地，我们遍历该序列，假设当前遍历到元素 x，那么我们首先检查 x 所属于的桶是否已经存在元素，
如果存在，那么我们就找到了一对符合条件的元素，否则我们继续检查两个相邻的桶内是否存在符合条件的元素。
*/
func containsNearbyAlmostDuplicateBucket(nums []int, k, t int) bool {
	mp := map[int]int{}
	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x-y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x-y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/7WqeDu/solution/zhi-he-xia-biao-zhi-chai-du-zai-gei-ding-94ei/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
