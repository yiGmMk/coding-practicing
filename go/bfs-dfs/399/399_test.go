package jzoffer

import (
	"fmt"
	"testing"
)

func Test399(t *testing.T) {
	fmt.Println(calcEquation(
		// [["a","b"],["c","b"],["d","b"],["w","x"],["y","x"],["z","x"],["w","d"]]
		[][]string{
			{"a", "b"}, {"c", "b"}, {"d", "b"}, {"w", "x"}, {"y", "x"}, {"z", "x"}, {"w", "d"}},
		[]float64{2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0},
		// [["a","c"],["b","c"],["a","e"],["a","a"],["x","x"],["a","z"]]
		[][]string{{"a", "c"}, {"b", "c"}, {"a", "e"}, {"a", "a"}, {"x", "x"}, {"a", "z"}}))

	/*
	 * 输入：equations = [["a","b"],["b","c"]], values = [2.0,3.0], queries =
	 * [["a","c"],["b","a"],["a","e"],["a","a"],["x","x"]]
	 * 输出：[6.00000,0.50000,-1.00000,1.00000,-1.00000]
	 * 解释：
	 * 条件：a / b = 2.0, b / c = 3.0
	 * 问题：a / c = ?, b / a = ?, a / e = ?, a / a = ?, x / x = ?
	 * 结果：[6.0, 0.5, -1.0, 1.0, -1.0 ]
	 * 注意：x 是未定义的 => -1.0
	 */
	fmt.Println(calcEquation(
		[][]string{{"a", "b"}, {"b", "c"}},
		[]float64{2.0, 3.0},
		[][]string{{"a", "c"}, {"b", "a"}, {"a", "e"}, {"a", "a"}, {"x", "x"}}))

}