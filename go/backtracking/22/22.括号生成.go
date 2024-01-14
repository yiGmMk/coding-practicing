/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 *
 * https://leetcode.cn/problems/generate-parentheses/description/
 *
 * algorithms
 * Medium (77.55%)
 * Likes:    3483
 * Dislikes: 0
 * Total Accepted:    780.3K
 * Total Submissions: 1M
 * Testcase Example:  '3'
 *
 * 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入：n = 3
 * 输出：["((()))","(()())","(())()","()(())","()()()"]
 *
 *
 * 示例 2：
 *
 *
 * 输入：n = 1
 * 输出：["()"]
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= n <= 8
 *
 *
 */
package jzoffer

// @lc code=start

func generateParenthesisDFS(n int) (res []string) {
	var dfs func(path string, l, r int)
	dfs = func(path string, l, r int) {
		if l == r && l == n {
			res = append(res, path)
			return
		}
		if l < r {
			return
		}
		if l < n {
			dfs(path+"(", l+1, r)
		}
		if r < n {
			dfs(path+")", l, r+1)
		}
		return
	}

	dfs("", 0, 0)
	return
}

// 每个位置枚举 所有情况, 实际每个位置就2中情况是'(' 或者 ')',不剪支会超时
func generateParenthesisEnumDfs(n int) (res []string) {
	bytes := make([]byte, n*2)
	for i := 0; i < n; i++ {
		bytes[i*2] = '('
		bytes[i*2+1] = ')'
	}

	check := func(s string) bool {
		l, r := 0, 0
		for _, v := range s {
			if v == '(' {
				l++
			} else {
				r++
			}
			if l < r {
				return false
			}
		}
		return true
	}

	var dfs func(path []byte, used map[int]bool, idx int)
	has := make(map[string]bool)
	dfs = func(path []byte, used map[int]bool, idx int) {
		if !check(string(path[:idx])) {
			return
		}
		if idx == n*2-1 {
			if check(string(path)) {
				if has[string(path)] {
					return
				}
				has[string(path)] = true
				res = append(res, string(path))
			}
			return
		}
		l, r := false, false
		for i := 0; i < len(bytes); i++ {
			if used[i] {
				continue
			}
			if bytes[i] == '(' {
				if l {
					continue
				}
				l = true
			}
			if bytes[i] == ')' {
				if r {

					continue
				}
				r = true
			}
			path[idx] = bytes[i]
			used[i] = true
			dfs(path, used, idx+1)
			used[i] = false
		}
	}
	used := make(map[int]bool)
	dfs(append([]byte{}, bytes...), used, 0)
	return
}

func generateParenthesis(n int) (res []string) {
	bytes := make([]byte, n*2)
	for i := 0; i < n; i++ {
		bytes[i*2] = '('
		bytes[i*2+1] = ')'
	}

	check := func(s string) bool {
		l, r := 0, 0
		for _, v := range s {
			if v == '(' {
				l++
			} else {
				r++
			}
			if l < r {
				return false
			}
		}
		return true
	}

	var dfs func(path []byte, l, r, idx int)
	has := make(map[string]bool)
	dfs = func(path []byte, l, r, idx int) {
		if idx == n*2-1 {
			if check(string(path)) {
				if has[string(path)] {
					return
				}
				has[string(path)] = true
				res = append(res, string(path))
			}
			return
		}
		if l < n {
			path[idx] = '('
			dfs(path, l+1, r, idx+1)
		}

		if r < n {
			path[idx] = ')'
			dfs(path, l, r+1, idx+1)
		}
	}
	dfs(append([]byte{}, bytes...), 0, 0, 0)
	return
}

// @lc code=end
