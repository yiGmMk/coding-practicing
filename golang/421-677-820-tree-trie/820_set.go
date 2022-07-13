package jzoffer

import (
	datastructure "github.com/duke-git/lancet/v2/datastructure/set"
	"github.com/liyue201/gostl/ds/set"
)

// https://leetcode.cn/problems/iSwD2y/solution/zui-duan-de-dan-ci-bian-ma-by-leetcode-s-qjxw/
func minimumLengthEncodingSet(words []string) int {
	s := set.New()
	for i, _ := range words {
		s.Insert(words[i])
	}
	out := 0
	for _, work := range words {
		for j := 1; j < len(work); j++ {
			s.Erase(work[j:])
		}
	}
	for iter := s.Begin(); iter.IsValid(); iter.Next() {
		out += len(iter.Value().(string)) + 1
	}
	return out
}

func minimumLengthEncodingSet2(words []string) int {
	s := datastructure.NewSet[string](words...)

	out := 0
	for _, work := range words {
		for j := 1; j < len(work); j++ {
			s.Delete(work[j:])
		}
	}
	s.Iterate(func(value string) {
		out += len(value) + 1
	})
	return out
}
