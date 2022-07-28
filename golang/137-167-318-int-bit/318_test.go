package jzoffer

import "testing"

var _318TestCases = []struct {
	words []string
	res   int
}{
	{
		words: []string{"abcw", "baz", "foo", "bar", "xtfn", "abcdef"},
		res:   16,
	},
	{
		words: []string{"a", "ab", "abc", "d", "cd", "bcd", "abcd"},
		res:   4,
	},
	{
		words: []string{"a", "aa", "aaa", "aaaa"},
		res:   0,
	},
	{
		words: []string{"bdcecbcadca", "caafd", "bcadc", "eaedfcd", "fcdecf", "dee", "bfedd", "ffafd", "eceaffa", "caabe", "fbdb", "acafbccaa", "cdc", "ecfdebaafde", "cddbabf", "adc", "cccce", "cbbe", "beedf", "fafbfdcb", "ceecfabedbd", "aadbedeaf", "cffdcfde", "fbbdfdccce", "ccada", "fb", "fa", "ec", "dddafded", "accdda", "acaad", "ba", "dabe", "cdfcaa", "caadfedd", "dcdcab", "fadbabace", "edfdb", "dbaaffdfa", "efdffceeeb", "aefdf", "fbadcfcc", "dcaeddd", "baeb", "beddeed", "fbfdffa", "eecacbbd", "fcde", "fcdb", "eac", "aceffea", "ebabfffdaab", "eedbd", "fdeed", "aeb", "fbb", "ad", "bcafdabfbdc", "cfcdf", "deadfed", "acdadbdcdb", "fcbdbeeb", "cbeb", "acbcafca", "abbcbcbaef", "aadcafddf", "bd", "edcebadec", "cdcbabbdacc", "adabaea", "dcebf", "ffacdaeaeeb", "afedfcadbb", "aecccdfbaff", "dfcfda", "febb", "bfffcaa", "dffdbcbeacf", "cfa", "eedeadfafd", "fcaa", "addbcad", "eeaaa", "af", "fafc", "bedbbbdfae", "adfecadcabe", "efffdaa", "bafbcbcbe", "fcafabcc", "ec", "dbddd", "edfaeabecee", "fcbedad", "abcddfbc", "afdafb", "afe", "cdad", "abdffbc", "dbdbebdbb"},
		res:   45,
	},
}

func Test318(t *testing.T) {
	for _, testCase := range _318TestCases {
		result := maxProduct01(testCase.words)
		if result != testCase.res {
			t.Errorf("longestCommonPrefix(%v) = %v, expected %v", testCase.words, result, testCase.res)
		}

		result = maxProduct(testCase.words)
		if result != testCase.res {
			t.Errorf("longestCommonPrefix(%v) = %v, expected %v", testCase.words, result, testCase.res)
		}
	}
}
