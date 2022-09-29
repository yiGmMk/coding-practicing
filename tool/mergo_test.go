package tool

import (
	"fmt"
	"testing"

	"github.com/imdario/mergo"
)

// https://github.com/imdario/mergo
// 知乎: https://zhuanlan.zhihu.com/p/112806927
func TestMergo(t *testing.T) {
	// mergo 不处理未导出的值
	type Person struct {
		Name           string
		Age            int
		_salary        int
		WithDefaultVal int
		House          []string
	}

	sp := Person{Name: "jack", Age: 11, _salary: 22, WithDefaultVal: 33, House: []string{"bj", "sz"}}

	// 默认不覆盖有值的
	var spCopy Person
	spCopy.WithDefaultVal = 888
	spCopy.House = []string{"nj"}
	if err := mergo.Merge(&spCopy, sp); err != nil {
		fmt.Printf("map:%+v", err)
		return
	}
	fmt.Println("not override", spCopy)

	// 覆盖
	if err := mergo.MergeWithOverwrite(&spCopy, sp); err != nil {
		fmt.Printf("map:%+v", err)
		return
	}
	fmt.Println("    override", spCopy)

	// 切片,带参数可以append,而不是覆盖
	spCopy.House = []string{"nj"}
	if err := mergo.MergeWithOverwrite(&spCopy, sp, mergo.WithAppendSlice); err != nil {
		fmt.Printf("map:%+v", err)
		return
	}
	fmt.Println("    append", spCopy)

	// map的val需要是interface
	mp := make(map[string]interface{})
	if err := mergo.Map(&mp, sp); err != nil {
		fmt.Printf("merge:%+v ", err)
		return
	}
	fmt.Println(mp)
}
