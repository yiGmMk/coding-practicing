package jzoffer

import "testing"

var _648TestCases = []struct {
	dictionary []string
	sentence   string
	expected   string
}{
	/*
	 * 输入：dictionary = ["cat","bat","rat"], sentence = "the cattle was rattled by the battery"
	 * 输出："the cat was rat by the bat"
	 */
	{
		dictionary: []string{"cat", "bat", "rat"},
		sentence:   "the cattle was rattled by the battery",
		expected:   "the cat was rat by the bat",
	},
	/*
	 * 输入：dictionary = ["a","b","c"], sentence = "aadsfasf absbs bbab cadsfafs"
	 * 输出："a a b c"
	 */
	{
		dictionary: []string{"a", "b", "c"},
		sentence:   "aadsfasf absbs bbab cadsfafs",
		expected:   "a a b c",
	},
}

func Test648(t *testing.T) {
	for _, tc := range _648TestCases {
		actual := replaceWords(tc.dictionary, tc.sentence)
		if actual != tc.expected {
			t.Errorf("expect:%s,got %v", tc.expected, actual)
		}

		actual = replaceWordsMap(tc.dictionary, tc.sentence)
		if actual != tc.expected {
			t.Errorf("using map,expect:%s,got %v", tc.expected, actual)
		}
	}
}
