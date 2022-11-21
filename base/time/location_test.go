package base

import (
	"fmt"
	"testing"
	"time"
)

// 时区转换
func TestLocation(t *testing.T) {
	locSrc, err := time.LoadLocation("America/Los_Angeles")
	if err != nil {
		fmt.Println("load time location failed:", err)
		return
	}

	date := time.Date(2020, 6, 18, 06, 0, 0, 0, locSrc)
	fmt.Println(date) // 美国西部洛杉矶时间，即太平洋时间

	locTo, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("load time location failed:", err)
		return
	}

	t1 := date.In(locTo) // 转换成美国东部纽约时间表示
	fmt.Println(t1)
}
