package jzoffer

import "testing"

func Test211(t *testing.T) {

	c := Constructor()
	c.AddWord("bad")
	c.AddWord("dad")
	c.AddWord("mad")
	if c.Search("pad") {
		t.Error("pad should be false")
	}
	if !c.Search("bad") {
		t.Error("bad should be true")
	}
	if !c.Search(".ad") {
		t.Error(".ad should be true")
	}
	if !c.Search("b..") {
		t.Error("b.. should be true")
	}

}
