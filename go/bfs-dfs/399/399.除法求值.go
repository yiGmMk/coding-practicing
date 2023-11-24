/*
* @lc app=leetcode.cn id=399 lang=golang
*
* [399] 除法求值
*
* https://leetcode.cn/problems/evaluate-division/description/
*
  - algorithms
  - Medium (59.00%)
  - Likes:    1043
  - Dislikes: 0
  - Total Accepted:    86.7K
  - Total Submissions: 147.1K
  - Testcase Example:  '[["a","b"],["b","c"]]\n' +
    '[2.0,3.0]\n' +
    '[["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]'

*
* 给你一个变量对数组 equations 和一个实数值数组 values 作为已知条件，其中 equations[i] = [Ai, Bi] 和
* values[i] 共同表示等式 Ai / Bi = values[i] 。每个 Ai 或 Bi 是一个表示单个变量的字符串。
*
* 另有一些以数组 queries 表示的问题，其中 queries[j] = [Cj, Dj] 表示第 j 个问题，请你根据已知条件找出 Cj / Dj
* = ? 的结果作为答案。
*
* 返回 所有问题的答案 。如果存在某个无法确定的答案，则用 -1.0 替代这个答案。如果问题中出现了给定的已知条件中没有出现的字符串，也需要用 -1.0
* 替代这个答案。
*
* 注意：输入总是有效的。你可以假设除法运算中不会出现除数为 0 的情况，且不存在任何矛盾的结果。
*
* 注意：未在等式列表中出现的变量是未定义的，因此无法确定它们的答案。
*
*
*
* 示例 1：
*
*
* 输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries =
* [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
* 输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
* 解释：
* 条件：a / b = 2.0, b / c = 3.0
* 问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
* 结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
* 注意：x 是未定义的 => -1.0
*
* 示例 2：
*
*
* 输入：equations = [["a","b"],["b","c"],["bc","cd"]], values = [1.5,2.5,5.0],
* queries = [["a","c"],["c","b"],["bc","cd"],["cd","bc"]]
* 输出：[3.75000,0.40000,5.00000,0.20000]
*
*
* 示例 3：
*
*
* 输入：equations = [["a","b"]], values = [0.5], queries =
* [["a","b"],["b","a"],["a","c"],["x","y"]]
* 输出：[0.50000,2.00000,-1.00000,-1.00000]
*
*
*
*
* 提示：
*
*
* 1 <= equations.length <= 20
* equations[i].length == 2
* 1 <= Ai.length, Bi.length <= 5
* values.length == equations.length
* 0.0 < values[i] <= 20.0
* 1 <= queries.length <= 20
* queries[i].length == 2
* 1 <= Cj.length, Dj.length <= 5
* Ai, Bi, Cj, Dj 由小写英文字母与数字组成
*
*
*/
package jzoffer

// @lc code=start
func calcEquation1(equations [][]string, values []float64, queries [][]string) (res []float64) {
	// 建图
	graph := make(map[string]map[string]float64)
	for i, v := range equations {
		if _, ok := graph[v[0]]; !ok {
			graph[v[0]] = map[string]float64{v[1]: values[i]}
		} else {
			graph[v[0]][v[1]] = values[i]
		}

		if _, ok := graph[v[1]]; !ok {
			graph[v[1]] = map[string]float64{v[0]: 1 / values[i]}
		} else {
			graph[v[1]][v[0]] = 1 / values[i]
		}
	}
	res = make([]float64, len(queries))
	// dfs 找起点
	for i, v := range queries {
		// 起点,终点 必须都在图里才有解
		nodes0, ok := graph[v[0]]
		if !ok {
			res[i] = -1.0
			continue
		}
		if _, ok := graph[v[1]]; !ok {
			res[i] = -1.0
			continue
		}
		// 都在图里,且相同的 直接返回1
		if v[0] == v[1] {
			res[i] = 1.0
			continue
		}
		// 在图里有对应的edge
		if val, ok := nodes0[v[1]]; ok {
			res[i] = val
			continue
		}

		// bfs查找多层的路径,把边的权重带上
		type edgeNode struct {
			node   string
			weight float64
		}
		queue := []edgeNode{}
		for k := range nodes0 {
			queue = append(queue, edgeNode{k, nodes0[k]})
		}
		visited := map[string]bool{v[0]: true}
		for len(queue) > 0 {
			top := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if top.node == v[1] {
				res[i] = top.weight
				break
			}
			visited[top.node] = true
			for k, kWeight := range graph[top.node] {
				if visited[k] {
					continue
				}
				queue = append(queue, edgeNode{k, kWeight * top.weight})
			}
		}
		if res[i] < 0.0000000000000000000001 { // 没有设置过值的0 改成-1
			res[i] = -1.0
		}
	}
	return
}

func calcEquation(equations [][]string, values []float64, queries [][]string) (res []float64) {
	// 建图
	graph := make(map[string]map[string]float64)
	for i, v := range equations {
		if _, ok := graph[v[0]]; !ok {
			graph[v[0]] = map[string]float64{v[1]: values[i]}
		} else {
			graph[v[0]][v[1]] = values[i]
		}

		if _, ok := graph[v[1]]; !ok {
			graph[v[1]] = map[string]float64{v[0]: 1 / values[i]}
		} else {
			graph[v[1]][v[0]] = 1 / values[i]
		}
	}
	res = make([]float64, len(queries))
	// dfs 找起点
	for i, v := range queries {
		start, end := v[0], v[1]
		nodes0, ok0 := graph[start]
		_, ok1 := graph[end]
		if !ok0 || !ok1 { // 起点,终点 必须都在图里才有解
			res[i] = -1.0
			continue
		}
		if start == end { // 都在图里,且相同的 直接返回1
			res[i] = 1.0
			continue
		}

		// bfs查找多层的路径,把边的权重带上
		type edgeNode struct {
			node   string
			weight float64
		}
		queue := []edgeNode{}
		for k := range nodes0 {
			queue = append(queue, edgeNode{k, nodes0[k]})
		}
		visited := map[string]bool{start: true}
		for len(queue) > 0 {
			top := queue[len(queue)-1]
			queue = queue[:len(queue)-1]
			if top.node == end {
				res[i] = top.weight
				break
			}
			visited[top.node] = true
			for k, kWeight := range graph[top.node] {
				if visited[k] {
					continue
				}
				queue = append(queue, edgeNode{k, kWeight * top.weight})
			}
		}
		if res[i] < 0.0000000000000000000001 { // 没有设置过值的0 改成-1
			res[i] = -1.0
		}
	}
	return
}

// @lc code=end
