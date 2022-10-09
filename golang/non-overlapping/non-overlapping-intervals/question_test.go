package jzoffer

import "testing"

var _qTestCases = []struct {
	starting, ending []int
	expected         int
}{
	{
		starting: []int{5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7},
		ending:   []int{7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8},
		expected: 2,
	},
	{
		starting: []int{1, 2, 4, 3, 7},
		ending:   []int{3, 4, 6, 5, 8},
		expected: 1,
	},
	{
		starting: []int{5, 2, 3, 7},
		ending:   []int{7, 2, 4, 8},
		expected: 2,
	},
	{
		starting: []int{5, 2, 2, 3, 7, 7, 7},
		ending:   []int{7, 2, 2, 4, 8, 8, 8},
		expected: 2,
	},
}

func TestThreeNonOverlappingIntervals(t *testing.T) {
	for _, v := range _qTestCases {
		res := getThreeNonOverlappingIntervals(v.starting, v.ending)
		if res != v.expected {
			t.Errorf("1.want:%d,got:%d", v.expected, res)
		}

		res = getThreeNonOverlappingIntervalsTimeout(v.starting, v.ending)
		if res != v.expected {
			t.Errorf("2.want:%d,got:%d", v.expected, res)
		}

		res = getThreeNonOverlappingIntervalsTimeout2(v.starting, v.ending)
		if res != v.expected {
			t.Errorf("3.want:%d,got:%d", v.expected, res)
		}
	}
}

func BenchmarkNonintervals(b *testing.B) {
	n := len(_qTestCases)
	debug = false
	for i := 0; i < b.N; i++ {
		getThreeNonOverlappingIntervalsTimeout(_qTestCases[i%n].starting, _qTestCases[i%n].ending)
	}
}

func BenchmarkNonintervals1(b *testing.B) {
	n := len(_qTestCases)
	debug = false
	for i := 0; i < b.N; i++ {
		getThreeNonOverlappingIntervalsTimeout2(_qTestCases[i%n].starting, _qTestCases[i%n].ending)
	}
}

func BenchmarkNonintervals2(b *testing.B) {
	debug = false
	for i := 0; i < b.N; i++ {
		getThreeNonOverlappingIntervalsTimeout2(_qTestCases[0].starting, _qTestCases[0].ending)
	}
}
