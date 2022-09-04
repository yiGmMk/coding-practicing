package jzoffer

import "testing"

var _820TestCase = []struct {
	words []string
	res   int
}{
	{
		/** 输入：words = ["time", "me", "bell"] * 输出：10		*/
		words: []string{"time", "me", "bell"},
		res:   10,
	},
	{
		/*
		 * 输入：words = ["t"] * 输出：2
		 */
		words: []string{"t"},
		res:   2,
	},
}

func Test820(t *testing.T) {
	for _, tc := range _820TestCase {

		res := minimumLengthEncodingSet(tc.words)
		if res != tc.res {
			t.Errorf("%v, want: %v", res, tc.res)
		}

		res = minimumLengthEncodingSet2(tc.words)
		if res != tc.res {
			t.Errorf("%v, want: %v", res, tc.res)
		}

		if res := minimumLengthEncoding(tc.words); res != tc.res {
			t.Errorf("got %v, want %v", res, tc.res)
		}
	}
}
