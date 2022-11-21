package base

import (
	"fmt"
	"testing"
	"time"
)

func TestHasMonotonic(t *testing.T) {
	// 没有单调时间
	fmt.Println("-----------date-----------------")
	constructTimeByDate()
	fmt.Println("-----------parse-----------------")
	constructTimeByParse()

	// 通过time.Now函数获取的当前时间中则既包含挂钟时间，也包含单调时间
	fmt.Println("-----------now-----------------")
	time.Sleep(time.Second)
	now := time.Now()
	fmt.Println(now)
	dumpWallAndExt(now)
}
