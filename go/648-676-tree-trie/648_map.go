package jzoffer

import "strings"

// https://leetcode.cn/problems/UhWRSj/solution/ti-huan-dan-ci-by-leetcode-solution-9reh/
func replaceWordsMap(dictionary []string, sentence string) string {
	trie := make(map[string]bool)
	for _, word := range dictionary {
		trie[word] = true
	}

	words := strings.Split(sentence, " ")
	for i, word := range words {
		for j := 1; j <= len(word); j++ {
			if trie[word[:j]] {
				words[i] = word[:j]
				break
			}
		}
	}

	return strings.Join(words, " ")
}
