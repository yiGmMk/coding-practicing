package util

import (
	"strconv"
	"strings"
	"unicode"
)

func Str2Int(strs []string) ([]int, error) {
	out := []int{}
	for _, str := range strs {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}
	return out, nil
}

// Leetcode 代码样例中给的字符串转数组
// [0,null,1] => []string{"0", "null", "1"}
func LeetcodeArrayStringToArray(array string) []string {
	res := []string{}
	if array == "" {
		return res
	}
	if array == "[]" {
		return res
	}
	// 删除空格
	array = strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, array)
	array = array[1 : len(array)-1]
	res = strings.Split(array, ",")
	return res
}
