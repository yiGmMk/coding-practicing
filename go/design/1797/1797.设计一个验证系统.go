/*
* @lc app=leetcode.cn id=1797 lang=golang
*
* [1797] 设计一个验证系统
*
* https://leetcode.cn/problems/design-authentication-manager/description/
*
  - algorithms
  - Medium (58.09%)
  - Likes:    42
  - Dislikes: 0
  - Total Accepted:    15.1K
  - Total Submissions: 24.6K
  - Testcase Example:  '["AuthenticationManager","renew","generate","countUnexpiredTokens","generate","renew","renew","countUnexpiredTokens"]\n' +
    '[[5],["aaa",1],["aaa",2],[6],["bbb",7],["aaa",8],["bbb",10],[15]]'

*
* 你需要设计一个包含验证码的验证系统。每一次验证中，用户会收到一个新的验证码，这个验证码在 currentTime 时刻之后 timeToLive
* 秒过期。如果验证码被更新了，那么它会在 currentTime （可能与之前的 currentTime 不同）时刻延长 timeToLive 秒。
*
* 请你实现 AuthenticationManager 类：
*
*
* AuthenticationManager(int timeToLive) 构造 AuthenticationManager 并设置
* timeToLive 参数。
* generate(string tokenId, int currentTime) 给定 tokenId ，在当前时间 currentTime
* 生成一个新的验证码。
* renew(string tokenId, int currentTime) 将给定 tokenId 且 未过期 的验证码在 currentTime
* 时刻更新。如果给定 tokenId 对应的验证码不存在或已过期，请你忽略该操作，不会有任何更新操作发生。
* countUnexpiredTokens(int currentTime) 请返回在给定 currentTime 时刻，未过期 的验证码数目。
*
*
* 如果一个验证码在时刻 t 过期，且另一个操作恰好在时刻 t 发生（renew 或者 countUnexpiredTokens 操作），过期事件 优先于
* 其他操作。
*
*
*
* 示例 1：
*
*
* 输入：
* ["AuthenticationManager", "renew", "generate", "countUnexpiredTokens",
* "generate", "renew", "renew", "countUnexpiredTokens"]
* [[5], ["aaa", 1], ["aaa", 2], [6], ["bbb", 7], ["aaa", 8], ["bbb", 10],
* [15]]
* 输出：
* [null, null, null, 1, null, null, null, 0]
*
* 解释：
* AuthenticationManager authenticationManager = new AuthenticationManager(5);
* // 构造 AuthenticationManager ，设置 timeToLive = 5 秒。
* authenticationManager.renew("aaa", 1); // 时刻 1 时，没有验证码的 tokenId 为 "aaa"
* ，没有验证码被更新。
* authenticationManager.generate("aaa", 2); // 时刻 2 时，生成一个 tokenId 为 "aaa"
* 的新验证码。
* authenticationManager.countUnexpiredTokens(6); // 时刻 6 时，只有 tokenId 为 "aaa"
* 的验证码未过期，所以返回 1 。
* authenticationManager.generate("bbb", 7); // 时刻 7 时，生成一个 tokenId 为 "bbb"
* 的新验证码。
* authenticationManager.renew("aaa", 8); // tokenId 为 "aaa" 的验证码在时刻 7 过期，且 8
* >= 7 ，所以时刻 8 的renew 操作被忽略，没有验证码被更新。
* authenticationManager.renew("bbb", 10); // tokenId 为 "bbb" 的验证码在时刻 10
* 没有过期，所以 renew 操作会执行，该 token 将在时刻 15 过期。
* authenticationManager.countUnexpiredTokens(15); // tokenId 为 "bbb" 的验证码在时刻
* 15 过期，tokenId 为 "aaa" 的验证码在时刻 7 过期，所有验证码均已过期，所以返回 0 。
*
*
*
*
* 提示：
*
*
* 1
* 1
* 1
* tokenId 只包含小写英文字母。
* 所有 generate 函数的调用都会包含独一无二的 tokenId 值。
* 所有函数调用中，currentTime 的值 严格递增 。
* 所有函数的调用次数总共不超过 2000 次。
*
*
*/
package jzoffer

import "container/heap"

// @lc code=start
// 1. 哈希表: https://leetcode.cn/problems/design-authentication-manager/solution/she-ji-yi-ge-yan-zheng-xi-tong-by-leetco-kqqb/
type AuthenticationManager0 struct {
	t int
	m map[string]int
}

func Constructor0(timeToLive int) AuthenticationManager0 {
	return AuthenticationManager0{t: timeToLive, m: make(map[string]int)}
}

func (this *AuthenticationManager0) Generate(tokenId string, currentTime int) {
	this.m[tokenId] = currentTime + this.t
}

func (this *AuthenticationManager0) Renew(tokenId string, currentTime int) {
	if v, ok := this.m[tokenId]; ok && v > currentTime {
		this.m[tokenId] = currentTime + this.t
	}
}

func (this *AuthenticationManager0) CountUnexpiredTokens(currentTime int) (res int) {
	for _, v := range this.m {
		if v > currentTime {
			res++
		}
	}
	return
}

// 每次操作时删除超时的,使用优先队列按时间排序(小顶堆),超时越早的放前面
// ----------map+priorityqueue----------------
type (
	Node struct {
		token string
		t     int
		// index 是 entry 在 heap 中的索引号
		// entry 加入 Priority Queue 后， Priority 会变化时，很有用
		index int
	}
	PQ                    []*Node
	AuthenticationManager struct {
		t  int
		m  map[string]*Node
		pq *PQ
	}
)

//------priorityqueue---------------------

func (pq PQ) Len() int { return len(pq) }

func (h *PQ) Less(i, j int) bool {
	return (*h)[i].t < (*h)[j].t
}

func (pq PQ) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (h *PQ) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

func (h *PQ) Push(v interface{}) {
	if val, ok := v.(*Node); ok {
		val.index = len(*h)
		*h = append(*h, val)
	}
}

func Constructor(timeToLive int) AuthenticationManager {
	return AuthenticationManager{
		t:  timeToLive,
		m:  make(map[string]*Node),
		pq: &PQ{},
	}
}

func (this *AuthenticationManager) Remove(currentTime int) {
	for this.pq.Len() > 0 {
		if (*this.pq)[0].t > currentTime {
			return
		}
		v, ok := heap.Pop(this.pq).(*Node)
		if ok && v != nil {
			delete(this.m, v.token)
		}
	}
}

func (this *AuthenticationManager) Generate(tokenId string, currentTime int) {
	this.Remove(currentTime)
	node := &Node{token: tokenId, t: currentTime + this.t}
	heap.Push(this.pq, node)
	this.m[tokenId] = node
}

func (this *AuthenticationManager) Renew(tokenId string, currentTime int) {
	this.Remove(currentTime)
	if v, ok := this.m[tokenId]; !ok {
		return
	} else {
		v.t = currentTime + this.t
		heap.Fix(this.pq, v.index)
	}
}

func (this *AuthenticationManager) CountUnexpiredTokens(currentTime int) int {
	this.Remove(currentTime)
	return this.pq.Len()
}

/**
 * Your AuthenticationManager object will be instantiated and called as such:
 * obj := Constructor(timeToLive);
 * obj.Generate(tokenId,currentTime);
 * obj.Renew(tokenId,currentTime);
 * param_3 := obj.CountUnexpiredTokens(currentTime);
 */
// @lc code=end
