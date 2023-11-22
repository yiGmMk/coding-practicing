package jzoffer

import (
	"fmt"
	"testing"
)

func Test208(t *testing.T) {
	// utf-8 变长编码,使用1~3个字节存储
	str := "01abAB,./中"
	for i, v := range str {
		fmt.Printf("[%d].%s:string[i]=%c,rune-'0'=%d,rune-'a'=%d,rune=%v,%%c=%c\n",
			i, string(v), str[i], v-'0', v-'a', v, v)
	}

	// [0].0:string[i]=0,rune-'0'=0,rune-'a'=-49,rune=48,%c=0
	// [1].1:string[i]=1,rune-'0'=1,rune-'a'=-48,rune=49,%c=1
	// [2].a:string[i]=a,rune-'0'=49,rune-'a'=0,rune=97,%c=a
	// [3].b:string[i]=b,rune-'0'=50,rune-'a'=1,rune=98,%c=b
	// [4].A:string[i]=A,rune-'0'=17,rune-'a'=-32,rune=65,%c=A
	// [5].B:string[i]=B,rune-'0'=18,rune-'a'=-31,rune=66,%c=B
	// [6].,:string[i]=,,rune-'0'=-4,rune-'a'=-53,rune=44,%c=,
	// [7]..:string[i]=.,rune-'0'=-2,rune-'a'=-51,rune=46,%c=.
	// [8]./:string[i]=/,rune-'0'=-1,rune-'a'=-50,rune=47,%c=/
	// [9].中:string[i]=ä,rune-'0'=19965,rune-'a'=19916,rune=20013,%c=中
}
