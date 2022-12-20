package jzoffer

import "testing"

var _1300TestCase = []struct {
	nums   []int
	target int
	expect int
}{
	{
		nums:   []int{4, 9, 3},
		target: 10,
		expect: 3,
	},
	{
		nums:   []int{2, 3, 5},
		target: 10,
		expect: 5,
	},
	{
		nums:   []int{60864, 25176, 27249, 21296, 20204},
		target: 56803,
		expect: 11361,
	},
}

func Test1300(t *testing.T) {
	for i, v := range _1300TestCase {
		got := findBestValue(v.nums, v.target)
		if got != v.expect {
			t.Errorf("[%d].got:%d,want:%d", i, got, v.expect)
		}
	}
}
