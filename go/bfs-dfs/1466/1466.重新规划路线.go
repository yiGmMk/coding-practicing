/*
 * @lc app=leetcode.cn id=1466 lang=golang
 *
 * [1466] 重新规划路线
 *
 * https://leetcode.cn/problems/reorder-routes-to-make-all-paths-lead-to-the-city-zero/description/
 *
 * algorithms
 * Medium (51.13%)
 * Likes:    147
 * Dislikes: 0
 * Total Accepted:    17.3K
 * Total Submissions: 33.8K
 * Testcase Example:  '6\n[[0,1],[1,3],[2,3],[4,0],[4,5]]'
 *
 * n 座城市，从 0 到 n-1 编号，其间共有 n-1
 * 条路线。因此，要想在两座不同城市之间旅行只有唯一一条路线可供选择（路线网形成一颗树）。去年，交通运输部决定重新规划路线，以改变交通拥堵的状况。
 *
 * 路线用 connections 表示，其中 connections[i] = [a, b] 表示从城市 a 到 b 的一条有向路线。
 *
 * 今年，城市 0 将会举办一场大型比赛，很多游客都想前往城市 0 。
 *
 * 请你帮助重新规划路线方向，使每个城市都可以访问城市 0 。返回需要变更方向的最小路线数。
 *
 * 题目数据 保证 每个城市在重新规划路线方向后都能到达城市 0 。
 *
 *
 *
 * 示例 1：
 *
 *
 *
 * 输入：n = 6, connections = [[0,1],[1,3],[2,3],[4,0],[4,5]]
 * 输出：3
 * 解释：更改以红色显示的路线的方向，使每个城市都可以到达城市 0 。
 *
 * 示例 2：
 *
 *
 *
 * 输入：n = 5, connections = [[1,0],[1,2],[3,2],[3,4]]
 * 输出：2
 * 解释：更改以红色显示的路线的方向，使每个城市都可以到达城市 0 。
 *
 * 示例 3：
 *
 * 输入：n = 3, connections = [[1,0],[2,0]]
 * 输出：0
 *
 *
 *
 *
 * 提示：
 *
 *
 * 2 <= n <= 5 * 10^4
 * connections.length == n-1
 * connections[i].length == 2
 * 0 <= connections[i][0], connections[i][1] <= n-1
 * connections[i][0] != connections[i][1]
 *
 *
 */
package jzoffer

import "sort"

// @lc code=start
func minReorderBfs(n int, connections [][]int) (res int) {
	// 反向建图   k存v[1] val存v[0] ,res是不需要翻转的路线数 返回 n-1-res
	/*
		0->1->3<-2
		 <-4->5

		先找出0和0的下一级,在通过bfs逐层找节点,如果一个节点有与被访问节点的反向路径,res++
	*/
	sort.Slice(connections, func(i, j int) bool {
		return connections[i][1] < connections[j][1]
	})
	graph := make(map[int][]int, n)
	for _, v := range connections {
		if v[1] == 0 {
			res++
			graph[0] = append(graph[0], v[0])
			continue
		}
		graph[v[0]] = append(graph[v[0]], v[1]) // v[0]-> v[1]
	}

	queue := graph[0]
	if len(queue) == n-1 {
		return n - 1 - res
	}
	visited := map[int]bool{0: true}
	for len(queue) > 0 {
		top := queue[len(queue)-1] // 取出队首
		visited[top] = true
		queue = queue[:len(queue)-1]
		if val, ok := graph[top]; ok {
			for _, v := range val {
				if visited[v] {
					continue
				}
				queue = append(queue, v)
			}
		}
		// 不二分会超时
		idx := sort.Search(n-1, func(i int) bool {
			return connections[i][1] >= top
		})
		for i := idx; i < len(connections); i++ {
			v := connections[i]
			if v[1] < top {
				continue
			}
			if v[1] > top {
				break
			}
			if !visited[v[0]] {
				res++
				queue = append(queue, v[0])
			}
		}
	}
	return n - 1 - res
}

// todo
// dfs解法
func minReorderDfs(n int, connections [][]int) (ans int) {
	g := make([][]int, n)
	for _, v := range connections {
		g[v[0]] = append(g[v[0]], v[1])  //正向边
		g[v[1]] = append(g[v[1]], -v[0]) //反向边
	}
	var dfs func(int, int)
	dfs = func(c, p int) {
		// ->c-> x
		for _, x := range g[c] {
			if x != p && -x != p {
				if x > 0 {
					ans++
				} else {
					x = -x
				}
				dfs(x, c)
			}
		}
	}
	dfs(0, n)
	return
}

type dirPath struct {
	city      int
	direction bool
}

func minReorder(n int, connections [][]int) int {
	graph := make(map[int][]dirPath)
	for _, con := range connections {
		from, to := con[0], con[1]
		graph[from] = append(graph[from], dirPath{to, true})
		graph[to] = append(graph[to], dirPath{from, false})
	}
	visited, ans := make(map[int]bool), 0
	var dfs func(int)
	dfs = func(i int) {
		visited[i] = true
		for _, p := range graph[i] {
			if !visited[p.city] {
				dfs(p.city)
				if p.direction == true {
					ans += 1
				}
			}
		}
	}
	dfs(0)
	return ans
}

// @lc code=end
