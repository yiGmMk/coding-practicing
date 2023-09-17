package jzoffer

import (
	"strings"
	"testing"
)

func Test151(t *testing.T) {
	ts := map[string]string{
		"the sky is blue": "blue is sky the",
		"  hello world  ": "world hello",
	}

	for k, v := range ts {
		if got := reverseWords1(k); got != v {
			t.Errorf("got:%s,want:%s", got, v)
		}
		if got := reverseWords(k); got != v {
			t.Errorf("got:%s,want:%s", got, v)
		}
	}
}

func reverseWords1(s string) string {
	t := strings.Fields(s)          //使用该函数可切割单个/多个空格，提取出单词
	for i := 0; i < len(t)/2; i++ { //遍历数组的前半段直接交换即可
		t[i], t[len(t)-1-i] = t[len(t)-1-i], t[i]
	}
	new := strings.Join(t, " ") //重新使用空格连接多个单词
	return (new)
}

// 作者：Jovial Dewdney1Qm
// 链接：https://leetcode.cn/problems/reverse-words-in-a-string/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
