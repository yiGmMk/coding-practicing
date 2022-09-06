/*
* @lc app=leetcode.cn id=380 lang=golang
*
* [380] O(1) 时间插入、删除和获取随机元素
*
* https://leetcode.cn/problems/insert-delete-getrandom-o1/description/
*
  - algorithms
  - Medium (52.81%)
  - Likes:    596
  - Dislikes: 0
  - Total Accepted:    90K
  - Total Submissions: 170.4K
  - Testcase Example:  '["RandomizedSet","insert","remove","insert","getRandom","remove","insert","getRandom"]\n' +
    '[[],[1],[2],[2],[],[1],[2],[]]'

*
* 实现RandomizedSet 类：
*
*
*
*
* RandomizedSet() 初始化 RandomizedSet 对象
* bool insert(int val) 当元素 val 不存在时，向集合中插入该项，并返回 true ；否则，返回 false 。
* bool remove(int val) 当元素 val 存在时，从集合中移除该项，并返回 true ；否则，返回 false 。
* int getRandom() 随机返回现有集合中的一项（测试用例保证调用此方法时集合中至少存在一个元素）。每个元素应该有 相同的概率 被返回。
*
*
* 你必须实现类的所有函数，并满足每个函数的 平均 时间复杂度为 O(1) 。
*
*
*
* 示例：
*
*
* 输入
* ["RandomizedSet", "insert", "remove", "insert", "getRandom", "remove",
* "insert", "getRandom"]
* [[], [1], [2], [2], [], [1], [2], []]
* 输出
* [null, true, false, true, 2, true, false, 2]
*
* 解释
* RandomizedSet randomizedSet = new RandomizedSet();
* randomizedSet.insert(1); // 向集合中插入 1 。返回 true 表示 1 被成功地插入。
* randomizedSet.remove(2); // 返回 false ，表示集合中不存在 2 。
* randomizedSet.insert(2); // 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
* randomizedSet.getRandom(); // getRandom 应随机返回 1 或 2 。
* randomizedSet.remove(1); // 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
* randomizedSet.insert(2); // 2 已在集合中，所以返回 false 。
* randomizedSet.getRandom(); // 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
*
*
*
*
* 提示：
*
*
* -2^31 <= val <= 2^31 - 1
* 最多调用 insert、remove 和 getRandom 函数 2 * 10^5 次
* 在调用 getRandom 方法时，数据结构中 至少存在一个 元素。
*
*
*
*
*/
package jzoffer

// @lc code=start
import "math/rand"

// 没有完全实现O(1),查询需要优化
type RandomizedSet1 struct {
	im map[int]struct{}
}

func Constructor1() RandomizedSet1 {
	m := make(map[int]struct{})
	return RandomizedSet1{im: m}
}

func (this *RandomizedSet1) Insert(val int) bool {
	if _, ok := this.im[val]; ok {
		return false
	}
	this.im[val] = struct{}{}
	return true
}

func (this *RandomizedSet1) Remove(val int) bool {
	if _, ok := this.im[val]; ok {
		delete(this.im, val)
		return true
	}

	return false
}

func (this *RandomizedSet1) GetRandom() int {
	num := len(this.im)
	if num == 0 {
		return 0
	}

	i, s := 0, rand.Int()%num
	for k := range this.im {
		if i == s {
			return k
		}
		i++
	}

	return 0
}

// /---------------------optimize,随机数O(1)-------------------------------
// 引入?实现O(1)获取随机数,(数组/链表/堆/栈,最先想到数组,很容易实现O(1))
type RandomizedSet struct {
	im   map[int]int
	nums []int
}

func Constructor() RandomizedSet {
	m := make(map[int]int)
	return RandomizedSet{im: m}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.im[val]; ok {
		return false
	}
	this.nums = append(this.nums, val)
	this.im[val] = len(this.nums) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	if i, ok := this.im[val]; ok {
		if len(this.nums) > 1 {
			v := this.nums[len(this.nums)-1] // 最后一个移动到i的位置上,改下slice
			this.nums[i] = v
			this.nums = this.nums[:len(this.nums)-1]
			this.im[v] = i
		} else {
			this.nums = []int{}
		}
		delete(this.im, val)
		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	num := len(this.im)
	if num == 0 {
		return 0
	}

	i := rand.Int() % num
	return this.nums[i]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
// @lc code=end
