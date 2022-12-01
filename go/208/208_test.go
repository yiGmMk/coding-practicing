package jzoffer

import "testing"

/*
 * 输入
 * ["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
 * [[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
 * 输出
 * [null, null, true, false, true, null, true]
 *
 */
func Test208(t *testing.T) {

	trie := Constructor()
	trie.Insert("apple")
	res := trie.Search("apple")
	if !res {
		t.Error("apple not found")
	}
	res = trie.Search("app")
	if res {
		t.Error("app found")
	}
	trie.StartsWith("app")
	trie.Insert("app")
	res = trie.Search("app")
	if !res {
		t.Error("app not found")
	}
}
