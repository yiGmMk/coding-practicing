package jzoffer

import "testing"

func Test677(t *testing.T) {
	m := Constructor()
	m.Insert("apple", 3)
	sum := m.Sum("ap")
	if sum != 3 {
		t.Error("should be 3")
	}
	m.Insert("app", 2)
	sum = m.Sum("ap")
	if sum != 5 {
		t.Error("should be 5")
	}
}
