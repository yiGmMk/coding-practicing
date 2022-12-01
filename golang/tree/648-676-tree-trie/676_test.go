package jzoffer

import "testing"

func Test676(t *testing.T) {
	m := Constructor()
	m.BuildDict([]string{"hello", "world", "leetcode"})
	if m.Search("hello") {
		t.Error("should be false")
	}
	if !m.Search("hhllo") {
		t.Error("should be true")
	}

	if m.Search("hell") {
		t.Error("should be false")
	}

	if m.Search("leetcoded") {
		t.Error("should be false")
	}
}
