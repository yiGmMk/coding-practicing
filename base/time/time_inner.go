package base

import (
	"fmt"
	"time"
	"unsafe"
)

// $GOROOT/src/time.go(go1.14) Time的内部结构
// type Time struct {
// 	wall uint64
// 	ext  int64
// 	loc  *Location
// }

// chapter9/sources/go-time-operations/construct_time_with_func_date.go
// <go语言精进之路>(https://weread.qq.com/web/reader/b8f32d2072895edbb8fbb04ka3f32db0244a3f390d88bb9)
// 1.挂钟时间与单调时间(单调时间则是永远不会出现“时间倒流”现象的。
// 1.0单调时间表示的是程序进程启动之后流逝的时间，两次采集的单调时间之差永远不可能为负数。)
// 1.1在hasMonotonic比特位为1的情况下，ext字段表示程序进程启动后的单调流逝时间，以纳秒为单位。
// 1.2而当hasMonotonic为0时，time.Time结构体仅表示挂钟时间
// 1.3通过time.Now函数获取的当前时间中则既包含挂钟时间，也包含单调时间
func dumpWallAndExt(t time.Time) {
	var hasMonotonic int

	// 输出wall字段的值
	pWall := (*uint64)(unsafe.Pointer(&t))
	fmt.Printf("0x%X\n", *pWall)
	if (1<<63)&(*pWall) != 0 {
		hasMonotonic = 1
	}
	fmt.Printf("hasMonotonic = %d\n", hasMonotonic)

	// 输出ext字段的值
	pExt := (*int64)(unsafe.Pointer((uintptr(unsafe.Pointer(&t)) + unsafe.Sizeof(uint64(0)))))
	fmt.Printf("0x%X\n", *pExt)
	fmt.Printf("%d\n", *pExt/86400/365) // 粗略计算距今的年数
}

func constructTimeByDate() {
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println("load time location failed:", err)
		return
	}
	t := time.Date(2020, 6, 18, 06, 0, 0, 10000, loc)
	fmt.Println(t)

	dumpWallAndExt(t)
}

func constructTimeByParse() {
	t, _ := time.Parse(time.RFC3339, "2020-06-18T06:00:00.00001+08:00")
	fmt.Println(t)
	dumpWallAndExt(t)
}
