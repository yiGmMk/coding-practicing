# param

一些环境变量，这些环境变量可以影响Go程序运行时的行为并输出Go运行时的一些信息以辅助在线诊断Go程序的问题

## GOMAXPROCS

GOMAXPROCS环境变量可用于设置Go程序启动后的逻辑处理器P的数量
从go1.5开始默认值为CPU核数

## GOGC

GOGC是一个整数值，默认为100。这个值代表一个比例值，100表示100%。
这个比例值的分子是上一次GC结束后到当前时刻新分配的堆内存大小（设为M），分母则是上一次GC结束后堆内存上的活动对象内存数据（live data，可达内存）的大小（设为N）

[GC](../runtime/RADME.md)

## GODEBUG

### schedtrace与scheddetail

```sh
GODEBUG=schedtrace=1000 godoc -http=:6000
```

schedtrace=1000中的1000的单位是毫秒，即每1000毫秒（1秒）输出一次goroutine调度器的内部信息

scheddetail=1，Go运行时会将更为详尽的调度器信息输出

### asyncpreemptoff(go 1.14+)

可以通过GODEBUG= asyncpreemptoff=1关闭这一新增的特性(抢占式调度)

#### 抢占式调度

go1.13及以前是协作式调度,只在有函数调用的地方才能插入抢占代码（埋点）

Go 1.14版本增加了基于系统信号的异步抢占调度，这样上面的deadloop所在的goroutine也可以被抢占了

## 其他

[参考](https://tip.golang.org/pkg/runtime/#hdr-Environment_Variables)
