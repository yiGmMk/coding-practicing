/*
 * @lc app=leetcode.cn id=1233 lang=golang
 *
 * [1233] 删除子文件夹
 *
 * https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/description/
 *
 * algorithms
 * Medium (56.75%)
 * Likes:    101
 * Dislikes: 0
 * Total Accepted:    17.8K
 * Total Submissions: 31.5K
 * Testcase Example:  '["/a","/a/b","/c/d","/c/d/e","/c/f"]'
 *
 * 你是一位系统管理员，手里有一份文件夹列表 folder，你的任务是要删除该列表中的所有 子文件夹，并以 任意顺序 返回剩下的文件夹。
 *
 * 如果文件夹 folder[i] 位于另一个文件夹 folder[j] 下，那么 folder[i] 就是 folder[j] 的 子文件夹 。
 *
 * 文件夹的「路径」是由一个或多个按以下格式串联形成的字符串：'/' 后跟一个或者多个小写英文字母。
 *
 *
 * 例如，"/leetcode" 和 "/leetcode/problems" 都是有效的路径，而空字符串和 "/" 不是。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：folder = ["/a","/a/b","/c/d","/c/d/e","/c/f"]
 * 输出：["/a","/c/d","/c/f"]
 * 解释："/a/b" 是 "/a" 的子文件夹，而 "/c/d/e" 是 "/c/d" 的子文件夹。
 *
 *
 * 示例 2：
 *
 *
 * 输入：folder = ["/a","/a/b/c","/a/b/d"]
 * 输出：["/a"]
 * 解释：文件夹 "/a/b/c" 和 "/a/b/d" 都会被删除，因为它们都是 "/a" 的子文件夹。
 *
 *
 * 示例 3：
 *
 *
 * 输入: folder = ["/a/b/c","/a/b/ca","/a/b/d"]
 * 输出: ["/a/b/c","/a/b/ca","/a/b/d"]
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= folder.length <= 4 * 10^4
 * 2 <= folder[i].length <= 100
 * folder[i] 只包含小写字母和 '/'
 * folder[i] 总是以字符 '/' 起始
 * 每个文件夹名都是 唯一 的
 *
 *
 */
package jzoffer

import (
	"sort"
	"strings"
)

// @lc code=start
//  1. Sort the folders lexicographically.
//  2. Insert the current element in an array and then loop
//     until we get rid of all of their subfolders, repeat this until no element is left.
func removeSubfolders0(folder []string) (res []string) {
	sort.Strings(folder)
	res = append(res, folder[0])
	l := 0
	for i := 1; i < len(folder); i++ {
		if strings.HasPrefix(folder[i], folder[l]) {
			a, b := strings.Split(folder[i], "/"), strings.Split(folder[l], "/")
			if len(a) != len(b) { // 说明是子目录
				continue
			}
		}
		res = append(res, folder[i])
		l = i
	}
	return
}

// 作者：LeetCode-Solution
// 链接：https://leetcode.cn/problems/remove-sub-folders-from-the-filesystem/solution/shan-chu-zi-wen-jian-jia-by-leetcode-sol-0x8d/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func removeSubfolders1(folder []string) (ans []string) {
	sort.Strings(folder)
	ans = append(ans, folder[0])
	for _, f := range folder[1:] {
		last := ans[len(ans)-1]
		if !strings.HasPrefix(f, last) || f[len(last)] != '/' {
			ans = append(ans, f)
		}
	}
	return
}

type trie struct {
	ref     int
	childer map[string]*trie
}

func removeSubfolders(folder []string) (ans []string) {
	root := &trie{ref: -1, childer: map[string]*trie{}}
	for i, v := range folder {
		path := strings.Split(v, "/")
		cur := root
		for _, v := range path {
			if _, ok := cur.childer[v]; !ok {
				cur.childer[v] = &trie{ref: -1, childer: map[string]*trie{}}
			}
			cur = cur.childer[v]
		}
		cur.ref = i
	}
	var dfs func(node *trie)
	dfs = func(node *trie) {
		if node.ref != -1 {
			ans = append(ans, folder[node.ref])
			return
		}
		for _, v := range node.childer {
			dfs(v)
		}
	}
	dfs(root)
	return
}

// @lc code=end
