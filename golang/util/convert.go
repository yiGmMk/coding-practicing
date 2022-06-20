package util

import "strconv"

func Str2Int(strs []string) ([]int, error) {
	out := []int{}
	for _, str := range strs {
		i, err := strconv.Atoi(str)
		if err != nil {
			return nil, err
		}
		out = append(out, i)
	}
	return out, nil
}
