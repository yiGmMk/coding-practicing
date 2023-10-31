package jzoffer

import (
	"testing"
)

func Test394(t *testing.T) {
	if "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef" !=
		decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef") {
		t.Error("3[z]2[2[y]pq4[2[jk]e1[f]]]ef failed")
	}
	if "aaabcbc" != (decodeString("3[a]2[bc]")) {
		t.Error("3[a]2[bc] failed")
	}

	if "accaccacc" != decodeString("3[a2[c]]") {
		t.Error("3[a2[c]] decode failed")
	}
	if "leetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcodeleetcode" !=
		decodeString("100[leetcode]") {
		t.Error("100[leetcode]")
	}
}
