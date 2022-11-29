# runtime

## GMP

## GC

### Pacing算法

根据实时获取的堆内存状态数据及上一轮GC的结果数据决定新一轮GC启动的最佳时间.
并通过负反馈机制在两次GC之间的清除阶段以及GC过程中的并发标记阶段调节用户goroutine参与清除及标记的时间比例
如果某用户goroutine分配内存过快，该算法就会延缓其分配速度，让其更多地参与清除或辅助扫描标记的工作。
其目标是保证新一轮GC完成后的堆大小不超过这一轮GC后堆大小的期望目标值，以实现对堆内存大小的有效控制。

```sh
go build -o $dir/main $dir/gc.go

GODEBUG='gctrace=1' GOGC=100 $dir/main
```

启用环境变量GODEBUG='gctrace=1',Go运行时以特定的格式在每次GC结束后都输出一行信息,包括
GC的CPU占用率、每次GC的各个阶段的耗时、GC前后堆内存大小变化以及每次GC的目标等.

```sh
gc 1 @0.010s 3%: 0.10+2.6+0.006 ms clock, 0.20+0.68/0.13/2.4+0.012 ms cpu, 5->5->4 MB, 4 MB goal, 0 MB stacks, 0 MB globals, 2 P
gc 2 @0.025s 10%: 0.072+5.7+0.008 ms clock, 0.14+5.7/0.050/0.038+0.016 ms cpu, 9->9->8 MB, 8 MB goal, 0 MB stacks, 0 MB globals, 2 P
gc 3 @0.062s 6%: 0.027+13+0.005 ms clock, 0.055+0.58/1.5/19+0.010 ms cpu, 19->19->16 MB, 16 MB goal, 0 MB stacks, 0 MB globals, 2 P
```

GOGC值的越大,GC启动后面对的堆内存大小也是成比例快速增加的。
如果GOGC设置不合理，很可能出现GC还未来得及启动，堆内存就被耗尽、导致程序崩溃的情况

### 参数 GOGC

GOGC是一个整数值，默认为100。这个值代表一个比例值，100表示100%。
这个比例值的分子是上一次GC结束后到当前时刻新分配的堆内存大小（设为M），分母则是上一次GC结束后堆内存上的活动对象内存数据（live data，可达内存）的大小（设为N）

[GOGC](../go-tool/runtime-param.md)
