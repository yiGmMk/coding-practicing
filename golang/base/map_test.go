package base

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	// 在遍历map的过程中,新增/删除key,新增/删除的会不会出现在遍历结果中未知,可能出现,也可能不出现
	for {
		var m = map[string]int{"a": 1, "b": 2, "c": 3}
		count := 0
		for k, v := range m {
			if count == 0 {
				delete(m, "a")
			}
			count++
			fmt.Println(k, v)
		}
		fmt.Println("count", count)

		if count == 2 {
			break
		}
	}

	/*
	 *a 1
	 *b 2
	 *c 3
	 *count 3
	 *
	 *
	 *b 2
	 *c 3
	 *count 2
	 */

}
