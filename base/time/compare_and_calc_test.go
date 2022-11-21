package base

import (
	"fmt"
	"testing"
	"time"
)

func TestCompare(t *testing.T) {
	now := time.Now()
	fmt.Println(now) //北京时间

	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		fmt.Println("load time location failed:", err)
		return
	}

	t1 := now.In(loc) // 转换成美国东部纽约时间表示
	fmt.Println(now == t1)

	// 使用equal方法
	fmt.Println("using equal", now.Equal(t1))
}

func TestCalc(t *testing.T) {
	subTwoTimeHasMonotonic := func(t *testing.T) {
		t1 := time.Now()
		time.Sleep(time.Second * 5)
		t2 := time.Now()
		diff := t2.Sub(t1)
		fmt.Printf("[hasMonotonic = 1] t2 - t1 = %v\n", diff)
	}

	subTwoTimeNoMonotonic := func(t *testing.T) {
		t1 := time.Date(2020, 6, 18, 0, 0, 0, 0, time.UTC)
		t2 := time.Date(2020, 6, 18, 12, 0, 0, 0, time.UTC)
		diff := t2.Sub(t1)
		fmt.Printf("[hasMonotonic = 0] t2 - t1 = %v\n", diff)
	}

	utilTwoTimeHasMonotonic := func(t *testing.T) {
		t1 := time.Now().AddDate(0, 0, 1)
		d := time.Until(t1)
		fmt.Printf("[hasMonotonic= 1] until t1 =%v \n ", d)
	}

	t.Run("subTwoTimeHasMonotonic", subTwoTimeHasMonotonic)
	t.Run("subTwoTimeNoMonotonic", subTwoTimeNoMonotonic)
	t.Run("utilTwoTimeHasMonotonic", utilTwoTimeHasMonotonic)
}
