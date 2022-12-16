package jzoffer

import "testing"

var _540TestCase = []struct {
	vs     []int
	expect int
}{
	{
		vs:     []int{1, 1, 2, 3, 3, 4, 4, 8, 8},
		expect: 2,
	},
	{
		vs:     []int{3, 3, 7, 7, 10, 11, 11},
		expect: 10,
	},
	{
		vs:     []int{1, 1, 2, 2, 3},
		expect: 3,
	},
}

func Test240(t *testing.T) {
	for i, v := range _540TestCase {
		got := singleNonDuplicate(v.vs)
		if got != v.expect {
			t.Errorf("[%d].got:%v,want:%v", i, got, v.expect)
		}
		got = singleNonDuplicate1(v.vs)
		if got != v.expect {
			t.Errorf("[%d]1.got:%v,want:%v", i, got, v.expect)
		}
		got = singleNonDuplicate2(v.vs)
		if got != v.expect {
			t.Errorf("[%d]2.got:%v,want:%v", i, got, v.expect)
		}
		got = singleNonDuplicate3(v.vs)
		if got != v.expect {
			t.Errorf("[%d]3.got:%v,want:%v", i, got, v.expect)
		}
	}
}
