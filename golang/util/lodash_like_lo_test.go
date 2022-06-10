package util

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/samber/lo"
	lop "github.com/samber/lo/parallel"
)

func TestLoApi(t *testing.T) {
	uniq := lo.Uniq([]int{12, 1, 2, 3, 4, 1})
	if !reflect.DeepEqual(uniq, []int{12, 1, 2, 3, 4}) {
		t.Errorf("uniq not right,uniq:%v", uniq)
	}

	result1 := lop.Map[int, string]([]int{1, 2, 3, 4}, func(x int, _ int) string {
		return "Hello"
	})
	result2 := lop.Map[int64, string]([]int64{1, 2, 3, 4}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})

	if !reflect.DeepEqual(result1, []string{"Hello", "Hello", "Hello", "Hello"}) {
		t.Errorf("result1 not right,result1:%v", result1)
	}
	if !reflect.DeepEqual(result2, []string{"1", "2", "3", "4"}) {
		t.Errorf("result2 not right,result2:%v", result2)
	}
}
