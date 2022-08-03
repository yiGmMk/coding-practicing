/*
 * @lc app=leetcode.cn id=211 lang=golang
 *
 * [211] 添加与搜索单词 - 数据结构设计
 *
 * https://leetcode.cn/problems/design-add-and-search-words-data-structure/description/
 *
 * algorithms
 * Medium (50.40%)
 * Likes:    446
 * Dislikes: 0
 * Total Accepted:    62.6K
 * Total Submissions: 124.2K
 * Testcase Example:  '["WordDictionary","addWord","addWord","addWord","search","search","search","search"]\n' +
  '[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]'
 *
 * 请你设计一个数据结构，支持 添加新单词 和 查找字符串是否与任何先前添加的字符串匹配 。
 *
 * 实现词典类 WordDictionary ：
 *
 *
 * WordDictionary() 初始化词典对象
 * void addWord(word) 将 word 添加到数据结构中，之后可以对它进行匹配
 * bool search(word) 如果数据结构中存在字符串与 word 匹配，则返回 true ；否则，返回  false 。word 中可能包含一些
 * '.' ，每个 . 都可以表示任何一个字母。
 *
 *
 *
 *
 * 示例：
 *
 *
 * 输入：
 *
 * ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]
 * [[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
 * 输出：
 * [null,null,null,null,false,true,true,true]
 *
 * 解释：
 * WordDictionary wordDictionary = new WordDictionary();
 * wordDictionary.addWord("bad");
 * wordDictionary.addWord("dad");
 * wordDictionary.addWord("mad");
 * wordDictionary.search("pad"); // 返回 False
 * wordDictionary.search("bad"); // 返回 True
 * wordDictionary.search(".ad"); // 返回 True
 * wordDictionary.search("b.."); // 返回 True
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= word.length <= 25
 * addWord 中的 word 由小写英文字母组成
 * search 中的 word 由 '.' 或小写英文字母组成
 * 最多调用 10^4 次 addWord 和 search
 *
 *
*/

package jzoffer

// @lc code=start

type trie struct {
	children [26]*trie
	istail   bool
}
type WordDictionary struct {
	root *trie
}

func Constructor() WordDictionary {
	return WordDictionary{root: &trie{children: [26]*trie{}}}
}

func (this *WordDictionary) AddWord(word string) {
	root := this.root
	for _, c := range word {
		if root.children[c-'a'] == nil {
			root.children[c-'a'] = &trie{children: [26]*trie{}}
		}
		root = root.children[c-'a']
	}
	root.istail = true
}

// 偶尔能过,偶尔超时
func (this *WordDictionary) SearchTimeOut(word string) bool {
	root := this.root
	nodes := []*trie{root}
	for _, c := range word {
		cur := []*trie{}
		for _, node := range nodes {
			if c != '.' {
				if node.children[c-'a'] != nil {
					cur = append(cur, node.children[c-'a'])
				}
				continue
			}
			for _, child := range node.children {
				if child != nil {
					cur = append(cur, child)
				}
			}
		}
		nodes = cur
	}
	for _, node := range nodes {
		if node.istail {
			return true
		}
	}
	return false
}

// dfs
func (this *WordDictionary) Search(word string) bool {
	var dfs func(int, *trie) bool
	dfs = func(index int, node *trie) bool {
		if index == len(word) {
			return node.istail
		}
		ch := word[index]
		if ch != '.' {
			child := node.children[ch-'a']
			if child != nil && dfs(index+1, child) {
				return true
			}
		} else {
			for i := range node.children {
				child := node.children[i]
				if child != nil && dfs(index+1, child) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0, this.root)

	// 作者：LeetCode-Solution
	// 链接：https://leetcode.cn/problems/design-add-and-search-words-data-structure/solution/tian-jia-yu-sou-suo-dan-ci-shu-ju-jie-go-n4ud/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end
