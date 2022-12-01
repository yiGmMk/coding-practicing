package jzoffer

import "testing"

var _qTestCases = []struct {
	starting, ending []int
	expected         int
}{
	{
		starting: []int{5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7, 5, 2, 2, 3, 7, 7, 7},
		ending:   []int{7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8, 7, 2, 2, 4, 8, 8, 8},
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
	debug = false
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

		res = getThreeNonOverlappingIntervals1(v.starting, v.ending)
		if res != v.expected {
			t.Errorf("4.want:%d,got:%d", v.expected, res)
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

// cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
// BenchmarkNonintervals1-2   	   40995	     27378 ns/op	   15126 B/op	     227 allocs/op
func BenchmarkNonintervals1(b *testing.B) {
	n := len(_qTestCases)
	debug = false
	for i := 0; i < b.N; i++ {
		getThreeNonOverlappingIntervalsTimeout2(_qTestCases[i%n].starting, _qTestCases[i%n].ending)
	}
}

// cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
// BenchmarkNonintervals2-2   	    9703	    103166 ns/op	   58712 B/op	     860 allocs/op
func BenchmarkNonintervals2(b *testing.B) {
	debug = false
	for i := 0; i < b.N; i++ {
		getThreeNonOverlappingIntervalsTimeout2(_qTestCases[0].starting, _qTestCases[0].ending)
	}
}

// cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
// BenchmarkNonintervals3-2   	   47332	     24897 ns/op	   15462 B/op	     239 allocs/op
func BenchmarkNonintervals3(b *testing.B) {
	n := len(_qTestCases)
	debug = false
	for i := 0; i < b.N; i++ {
		getThreeNonOverlappingIntervals(_qTestCases[i%n].starting, _qTestCases[i%n].ending)
	}
}

// cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
// BenchmarkNonintervals4-2   	   13000	     92333 ns/op	   58936 B/op	     869 allocs/op
func BenchmarkNonintervals4(b *testing.B) {
	debug = false
	for i := 0; i < b.N; i++ {
		getThreeNonOverlappingIntervals(_qTestCases[0].starting, _qTestCases[0].ending)
	}
}

// cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
// BenchmarkNonintervals5-2   	  210643	      5941 ns/op	     380 B/op	      10 allocs/op
func BenchmarkNonintervals5(b *testing.B) {
	n := len(_qTestCases)
	debug = false
	for i := 0; i < b.N; i++ {
		getThreeNonOverlappingIntervals1(_qTestCases[i%n].starting, _qTestCases[i%n].ending)
	}
}

// cpu: Intel(R) Xeon(R) Platinum 8269CY CPU @ 2.50GHz
// BenchmarkNonintervals6-2   	   58915	     20017 ns/op	     328 B/op	      10 allocs/op
func BenchmarkNonintervals6(b *testing.B) {
	debug = false
	for i := 0; i < b.N; i++ {
		getThreeNonOverlappingIntervals1(_qTestCases[0].starting, _qTestCases[0].ending)
	}
}
