# SIGNAL

## 同步信号

由程序执行错误引发的信号，包括SIGBUS（总线错误/硬件异常）、SIGFPE（算术异常）和SIGSEGV（段错误/无效内存引用）

Go运行时不会简单地将信号通过channel发送到用户层并等待用户层的异步处理，而是直接将信号转换成一个运行时**panic**并抛出。

## 异步信号

同步信号之外的信号都被Go划归为异步信号。异步信号不是由程序执行错误引起的，而是由其他程序或操作系统内核发出的.

- SIGHUP、SIGINT和SIGTERM这三个信号将导致程序直接退出
- SIGQUIT、SIGILL、SIGTRAP、SIGABRT、SIGSTKFLT、SIGEMT和SIGSYS在导致程序退出的同时，还会将程序退出时的栈状态打印出来
- **SIGPROF**信号则是被Go运行时用于实现运行时CPU性能剖析指标采集

## http服务gracefully exit

[样例](server/gracefully-exit-with-notify.go)

- 通过Notify捕获信号
- 使用http包提供的Shutdown来实现HTTP服务内部的退出清理工作,包括立即关闭所有listener、关闭所有空闲的连接、等待处于活动状态的连接处理完毕（变成空闲连接）等。
- http.Sever的RegisterOnShutdown方法,用于在服务关闭前清理其他资源、做收尾工作
